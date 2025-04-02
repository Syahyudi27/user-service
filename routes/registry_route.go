package routes

import (
	"user-service/controllers"
	routes "user-service/routes/user_route"

	"github.com/gin-gonic/gin"
)

type RegisteryRoute struct {
	controllers controllers.IRegistryController
	group       *gin.RouterGroup
}

type IRouterRegistry interface {
	Serve()
}

func NewRouterRegistry(controllers controllers.IRegistryController, group *gin.RouterGroup) IRouterRegistry {
	return &RegisteryRoute{
		controllers: controllers, 
		group: group}
}

func (r *RegisteryRoute) userRoute() routes.IUserRoute {
	return routes.NewUserRoute(r.controllers, r.group)
}


func (r *RegisteryRoute) Serve() {
	r.userRoute().Run()
}