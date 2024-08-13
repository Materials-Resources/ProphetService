package api

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rpc "github.com/materials-resources/microservices-proto/golang/catalog"
	"go.opentelemetry.io/otel/trace"
)

var _ rpc.CatalogServiceServer = CatalogApi{}

func NewCatalogApi(service service.CatalogService, tracer trace.Tracer,
) CatalogApi {
	return CatalogApi{service: service, tracer: tracer}
}

type CatalogApi struct {
	service service.CatalogService
	tracer  trace.Tracer
}

func (s CatalogApi) ListSuppliers(ctx context.Context, in *rpc.ListSuppliersRequest) (*rpc.ListSuppliersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) GetSupplier(ctx context.Context, in *rpc.GetSupplierRequest) (*rpc.GetSupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) CreateSupplier(ctx context.Context, in *rpc.CreateSupplierRequest) (*rpc.CreateSupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) UpdateSupplier(ctx context.Context, in *rpc.UpdateSupplierRequest) (*rpc.UpdateSupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) SetPrimarySupplier(ctx context.Context, in *rpc.SetPrimarySupplierRequest) (*rpc.SetPrimarySupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) GetProductBySupplier(ctx context.Context, in *rpc.GetProductBySupplierRequest) (*rpc.GetProductBySupplierResponse, error) {
	product, err := s.service.GetProductBySupplierPartNumber(ctx, in.GetSupplierPartNo(), in.GetSupplierId())
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			return nil, status.Error(codes.NotFound, "product not found")
		}

	}

	return &rpc.GetProductBySupplierResponse{
		Sn: *product.Sn,
	}, nil
}

func ConvertListProductsRequestToDomain(r *rpc.ListProductsRequest) *domain.Pagination {
	pagination := &domain.Pagination{
		Size: r.GetCursor().GetSize(),
	}

	switch r.GetCursor().GetDirection().(type) {
	case *rpc.Cursor_SelectPrevious:
		pagination.Direction = domain.CursorDirectionPrevious
		pagination.Cursor = int32(r.GetCursor().GetSelectPrevious())
	case *rpc.Cursor_SelectNext:
		pagination.Direction = domain.CursorDirectionNext
		pagination.Cursor = int32(r.GetCursor().GetSelectNext())
	}
	return pagination
}

func CreateListProductsResponse(products []*domain.Product, pagination *domain.PaginationMetadata) *rpc.ListProductsResponse {
	res := &rpc.ListProductsResponse{
		CursorMetadata: CreateCursorMetadata(pagination),
	}

	for _, p := range products {
		res.Products = append(res.Products, &rpc.ListProductsResponse_ProductDetails{
			Product:       CreateProduct(p),
			StockQuantity: float32(p.StockQuantity),
		})

	}
	return res
}

func (s CatalogApi) ListProducts(ctx context.Context, in *rpc.ListProductsRequest) (*rpc.ListProductsResponse, error) {
	pagination := ConvertListProductsRequestToDomain(in)

	products, paginationMetadata, err := s.service.ListProducts(ctx, *pagination)
	if err != nil {
		return nil, err

	}
	return CreateListProductsResponse(products, paginationMetadata), nil
}

func (s CatalogApi) CreateProduct(ctx context.Context, in *rpc.CreateProductRequest) (*rpc.CreateProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) GetProduct(ctx context.Context, in *rpc.GetProductRequest) (*rpc.GetProductResponse, error) {
	product, err := s.service.GetProduct(ctx, in.GetProductUid())
	if err != nil {
		return nil, err
	}
	return CreateGetProductResponse(product), nil
}

func ConvertUpdateProductRequestToDomain(r *rpc.UpdateProductRequest) *domain.Product {
	productDomain := &domain.Product{
		Uid: r.GetProductUid(),
	}

	paths := r.GetProduct().GetUpdateMask()

	for _, path := range paths.GetPaths() {
		if path == "name" {
			productDomain.Name = &r.GetProduct().Name
		}
		if path == "description" {
			productDomain.Description = &r.GetProduct().Description
		}
		if path == "product_group_sn" {
			productDomain.ProductGroupSn = &r.GetProduct().ProductGroupSn
		}
	}
	return productDomain
}

func (s CatalogApi) UpdateProduct(ctx context.Context, in *rpc.UpdateProductRequest) (*rpc.UpdateProductResponse, error) {
	err := s.service.UpdateProduct(ctx, ConvertUpdateProductRequestToDomain(in))
	if err != nil {
		return nil, err

	}
	return &rpc.UpdateProductResponse{}, nil
}

func (s CatalogApi) DeleteProduct(ctx context.Context, in *rpc.DeleteProductRequest) (*rpc.DeleteProductResponse, error) {
	err := s.service.DeleteProduct(ctx, in.GetProductUid())
	if err != nil {
		return nil, err

	}
	return &rpc.DeleteProductResponse{}, nil
}

func CreateListGroupsResponse(groups []*domain.ProductGroup) *rpc.ListGroupResponse {
	return &rpc.ListGroupResponse{
		ProductGroups: CreateProductGroupSlice(groups),
	}
}

func (s CatalogApi) ListGroups(ctx context.Context, in *rpc.ListGroupRequest) (*rpc.ListGroupResponse, error) {
	productGroup, err := s.service.ListGroups(ctx)
	if err != nil {
		return nil, err

	}

	return CreateListGroupsResponse(productGroup), nil

}

func ConvertCreateGroupRequestToDomain(r *rpc.CreateGroupRequest) *domain.ProductGroup {
	return &domain.ProductGroup{
		Sn:   r.GetSn(),
		Name: &r.Name,
	}
}

func (s CatalogApi) CreateGroup(ctx context.Context, in *rpc.CreateGroupRequest) (*rpc.CreateGroupResponse, error) {
	productGroup := ConvertCreateGroupRequestToDomain(in)
	err := s.service.CreateGroup(ctx, productGroup)
	if err != nil {
		return nil, err

	}
	return &rpc.CreateGroupResponse{}, nil
}

func CreateGetGroupResponse(group *domain.ProductGroup) *rpc.GetGroupResponse {
	return &rpc.GetGroupResponse{
		ProductGroup: &rpc.ProductGroup{
			Sn:   group.Sn,
			Name: *group.Name,
		},
	}
}

func (s CatalogApi) GetGroup(ctx context.Context, in *rpc.GetGroupRequest) (*rpc.GetGroupResponse, error) {
	productGroup, err := s.service.ReadGroup(ctx, in.GetSn())
	if err != nil {
		return nil, err

	}

	return CreateGetGroupResponse(productGroup), nil
}

func ConvertUpdateGroupRequestToDomain(r *rpc.UpdateGroupRequest) *domain.ProductGroup {
	productGroupDomain := &domain.ProductGroup{
		Sn: r.Sn,
	}
	paths := r.GetProductGroup().GetUpdateMask()

	for _, path := range paths.GetPaths() {
		if path == "name" {
			productGroupDomain.Name = &r.ProductGroup.Name
		}
	}
	return productGroupDomain
}

func (s CatalogApi) UpdateGroup(ctx context.Context, in *rpc.UpdateGroupRequest) (*rpc.UpdateGroupResponse, error) {
	productGroup := ConvertUpdateGroupRequestToDomain(in)
	err := s.service.UpdateGroup(ctx, productGroup)
	if err != nil {
		return nil, err

	}
	return &rpc.UpdateGroupResponse{}, nil
}

func (s CatalogApi) DeleteGroup(ctx context.Context, in *rpc.DeleteGroupRequest) (*rpc.DeleteGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func CreateProduct(p *domain.Product) *rpc.Product {

	r := &rpc.Product{
		Uid:              p.Uid,
		ProductGroupName: p.ProductGroupName,
	}

	if p.Sn != nil {
		r.Sn = *p.Sn
	}

	if p.Name != nil {
		r.Name = *p.Name
	}

	if p.Description != nil {
		r.Description = *p.Description

	}
	if p.ProductGroupSn != nil {
		r.ProductGroupSn = *p.ProductGroupSn

	}
	return r
}

func CreateGetProductResponse(p *domain.Product) *rpc.GetProductResponse {
	return &rpc.GetProductResponse{
		Product: CreateProduct(p),
	}
}

func CreateProductGroupSlice(groups []*domain.ProductGroup) []*rpc.ProductGroup {
	var result []*rpc.ProductGroup
	for _, g := range groups {
		result = append(result, &rpc.ProductGroup{
			Sn:   g.Sn,
			Name: *g.Name,
		})
	}
	return result
}

func CreateCursorMetadata(p *domain.PaginationMetadata) *rpc.CursorMetadata {
	return &rpc.CursorMetadata{
		PreviousCursor: int64(p.PreviousCursor),
		NextCursor:     int64(p.NextCursor),
		Total:          p.Total,
	}

}
