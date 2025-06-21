package catalog

import (
	domain "github.com/IbadT/catalog-service-golang-microservice.git/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	ListProdcuts(limit, offset int32) ([]domain.Product, error)
	GetProductById(id uuid.UUID) (domain.Product, error)
	SearchProductsByFilter(title, categoryId string, minPrice, maxPrice *float32) ([]domain.Product, error)
	CreateProduct(product *domain.Product) error
	UploadProductImg(imgPath string) error

	CreateCategory(category *domain.Category) error
	ListCategories() ([]domain.Category, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uuid.UUID) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) ListProdcuts(limit, offset int32) ([]domain.Product, error) {
	var products []domain.Product
	err := r.DB.
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&products).Error

	return products, err
}

func (r *repository) GetProductById(id uuid.UUID) (domain.Product, error) {
	var product domain.Product
	if err := r.DB.First(&product, "id = ?", id).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) SearchProductsByFilter(title, categoryId string, minPrice, maxPrice *float32) ([]domain.Product, error) {
	var products []domain.Product
	query := r.DB.Model(&Product{})
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}
	if minPrice != nil {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice != nil {
		query = query.Where("price <= ?", maxPrice)
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) CreateProduct(product *domain.Product) error {
	return r.DB.Create(&product).Error
}

func (r *repository) UploadProductImg(imgPath string) error {
	return nil
}

func (r *repository) CreateCategory(category *domain.Category) error {
	return r.DB.Create(&category).Error
}

func (r *repository) ListCategories() ([]domain.Category, error) {
	var dmCategories []domain.Category
	err := r.DB.Find(&dmCategories).Error
	return dmCategories, err
}

func (r *repository) UpdateProduct(product *domain.Product) error {
	return nil
}

func (r *repository) DeleteProduct(id uuid.UUID) error {
	return r.DB.Delete(&Product{}, id).Error
}

// ////////

// Если ты хочешь сделать код ещё чище или переиспользуемым, можешь вынести фильтры в GORM scopes:

// func FilterByTitle(title string) func(*gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if title != "" {
// 			return db.Where("title ILIKE ?", "%"+title+"%")
// 		}
// 		return db
// 	}
// }

// func FilterByPrice(min, max *float32) func(*gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if min != nil {
// 			db = db.Where("price >= ?", *min)
// 		}
// 		if max != nil {
// 			db = db.Where("price <= ?", *max)
// 		}
// 		return db
// 	}
// }

// И использовать так:
// err := db.Model(&Product{}).
//     Scopes(
//         FilterByTitle(title),
//         FilterByPrice(minPrice, maxPrice),
//     ).
//     Find(&products).Error
