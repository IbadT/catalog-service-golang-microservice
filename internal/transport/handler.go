package transportgrpc

import (
	"context"

	"github.com/IbadT/catalog-service-golang-microservice.git/internal/catalog"
	"github.com/IbadT/catalog-service-golang-microservice.git/internal/domain"
	productpb "github.com/IbadT/project-protos/proto/product"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	svc catalog.Service
	productpb.UnimplementedProductServiceServer
}

// !!!!!! в сервис отправлять сырые данные, а не сразу domain !!!!!!

type Handler interface {
	ListProducts(ctx context.Context, req *productpb.ListProductsRequest) (*productpb.ListProductsResponse, error)
	GetProductByID(ctx context.Context, req *productpb.GetProductByIDRequest) (*productpb.GetCartByIdResponse, error)
	SearchProductsByFilter(ctx context.Context, req *productpb.SearchProductByFilterRequest) (*productpb.SearchProductByFilterResponse, error)
	CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error)
	UploadProductImg(ctx context.Context, req *productpb.UploadProductImgRequest) (*productpb.UploadProductImgResponse, error)
	UpdateProductTitle(ctx context.Context, req *productpb.UpdateProductTitleRequest) (*productpb.UpdateProductTitleResponse, error)
	UpdateProductPrice(ctx context.Context, req *productpb.UpdateProductPriceRequest) (*productpb.UpdateProductPriceResponse, error)
	DeleteProduct(ctx context.Context, req *productpb.DeleteCatalogRequest) (*emptypb.Empty, error)
	CreateCategory(ctx context.Context, req *productpb.CreateCategoryRequest) (*productpb.CreateCategoryResponse, error)
	ListCategories(ctx context.Context, req *emptypb.Empty) (*productpb.ListCategoriesResponse, error)
}

func NewHandler(s catalog.Service) *handler {
	return &handler{svc: s}
}

func (h *handler) ListProducts(ctx context.Context, req *productpb.ListProductsRequest) (*productpb.ListProductsResponse, error) {
	limit := req.Limit
	offset := req.Offset

	products, err := h.svc.ListProdcuts(limit, offset)
	if err != nil {
		return &productpb.ListProductsResponse{}, err
	}

	productspb := make([]*productpb.Product, 0, len(products))
	for _, prod := range products {
		p := &productpb.Product{
			Id:          prod.ID.String(),
			Title:       prod.Title,
			Price:       float32(prod.Price),
			OldPrice:    float32(prod.OldPrice),
			Description: prod.Description,
			Rating:      prod.Rating,
			ImageUrl:    prod.ImageUrl,
			Mimetype:    prod.Mimetype,
			Count:       int32(prod.Count),
			Code:        prod.Code,
			CategoryId:  prod.CategoryId,
		}
		productspb = append(productspb, p)
	}

	return &productpb.ListProductsResponse{Products: productspb}, nil
}

func (h *handler) GetProductByID(ctx context.Context, req *productpb.GetProductByIDRequest) (*productpb.GetCartByIdResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &productpb.GetCartByIdResponse{}, err
	}

	product, err := h.svc.GetProductById(id)
	if err != nil {
		return &productpb.GetCartByIdResponse{}, err
	}

	prodpb := &productpb.Product{
		Id:          product.ID.String(),
		Title:       product.Title,
		Price:       float32(product.Price),
		OldPrice:    float32(product.OldPrice),
		Description: product.Description,
		Rating:      product.Rating,
		ImageUrl:    product.ImageUrl,
		Mimetype:    product.Mimetype,
		Count:       product.Count,
		Code:        product.Code,
		CategoryId:  product.CategoryId,
	}

	return &productpb.GetCartByIdResponse{Product: prodpb}, nil
}

func (h *handler) SearchProductsByFilter(ctx context.Context, req *productpb.SearchProductByFilterRequest) (*productpb.SearchProductByFilterResponse, error) {
	title := ""
	if req.Title != nil {
		title = req.Title.Value
	}

	var minPrice *int32
	if req.MinPrice != nil {
		minPrice = &req.MinPrice.Value
	}

	var maxPrice *int32
	if req.MaxPrice != nil {
		maxPrice = &req.MaxPrice.Value
	}

	categoryId := ""
	if req.CategoryId != nil {
		categoryId = req.CategoryId.Value
	}

	products, err := h.svc.SearchProductsByFilter(title, categoryId, *minPrice, *maxPrice)
	if err != nil {
		return &productpb.SearchProductByFilterResponse{}, err
	}

	prodspb := make([]*productpb.Product, 0, len(products))
	for _, prod := range products {
		p := &productpb.Product{
			Id:          prod.ID.String(),
			Title:       prod.Title,
			Price:       float32(prod.Price),
			Description: prod.Description,
			Rating:      prod.Rating,
			ImageUrl:    prod.ImageUrl,
			Mimetype:    prod.Mimetype,
			OldPrice:    float32(prod.OldPrice),
			Count:       prod.Count,
			Code:        prod.Code,
			CategoryId:  prod.CategoryId,
		}
		prodspb = append(prodspb, p)
	}

	return &productpb.SearchProductByFilterResponse{Products: prodspb}, nil
}

func (h *handler) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	title := req.Title
	price := int32(req.Price)
	count := int32(req.Count)
	code := string(req.Code)
	description := req.Description
	categoryId := req.CategoryId

	prod := domain.Product{
		ID:          uuid.New(),
		Title:       title,
		Price:       price,
		Count:       count,
		Code:        code,
		Description: description,
		CategoryId:  categoryId,
	}

	_, err := h.svc.CreateProduct(prod)
	if err != nil {
		return &productpb.CreateProductResponse{}, nil
	}

	prodpb := &productpb.Product{
		Id:          prod.ID.String(),
		Title:       prod.Title,
		Price:       float32(prod.Price),
		OldPrice:    float32(prod.OldPrice),
		Description: prod.Description,
		Rating:      prod.Rating,
		ImageUrl:    prod.ImageUrl,
		Mimetype:    prod.Mimetype,
		Count:       int32(prod.Count),
		Code:        prod.Code,
		CategoryId:  prod.CategoryId,
	}

	return &productpb.CreateProductResponse{Product: prodpb}, nil

}

func (h *handler) UploadProductImg(ctx context.Context, req *productpb.UploadProductImgRequest) (*productpb.UploadProductImgResponse, error) {
	panic("Not Implemented")
}

func (h *handler) UpdateProductTitle(ctx context.Context, req *productpb.UpdateProductTitleRequest) (*productpb.UpdateProductTitleResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &productpb.UpdateProductTitleResponse{}, err
	}
	title := req.Title

	prod, err := h.svc.UpdateProductTitle(id, title)

	prodpb := &productpb.Product{
		Id:          prod.ID.String(),
		Title:       prod.Title,
		Price:       float32(prod.Price),
		OldPrice:    float32(prod.OldPrice),
		Description: prod.Description,
		Rating:      prod.Rating,
		ImageUrl:    prod.ImageUrl,
		Mimetype:    prod.Mimetype,
		Count:       int32(prod.Count),
		Code:        prod.Code,
		CategoryId:  prod.CategoryId,
	}

	return &productpb.UpdateProductTitleResponse{Product: prodpb}, nil
}

func (h *handler) UpdateProductPrice(ctx context.Context, req *productpb.UpdateProductPriceRequest) (*productpb.UpdateProductPriceResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &productpb.UpdateProductPriceResponse{}, err
	}
	price := int32(req.Price)

	prod, err := h.svc.UpdateProductPrice(id, price)
	if err != nil {
		return &productpb.UpdateProductPriceResponse{}, err
	}

	prodpb := &productpb.Product{
		Id:          prod.ID.String(),
		Title:       prod.Title,
		Price:       float32(prod.Price),
		OldPrice:    float32(prod.OldPrice),
		Description: prod.Description,
		Rating:      prod.Rating,
		ImageUrl:    prod.ImageUrl,
		Mimetype:    prod.Mimetype,
		Count:       int32(prod.Count),
		Code:        prod.Code,
		CategoryId:  prod.CategoryId,
	}

	return &productpb.UpdateProductPriceResponse{Product: prodpb}, nil
}

func (h *handler) DeleteProduct(ctx context.Context, req *productpb.DeleteCatalogRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err = h.svc.DeleteProduct(id); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

// category
func (h *handler) CreateCategory(ctx context.Context, req *productpb.CreateCategoryRequest) (*productpb.CreateCategoryResponse, error) {
	categoryName := req.Category.CategoryName

	category := domain.Category{
		ID:           uuid.New(),
		CategoryName: categoryName,
	}

	c, err := h.svc.CreateCategory(category)
	if err != nil {
		return &productpb.CreateCategoryResponse{}, err
	}

	catpb := &productpb.Category{
		Id:           c.ID.String(),
		CategoryName: c.CategoryName,
	}

	return &productpb.CreateCategoryResponse{Category: catpb}, nil
}

func (h *handler) ListCategories(ctx context.Context, req *emptypb.Empty) (*productpb.ListCategoriesResponse, error) {
	categories, err := h.svc.ListCategories()
	if err != nil {
		return &productpb.ListCategoriesResponse{}, err
	}

	catspb := make([]*productpb.Category, 0, len(categories))
	for _, c := range categories {
		category := &productpb.Category{
			Id:           c.ID.String(),
			CategoryName: c.CategoryName,
		}
		catspb = append(catspb, category)
	}

	return &productpb.ListCategoriesResponse{Categories: catspb}, nil
}
