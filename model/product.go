package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string      `gorm:"column:name;not null"`
	Customers   []*Customer `gorm:"many2many:customer_with_products"`
}

func (p Product) TableName() string {
	return "mst_product"
}

func (p *Product) ToString() string {
	product, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return ""
	}
	return string(product)
}
