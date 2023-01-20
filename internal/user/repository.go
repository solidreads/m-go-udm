package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Create(user *User) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db}
}

func (repo *repo) Create(user *User) error {
	user.ID = uuid.New().String()

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("El usuario fue creado exitosamente con el ID: ", user.ID)
	return nil
}
