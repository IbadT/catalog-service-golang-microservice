package catalog

import (
	domain "github.com/IbadT/catalog-service-golang-microservice.git/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	ListProdcuts(limit, offset int32) ([]domain.Product, error)
	GetProductById(id uuid.UUID) (domain.Product, error)
	SearchProductsByFilter(title, categoryId string, minPrice, maxPrice *float32) ([]domain.Product, error)
	CreateProduct(title, code, description, categoryId string, price float32, count int32) (domain.Product, error)

	UploadProductImg(img byte) error
	UpdateProductTitle(id uuid.UUID, title string) (domain.Product, error)
	UpdateProductPrice(id uuid.UUID, price int32) (domain.Product, error)
	DeleteProduct(id uuid.UUID) error

	CreateCategory(categoryName string) (domain.Category, error)

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
	dm, err := s.repo.GetProductById(id)
	if err != nil {
		return domain.Product{}, err
	}

	return dm, nil
}

func (s *service) SearchProductsByFilter(title, categoryId string, minPrice, maxPrice *float32) ([]domain.Product, error) {
	return s.repo.SearchProductsByFilter(title, categoryId, minPrice, maxPrice)
}

func (s *service) CreateProduct(title, code, description, categoryId string, price float32, count int32) (domain.Product, error) {
	dmProduct := domain.NewProduct(title, code, description, categoryId, price, count)

	err := s.repo.CreateProduct(dmProduct)
	if err != nil {
		return domain.Product{}, err
	}
	return *dmProduct, nil
}

func (s *service) UploadProductImg(img byte) error {
	return s.repo.UploadProductImg(string(img))
}

func (s *service) UpdateProductTitle(id uuid.UUID, title string) (domain.Product, error) {
	product, err := s.repo.GetProductById(id)
	if err != nil {
		return domain.Product{}, err
	}

	product.Title = title
	if err = s.repo.UpdateProduct(&product); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) UpdateProductPrice(id uuid.UUID, price int32) (domain.Product, error) {
	product, err := s.repo.GetProductById(id)
	if err != nil {
		return domain.Product{}, err
	}

	product.OldPrice = product.Price
	product.Price = float32(price)
	if err = s.repo.UpdateProduct(&product); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) DeleteProduct(id uuid.UUID) error {
	return s.repo.DeleteProduct(id)
}

func (s *service) CreateCategory(categoryName string) (domain.Category, error) {
	dmCategory := domain.NewCategory(categoryName)

	if err := s.repo.CreateCategory(dmCategory); err != nil {
		return domain.Category{}, err
	}
	return *dmCategory, nil
}

func (s *service) ListCategories() ([]domain.Category, error) {
	return s.repo.ListCategories()
}
