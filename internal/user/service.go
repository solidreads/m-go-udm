package user

import (
	"log"
)

type Service interface {
	Create(firstNAme, lastName, email, phone string) (*User, error)
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

func (s *service) Create(firstNAme, lastName, email, phone string) (*User, error) {
	s.log.Println("service create user")
	user := User{
		FirstName: firstNAme,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
