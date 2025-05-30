package middlewares

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"user-service/common/response"
	"user-service/config"
	"user-service/constants"
	services "user-service/services/user_service"

	errConstant "user-service/constants/error"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// digunakan untuk menangani ERROR panic
func HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("Recovered from panic: %v", err)
				ctx.JSON(
					http.StatusInternalServerError,
					response.Response{
						Status:  constants.Error,
						Message: errConstant.ErrInternalServerError.Error(),
					})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

// untuk memberi batasan untuk request yang masuk ke dalam sistem
func Ratelimit(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusTooManyRequests, response.Response{
				Status:  constants.Error,
				Message: errConstant.ErrTooManyRequests.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func extractBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")
	if len(arrayToken) == 2 {
		return arrayToken[1]
	}
	return ""
}

func responseUnauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, response.Response{
		Status:  constants.Error,
		Message: message,
	})
	ctx.Abort()
}

func validateAPIKey(ctx *gin.Context) error {
	apiKey := ctx.GetHeader(constants.XApiKey)
	requestAt := ctx.GetHeader(constants.XRequestAt)
	serviceName := ctx.GetHeader(constants.XServiceName)
	signatureKey := config.Config.SignatureKey


	validateKey := fmt.Sprintf("%s:%s:%s", serviceName, signatureKey, requestAt)
	hash := sha256.New()
	hash.Write([]byte(validateKey))
	resultHash := hex.EncodeToString(hash.Sum(nil))

	if apiKey != resultHash {
		return errConstant.ErrUnauthorized
	}
	return nil

}

func validateBearerToken(ctx *gin.Context, token string) error {
	if !strings.HasPrefix(token, "Bearer ") {
		return errConstant.ErrUnauthorized
	}

	tokenString := extractBearerToken(token)
	if tokenString == "" {
		return errConstant.ErrUnauthorized
	}

	claims := &services.Claims{}
	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errConstant.ErrInvalidToken
		}

		jwtSecret := []byte(config.Config.JwtSecretKey)
		return jwtSecret, nil
	})

	if err != nil || !tokenJwt.Valid {
		return errConstant.ErrUnauthorized
	}

	userLogin := ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), constants.UserLogin, claims.User))
	ctx.Request = userLogin
	ctx.Set(constants.Token, token)
	return nil
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(constants.Authorization)
		if token == "" {
			fmt.Println("token is empty")
			responseUnauthorized(ctx, errConstant.ErrUnauthorized.Error())
			return
		}

		if err := validateBearerToken(ctx, token); err != nil {
			fmt.Println("error validate token", err)
			responseUnauthorized(ctx, err.Error())
			return
		}

		if err := validateAPIKey(ctx); err != nil {
			fmt.Println("error validate api key", err)
			responseUnauthorized(ctx, err.Error())
			return
		}

		ctx.Next()
	}
}
