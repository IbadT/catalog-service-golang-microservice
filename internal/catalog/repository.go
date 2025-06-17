package catalog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	ListProdcuts(limit, offset int32) error
	GetProductById(id uuid.UUID) error
	SearchProductByFilter() error
	CreateProduct() error
	UploadProductImg() error

	CreateCatalog() error
	GetCatalog() error
	UpdateCatalog() error
	DeleteCatalog(id uuid.UUID) error
}

type repository struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) ListProdcuts(limit, offset int32) error {
	return nil
}

func (r *repository) GetProductById(id uuid.UUID) error {
	return nil
}

func (r *repository) SearchProductByFilter() error {
	return nil
}

func (r *repository) CreateProduct() error {
	return nil
}

func (r *repository) UploadProductImg() error {
	return nil
}

func (r *repository) CreateCatalog() error {
	return nil
}

func (r *repository) GetCatalog() error {
	return nil
}

func (r *repository) UpdateCatalog() error {
	return nil
}

func (r *repository) DeleteCatalog(id uuid.UUID) error {
	return nil
}
