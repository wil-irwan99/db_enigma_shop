package repository

import (
	"gorm/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func (c *productRepository) Create(product *model.Product) error {
	result := c.db.Create(product).Error
	return result
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
