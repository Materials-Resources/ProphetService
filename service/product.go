package service

import (
	"context"
	"fmt"

	"github.com/materials-resources/s_prophet/core/domain"
	"github.com/materials-resources/s_prophet/core/port/repository"
	pb "github.com/materials-resources/s_prophet/proto/product/v1alpha0"
	"google.golang.org/grpc"
)

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductServer(server *grpc.Server, productRepository repository.ProductRepository) {
	productServer := &productService{productRepository: productRepository}
	pb.RegisterProductServiceServer(
		server,
		productServer,
	)
}

func (s *productService) CreateProduct(
	ctx context.Context, request *pb.CreateProductRequest,
) (*pb.CreateProductResponse, error) {
	s.productRepository.CreateProduct(ctx)
	return nil, nil
}
func (s *productService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	if product, err := s.productRepository.ReadProduct(
		ctx,
		req.GetId(),
	); err != nil {
		return nil, err
	} else {
		response := domainToProductResponse(product)
		return response, nil
	}
}

func (s *productService) GetProductBySupplier(
	ctx context.Context, request *pb.GetProductBySupplierRequest,
) (*pb.GetProductBySupplierResponse, error) {
	if sn, err := s.productRepository.ReadProductBySupplier(
		ctx,
		request.GetSupplierId(),
		request.GetSupplierPartNo(),
	); err != nil {
		return nil, err
	} else {
		response := pb.GetProductBySupplierResponse{Sn: *sn}
		return &response, nil
	}
}

func (s *productService) CreateProductGroup(ctx context.Context, req *pb.CreateProductGroupRequest) (*pb.
	CreateProductGroupResponse, error) {
	newProductGroup := domain.ProductGroup{
		Name: req.GetName(),
		SN:   req.GetSn(),
	}
	err := s.productRepository.CreateProductGroup(
		ctx,
		&newProductGroup,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil, nil

}

func domainToProductResponse(product *domain.Product) *pb.GetProductResponse {
	getProductResponse := &pb.GetProductResponse{
		Id:             product.ID,
		Sn:             product.SN,
		Name:           product.Name,
		Description:    product.Description,
		Price:          product.Price,
		UnitsAvailable: product.UnitsAvailable,
	}
	return getProductResponse
}
