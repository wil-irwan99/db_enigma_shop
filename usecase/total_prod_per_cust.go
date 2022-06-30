package usecase

import (
	"fmt"
	"gorm/repository"
	"gorm/utils"
)

type ProductPerCustomer struct {
	custRepo repository.CustomerRepository
}

func (p *ProductPerCustomer) CountProductPerCustomer(id string) (int, error) {
	cust, err := p.custRepo.FindByIdWithPreload(map[string]interface{}{"id": id}, "Products")
	utils.IsError(err)
	if err != nil {
		return 0, err
	}

	var count int = 0
	for _, v := range cust.Products {
		fmt.Println(v.ProductName)
		count += 1
	}
	return count, nil
}

func NewProductPerCustomer(custRepo repository.CustomerRepository) ProductPerCustomer {
	uscs := new(ProductPerCustomer)
	uscs.custRepo = custRepo
	return *uscs
}
