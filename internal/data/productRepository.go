package data

import (
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() []*domain.Product
	GetProductBy(id uuid.UUID) *domain.Product
	AddProduct(u *domain.Product) *domain.Product
	UpdateProduct(u *domain.Product) *domain.Product
	DeleteProduct(id uuid.UUID) bool
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (pr *productRepository) GetProducts() []*domain.Product {
	var products []*domain.Product
	pr.db.Find(&products)
	return products
}

func (pr *productRepository) GetProductBy(id uuid.UUID) *domain.Product {
	p := &domain.Product{}
	result := pr.db.First(p, id)

	if result.RowsAffected > 0 {
		return p
	}

	return nil
}

func (pr *productRepository) AddProduct(p *domain.Product) *domain.Product {
	pr.db.Create(p)
	return p
}

func (pr *productRepository) UpdateProduct(p *domain.Product) *domain.Product {
	pr.db.Save(p)
	return p
}

func (pr *productRepository) DeleteProduct(id uuid.UUID) bool {
	result := pr.db.Delete(&domain.Product{}, id)

	return result.RowsAffected > 0
}
