package api

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
	"github.com/materials-resources/s-prophet/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rpc "github.com/materials-resources/microservices-proto/golang/catalog"
	"go.opentelemetry.io/otel/trace"
)

func NewCatalogApi(tracer trace.Tracer, producer service.Producer,
) CatalogApi {
	return CatalogApi{tracer: tracer, producer: producer}
}

type CatalogApi struct {
	service  service.CatalogService
	tracer   trace.Tracer
	producer service.Producer
}

func (s CatalogApi) CreateSupplier(ctx context.Context, request *rpc.CreateSupplierRequest) (*rpc.CreateSupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) UpdateSupplier(ctx context.Context, request *rpc.UpdateSupplierRequest) (*rpc.UpdateSupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) SetPrimarySupplier(ctx context.Context, request *rpc.SetPrimarySupplierRequest) (*rpc.SetPrimarySupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) CustomerGetProduct(ctx context.Context, request *rpc.CustomerGetProductRequest) (*rpc.CustomerGetProductResponse, error) {
	product, err := s.service.CustomerGetProduct(ctx, request.GetProductUid())
	if err != nil {
		return nil, err
	}
	return CreateCustomerGetProductResponse(product), nil
}

func (s CatalogApi) ClerkListProducts(ctx context.Context, request *rpc.ClerkListProductsRequest) (*rpc.ClerkListProductsResponse, error) {
	pagination := ConvertClerkListProductsRequestToDomain(request)

	products, paginationMetadata, err := s.service.ClerkListProducts(ctx, *pagination)
	if err != nil {
		return nil, err

	}
	return CreateClerkListProductsResponse(products, paginationMetadata), nil
}

func (s CatalogApi) ClerkCreateProduct(ctx context.Context, request *rpc.ClerkCreateProductRequest) (*rpc.ClerkCreateProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) ClerkReadProduct(ctx context.Context, request *rpc.ClerkReadProductRequest) (*rpc.ClerkReadProductResponse, error) {
	product, err := s.service.ClerkGetProduct(ctx, request.GetProductUid())
	if err != nil {
		return nil, err
	}
	return CreateClerkReadProductResponse(product), nil
}

func (s CatalogApi) ClerkUpdateProduct(ctx context.Context, request *rpc.ClerkUpdateProductRequest) (*rpc.ClerkUpdateProductResponse, error) {
	err := s.service.ClerkUpdateProduct(ctx, ConvertClerkUpdateProductRequestToDomain(request))
	if err != nil {
		return nil, err

	}
	return &rpc.ClerkUpdateProductResponse{}, nil
}

func (s CatalogApi) ClerkDeleteProduct(ctx context.Context, request *rpc.ClerkDeleteProductRequest) (*rpc.ClerkDeleteProductResponse, error) {
	request.DeleteMode = rpc.DeleteMode_DELETE_MODE_SOFT
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) ClerkListGroups(ctx context.Context, request *rpc.ClerkListGroupRequest) (*rpc.ClerkListGroupResponse, error) {
	productGroup, err := s.service.ClerkListGroups(ctx)
	if err != nil {
		return nil, err

	}

	return CreateClerkListGroupsResponse(productGroup), nil

}

func (s CatalogApi) ClerkCreateGroup(ctx context.Context, request *rpc.ClerkCreateGroupRequest) (*rpc.ClerkCreateGroupResponse, error) {
	productGroup := ConvertClerkCreateGroupRequestToDomain(request)
	err := s.service.ClerkCreateGroup(ctx, productGroup)
	if err != nil {
		return nil, err

	}
	return &rpc.ClerkCreateGroupResponse{}, nil
}

func (s CatalogApi) ClerkReadGroup(ctx context.Context, request *rpc.ClerkReadGroupRequest) (*rpc.ClerkReadGroupResponse, error) {
	productGroup, err := s.service.ClerkReadGroup(ctx, request.GetSn())
	if err != nil {
		return nil, err

	}

	return CreateClerkReadGroupResponse(productGroup), nil
}

func (s CatalogApi) ClerkUpdateGroup(ctx context.Context, request *rpc.ClerkUpdateGroupRequest) (*rpc.ClerkUpdateGroupResponse, error) {
	productGroup := ConvertClerkUpdateGroupRequestToDomain(request)
	err := s.service.ClerkUpdateGroup(ctx, productGroup)
	if err != nil {
		return nil, err

	}
	return &rpc.ClerkUpdateGroupResponse{}, nil
}

func (s CatalogApi) ClerkDeleteGroup(ctx context.Context, request *rpc.ClerkDeleteGroupRequest) (*rpc.ClerkDeleteGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) ListSuppliers(ctx context.Context, request *rpc.ListSuppliersRequest) (
	*rpc.ListSuppliersResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s CatalogApi) GetSupplier(
	ctx context.Context,
	request *rpc.GetSupplierRequest,
) (*rpc.GetSupplierResponse, error) {
	//TODO implement me
	return nil, nil
}

func (s CatalogApi) GetProductBySupplier(
	ctx context.Context, request *rpc.GetProductBySupplierRequest,
) (*rpc.GetProductBySupplierResponse, error) {
	product, err := s.service.GetProductBySupplierPartNumber(ctx, request.GetSupplierPartNo(), request.GetSupplierId())
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
