package usecase

import (
	"errors"
	"gorm/model"
	"gorm/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationUsecase struct {
	custRepo     repository.CustomerRepository
	userCredRepo repository.UserCredentialRepository
}

func (a *AuthenticationUsecase) Authentication(uname string, pass string) (model.Customer, error) {

	result, err := a.userCredRepo.FindByUsername(uname)
	if err != nil {
		return model.Customer{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(pass))
	if err != nil {
		return model.Customer{}, errors.New("salah password")
	}

	customer, err := a.custRepo.FindByIdWithPreload(map[string]interface{}{"user_credential_id": result.ID}, "UserCredential")
	if err != nil {
		return model.Customer{}, err
	}

	customer.UserCredential.IsActive = true
	err = a.custRepo.UpdateBy(&customer)
	if err != nil {
		return model.Customer{}, err
	}
	log.Println("login success...")
	return customer, nil
}

func NewAuthenticationUsecase(custRepo repository.CustomerRepository, userCredRepo repository.UserCredentialRepository) AuthenticationUsecase {
	return AuthenticationUsecase{
		custRepo:     custRepo,
		userCredRepo: userCredRepo,
	}
}
