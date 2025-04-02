package seeders

import "gorm.io/gorm"

type SeederRegistry struct {
	db *gorm.DB
}

type ISeederRegistry interface {
	Run()
}

func NewSeederRegistry(db *gorm.DB) ISeederRegistry {
	return &SeederRegistry{db: db}
}

func (s *SeederRegistry) Run() {
	RunRoleSeeders(s.db)
	RunUserSeeders(s.db)
}


