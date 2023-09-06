package grpc_service

import (
	"context"
	"github.com/materials-resources/s_prophet/pkg/database"
	rpc "github.com/materials-resources/s_prophet/proto/product/v1alpha0"
)

type ProductServer struct {
	rpc.UnimplementedProductServiceServer
	ProductHandler *database.ProductHandler
}

func (s *ProductServer) GetProduct(ctx context.Context, req *rpc.GetProductRequest) (*rpc.GetProductResponse, error) {
	grpcRes := new(rpc.GetProductResponse)
	hRes, err := s.ProductHandler.SelectProduct(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}

	for _, product := range *hRes {
		grpcRes.Products = append(grpcRes.Products, &rpc.GetProductResponse_Product{
			Id:             product.InvMastUid,
			ProductSn:      product.ItemId,
			Name:           product.ItemDesc,
			Description:    product.ExtendedDesc,
			ProductGroupId: product.DefaultProductGroup,
			Price:          product.Price1,
		})
	}

	return grpcRes, nil
}

func (s *ProductServer) GetGroupProducts(ctx context.Context, req *rpc.GetGroupProductsRequest) (*rpc.GetGroupProductsResponse, error) {
	grpcRes := new(rpc.GetGroupProductsResponse)

	qr, err := s.ProductHandler.SelectProductsInProductGroup(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	for _, r := range *qr {
		grpcRes.ProductId = append(grpcRes.ProductId, r.InvMastUid)
	}

	return grpcRes, nil

}
func (s *ProductServer) GetProductGroups(ctx context.Context, req *rpc.GetProductGroupsRequest) (*rpc.GetProductGroupsResponse, error) {
	grpcResponse := new(rpc.GetProductGroupsResponse)

	qr, err := s.ProductHandler.SelectProductGroups(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range *qr {
		grpcResponse.ProductGroups = append(grpcResponse.ProductGroups, &rpc.GetProductGroupsResponse_ProductGroup{
			Id:   r.ProductGroupId,
			Name: r.ProductGroupDesc,
		})

	}

	return grpcResponse, nil
}
func (s *ProductServer) CreateProductGroup(ctx context.Context, req *rpc.CreateProductGroupRequest) (*rpc.CreateProductGroupResponse, error) {

	return nil, nil
}
