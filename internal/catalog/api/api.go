package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/s-prophet/internal/catalog/service"
	catalogv1 "github.com/materials-resources/s-prophet/internal/grpc/catalog"
	"github.com/materials-resources/s-prophet/internal/grpc/catalog/catalogconnect"
)

func NewCatalogServiceHandler(svc service.CatalogService) *CatalogServiceHandler {
	return &CatalogServiceHandler{
		service: svc,
	}
}

type CatalogServiceHandler struct {
	service service.CatalogService
}

func (h CatalogServiceHandler) ListProducts(ctx context.Context, req *connect.Request[catalogv1.ListProductsRequest]) (*connect.Response[catalogv1.ListProductsResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) CreateProduct(ctx context.Context, req *connect.Request[catalogv1.CreateProductRequest]) (*connect.Response[catalogv1.CreateProductResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) GetProduct(ctx context.Context, req *connect.Request[catalogv1.GetProductRequest]) (*connect.Response[catalogv1.GetProductResponse], error) {
	res, err := h.service.GetProduct(ctx, req.Msg)

	if err != nil {
		return nil, err
	}

	return connect.NewResponse[catalogv1.GetProductResponse](res), nil
}

func (h CatalogServiceHandler) UpdateProduct(ctx context.Context, req *connect.Request[catalogv1.UpdateProductRequest]) (*connect.Response[catalogv1.UpdateProductResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) DeleteProduct(ctx context.Context, req *connect.Request[catalogv1.DeleteProductRequest]) (*connect.Response[catalogv1.DeleteProductResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) ListGroups(ctx context.Context, req *connect.Request[catalogv1.ListGroupRequest]) (*connect.Response[catalogv1.ListGroupResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) CreateGroup(ctx context.Context, req *connect.Request[catalogv1.CreateGroupRequest]) (*connect.Response[catalogv1.CreateGroupResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) GetGroup(ctx context.Context, req *connect.Request[catalogv1.GetGroupRequest]) (*connect.Response[catalogv1.GetGroupResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) UpdateGroup(ctx context.Context, req *connect.Request[catalogv1.UpdateGroupRequest]) (*connect.Response[catalogv1.UpdateGroupResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) DeleteGroup(ctx context.Context, req *connect.Request[catalogv1.DeleteGroupRequest]) (*connect.Response[catalogv1.DeleteGroupResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) ListSuppliers(ctx context.Context, req *connect.Request[catalogv1.ListSuppliersRequest]) (*connect.Response[catalogv1.ListSuppliersResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) GetSupplier(ctx context.Context, req *connect.Request[catalogv1.GetSupplierRequest]) (*connect.Response[catalogv1.GetSupplierResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) CreateSupplier(ctx context.Context, req *connect.Request[catalogv1.CreateSupplierRequest]) (*connect.Response[catalogv1.CreateSupplierResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) UpdateSupplier(ctx context.Context, req *connect.Request[catalogv1.UpdateSupplierRequest]) (*connect.Response[catalogv1.UpdateSupplierResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) SetPrimarySupplier(ctx context.Context, req *connect.Request[catalogv1.SetPrimarySupplierRequest]) (*connect.Response[catalogv1.SetPrimarySupplierResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h CatalogServiceHandler) GetProductBySupplier(ctx context.Context, req *connect.Request[catalogv1.GetProductBySupplierRequest]) (*connect.Response[catalogv1.GetProductBySupplierResponse], error) {
	//TODO implement me
	panic("implement me")
}

var _ catalogconnect.CatalogServiceHandler = (*CatalogServiceHandler)(nil)
