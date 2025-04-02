package services

import (
	"user-service/repositories"
	services "user-service/services/user_service"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{repository: repository}
}

func (r *Registry) GetUser() services.IUserService {
	return services.NewuserService(r.repository)
}