package repositories

import (
	repositories "user-service/repositories/user_repository"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetUser() repositories.IUserRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry{
	return &Registry{db: db}
}

func (rr *Registry) GetUser() repositories.IUserRepository{
	return repositories.NewUserRepository(rr.db)
}