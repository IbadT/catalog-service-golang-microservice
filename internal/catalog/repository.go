package catalog

import (
	domain "github.com/IbadT/catalog-service-golang-microservice.git/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	ListProdcuts(limit, offset int32) ([]domain.Product, error)
	GetProductById(id uuid.UUID) (domain.Product, error)
	SearchProductsByFilter() ([]domain.Product, error)
	CreateProduct(product domain.Product) error
	UploadProductImg(imgPath string) error

	CreateCategory(catalog domain.Category) error
	ListCategories() ([]domain.Category, error)
	UpdateProduct() error
	DeleteProduct(id uuid.UUID) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) ListProdcuts(limit, offset int32) ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) GetProductById(id uuid.UUID) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) SearchProductsByFilter() ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) CreateProduct(product domain.Product) error {
	return nil
}

func (r *repository) UploadProductImg(imgPath string) error {
	return nil
}

func (r *repository) CreateCategory(category domain.Category) error {
	return nil
}

func (r *repository) ListCategories() ([]domain.Category, error) {
	return nil, nil
}

func (r *repository) UpdateProduct() error {
	return nil
}

func (r *repository) DeleteProduct(id uuid.UUID) error {
	return nil
}
