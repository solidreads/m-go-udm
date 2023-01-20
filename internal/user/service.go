package user

import (
	"log"
)

type Service interface {
	Create(firstNAme, lastName, email, phone string) (string, error)
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s *service) Create(firstNAme, lastName, email, phone string) (string, error) {
	s.log.Println("service create user")
	s.repo.Create(&User{
		FirstName: firstNAme,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	})
	return "", nil
}
