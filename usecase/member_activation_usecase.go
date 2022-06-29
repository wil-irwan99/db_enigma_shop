package usecase

import (
	"gorm/model"
	"gorm/repository"
	"log"
)

type MemberActivationUsecase struct {
	customerRepo repository.CustomerRepository
}

func (m *MemberActivationUsecase) MemberActivation(id string) (model.Customer, error) {
	customer, err := m.customerRepo.FindById(id)
	if err != nil {
		return model.Customer{}, err
	}

	customer.IsMember = "yes"
	err = m.customerRepo.UpdateBy(&customer)
	if err != nil {
		return model.Customer{}, err
	}
	log.Println("member activation success...")
	return customer, nil

}

func NewMemberActivitionUsecase(customerRepo repository.CustomerRepository) MemberActivationUsecase {
	usecase := new(MemberActivationUsecase)
	usecase.customerRepo = customerRepo
	return *usecase
}
