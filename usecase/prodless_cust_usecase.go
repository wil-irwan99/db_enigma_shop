package usecase

import (
	"gorm/model"
	"gorm/repository"
	"gorm/utils"
)

type ProdlessCust struct {
	custRepo repository.CustomerRepository
}

func (p *ProdlessCust) FindProdlesCust() ([]model.Customer, error) {
	result, err := p.custRepo.FindAll()
	utils.IsError(err)

	var prodlesCust []model.Customer
	for _, v := range result {
		res := p.custRepo.CountAssociation(&v, "Products")
		if res == 0 {
			prodlesCust = append(prodlesCust, v)
		}
	}
	return prodlesCust, nil
}

func NewProdlessCust(custRepo repository.CustomerRepository) ProdlessCust {
	uscs := new(ProdlessCust)
	uscs.custRepo = custRepo
	return *uscs
}
