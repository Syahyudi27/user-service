package controllers

import (
	controller "user-service/controllers/user_controller"
	"user-service/services"
)

type RegistryStruct struct {	
	registryService services.IServiceRegistry
}

type IRegistryController interface {
	GetUserController() controller.IUserController
}

func NewControllerRegistry(registryService services.IServiceRegistry) IRegistryController{
	return &RegistryStruct{registryService: registryService}
	}


func (r *RegistryStruct) GetUserController() controller.IUserController{
	return controller.NewUserController(r.registryService)
}