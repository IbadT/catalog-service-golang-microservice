package catalog

import (
	domain "github.com/IbadT/catalog-service-golang-microservice.git/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	ListProdcuts(limit, offset int32) ([]domain.Product, error)
	GetProductById(id uuid.UUID) (domain.Product, error)
	SearchProductsByFilter(title, categoryId string, minPrice, maxPrice int32) ([]domain.Product, error)
	CreateProduct(product domain.Product) (domain.Product, error) // не используется domain.Product в возврате
	UploadProductImg(img byte) error
	UpdateProductTitle(id uuid.UUID, title string) (domain.Product, error)
	UpdateProductPrice(id uuid.UUID, price int32) (domain.Product, error)
	DeleteProduct(id uuid.UUID) error

	CreateCategory(category domain.Category) (domain.Category, error)
	ListCategories() ([]domain.Category, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListProdcuts(limit, offset int32) ([]domain.Product, error) {
	return s.repo.ListProdcuts(limit, offset)
}

func (s *service) GetProductById(id uuid.UUID) (domain.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *service) SearchProductsByFilter(title, categoryId string, minPrice, maxPrice int32) ([]domain.Product, error) {
	return s.repo.SearchProductsByFilter()
}

func (s *service) CreateProduct(product domain.Product) (domain.Product, error) {
	return product, s.repo.CreateProduct(product)
}

func (s *service) UploadProductImg(img byte) error {
	return s.repo.UploadProductImg(string(img))
}

func (s *service) UpdateProductTitle(id uuid.UUID, title string) (domain.Product, error) {
	return domain.Product{}, s.repo.UpdateProduct()
}

func (s *service) UpdateProductPrice(id uuid.UUID, price int32) (domain.Product, error) {
	return domain.Product{}, s.repo.UpdateProduct()
}

func (s *service) DeleteProduct(id uuid.UUID) error {
	return s.repo.DeleteProduct(id)
}

func (s *service) CreateCategory(category domain.Category) (domain.Category, error) {
	return category, s.repo.CreateCategory(category)
}

func (s *service) ListCategories() ([]domain.Category, error) {
	return s.repo.ListCategories()
}
