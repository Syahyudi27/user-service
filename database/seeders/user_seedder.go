package seeders

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user-service/constants"
	"user-service/domain/models"
)

func RunUserSeeders(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("Admin"), bcrypt.DefaultCost)
	user := models.User{
		UUID:        uuid.New(),
		Name:        "Administrator",
		Username:    "admin",
		Password:    string(password),
		PhoneNumber: "08123456789",
		Email:       "admin@gmail.com",
		RoleID:      constants.Admin,
	}

	err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user: %v", err)
		panic(err)
	}
	logrus.Infof("user %s successfully seeded", user.Username)
}
