package transportgrpc

import (
	"context"

	"github.com/IbadT/catalog-service-golang-microservice.git/internal/catalog"
	productpb "github.com/IbadT/project-protos/proto/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	svc catalog.Service
	productpb.UnimplementedProductServiceServer
}

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
	panic("Not Implemented")
}

func (h *handler) GetProductByID(ctx context.Context, req *productpb.GetProductByIDRequest) (*productpb.GetCartByIdResponse, error) {
	panic("Not Implemented")
}

func (h *handler) SearchProductsByFilter(ctx context.Context, req *productpb.SearchProductByFilterRequest) (*productpb.SearchProductByFilterResponse, error) {
	panic("Not Implemented")
}

func (h *handler) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	panic("Not Implemented")
}

func (h *handler) UploadProductImg(ctx context.Context, req *productpb.UploadProductImgRequest) (*productpb.UploadProductImgResponse, error) {
	panic("Not Implemented")
}

func (h *handler) UpdateProductTitle(ctx context.Context, req *productpb.UpdateProductTitleRequest) (*productpb.UpdateProductTitleResponse, error) {
	panic("Not Implemented")
}

func (h *handler) UpdateProductPrice(ctx context.Context, req *productpb.UpdateProductPriceRequest) (*productpb.UpdateProductPriceResponse, error) {
	panic("Not Implemented")
}

func (h *handler) DeleteProduct(ctx context.Context, req *productpb.DeleteCatalogRequest) (*emptypb.Empty, error) {
	panic("Not Implemented")
}

// category
func (h *handler) CreateCategory(ctx context.Context, req *productpb.CreateCategoryRequest) (*productpb.CreateCategoryResponse, error) {
	panic("Not Implemented")
}

func (h *handler) ListCategories(ctx context.Context, req *emptypb.Empty) (*productpb.ListCategoriesResponse, error) {
	panic("Not Implemented")
}
