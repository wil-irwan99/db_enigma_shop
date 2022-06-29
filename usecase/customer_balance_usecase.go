package usecase

import (
	"errors"
	"gorm/model"
	"gorm/repository"
	"log"
)

type CustomerBalanceUsecase struct {
	customerRepo repository.CustomerRepository
}

func (c *CustomerBalanceUsecase) BalanceProcessing(id string, option string, nominal int) (model.Customer, error) {
	customer, err := c.customerRepo.FindById(id)
	if err != nil {
		return model.Customer{}, err
	}

	switch option {
	case "+":
		customer.Balance += nominal
	case "-":
		if customer.Balance >= nominal {
			customer.Balance -= nominal
		} else {
			return model.Customer{}, errors.New("saldo tidak cukup")
		}
	default:
		return model.Customer{}, errors.New("option salah")
	}

	err = c.customerRepo.UpdateBy(&customer)
	if err != nil {
		return model.Customer{}, err
	}
	log.Println("saldo processing success...")
	return customer, nil

}

func NewCustomerBalanceUsecase(customerRepo repository.CustomerRepository) CustomerBalanceUsecase {
	usecase := new(CustomerBalanceUsecase)
	usecase.customerRepo = customerRepo
	return *usecase
}
