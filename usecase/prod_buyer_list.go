package usecase

import (
	"gorm/repository"
	"gorm/utils"
)

type ProductBuyerList struct {
	custRepo repository.CustomerRepository
}

func (p *ProductBuyerList) FindProductBuyerList() (interface{}, error) {
	cust, err := p.custRepo.FindAllWithPreload("Products")
	utils.IsError(err)
	if err != nil {
		return nil, err
	}

	var cacaMarica []string

	for _, v := range cust { //bermasalah di nested loop
		for _, k := range v.Products {
			switch k.ID {
			case 8:
				cacaMarica = append(cacaMarica, v.Name)
			}
		}
	}
	result := map[string][]string{"Caca Marica": cacaMarica}
	return result, nil
}

func NewProductBuyerList(custRepo repository.CustomerRepository) ProductBuyerList {
	uscs := new(ProductBuyerList)
	uscs.custRepo = custRepo
	return *uscs
}
