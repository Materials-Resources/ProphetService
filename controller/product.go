package controller

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
	"github.com/materials-resources/s_prophet/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/product/v1alpha0"
	"google.golang.org/grpc"
)

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(server *grpc.Server, productRepository repository.ProductRepository) {
	productServer := &productService{productRepository: productRepository}
	rpc.RegisterProductServiceServer(
		server,
		productServer,
	)
}
func (s *productService) GetProduct(ctx context.Context, req *rpc.GetProductRequest) (*rpc.GetProductResponse, error) {
	if product, err := s.productRepository.Read(
		ctx,
		req.GetId(),
	); err != nil {
		return nil, err
	} else {
		response := domainToProductResponse(product)
		return response, nil
	}
}

func domainToProductResponse(product *domain.Product) *rpc.GetProductResponse {
	getProductResponse := &rpc.GetProductResponse{
		Id:          product.ID,
		Sn:          product.SN,
		Name:        product.Name,
		Description: product.Description,
	}
	return getProductResponse
}
