package repository

import (
	"gorm/model"

	"gorm.io/gorm"
)

type UserCredentialRepository interface {
	FindByUsername(username string) (model.UserCredential, error)
}

type userCredentialRepository struct {
	db *gorm.DB
}

func (u *userCredentialRepository) FindByUsername(username string) (model.UserCredential, error) {
	var usercred model.UserCredential
	err := u.db.Where("user_name = ?", username).Find(&usercred).Error
	if err != nil {
		return usercred, err
	}
	return usercred, nil
}

func NewUserCredsRepository(db *gorm.DB) UserCredentialRepository {
	repo := new(userCredentialRepository)
	repo.db = db
	return repo
}
