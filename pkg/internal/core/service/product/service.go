package product

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/product/v1alpha0"
	"google.golang.org/grpc"
)

type service struct {
	productRepository repository.ProductRepository
}

func NewProductServer(server *grpc.Server, productRepository repository.ProductRepository) {
	productServer := &service{productRepository: productRepository}
	rpc.RegisterProductServiceServer(
		server,
		productServer,
	)
}
func (s *service) GetProduct(ctx context.Context, req *rpc.GetProductRequest) (*rpc.GetProductResponse, error) {
	if product, err := s.productRepository.Read(req.GetId()); err != nil {
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
