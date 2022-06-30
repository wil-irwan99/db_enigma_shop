package repository

import (
	"errors"
	"gorm/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindBy(by string, vals ...interface{}) ([]model.Product, error)
	FindById(id int) (model.Product, error)
	FindAllBy(by map[string]interface{}) ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Create(product *model.Product) error {
	result := p.db.Create(product).Error
	return result
}

func (p *productRepository) FindBy(by string, vals ...interface{}) ([]model.Product, error) {
	var product []model.Product
	result := p.db.Unscoped().Where(by, vals...).Find(&product)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}
	return product, nil
}

func (p *productRepository) FindById(id int) (model.Product, error) {
	var product model.Product
	result := p.db.First(&product, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}
	return product, nil
}

func (p *productRepository) FindAllBy(by map[string]interface{}) ([]model.Product, error) {
	var product []model.Product
	result := p.db.Where(by).Find(&product)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}
	return product, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
