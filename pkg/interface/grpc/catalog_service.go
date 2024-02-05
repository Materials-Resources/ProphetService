package grpc

import (
	"context"
	"fmt"

	"github.com/materials-resources/s_prophet/pkg/domain/repositories"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
)

var tracer = otel.Tracer("catalog-service")

func NewCatalogService(server *grpc.Server, repo repositories.CatalogRepository) {
	svc := &catalogService{repo: repo}
	rpc.RegisterCatalogServiceServer(server, svc)
}

type catalogService struct {
	repo repositories.CatalogRepository
}

func (s catalogService) GetProduct(ctx context.Context, request *rpc.GetProductRequest) (*rpc.GetProductResponse, error,
) {
	fmt.Println(s.repo.FindProductByID(request.GetId()))
	return nil, nil
}

func (s catalogService) CreateProduct(
	ctx context.Context, request *rpc.CreateProductRequest,
) (*rpc.CreateProductResponse,
	error,
) {
	//TODO implement me
	panic("implement me")
}

func (s catalogService) DeleteProduct(
	ctx context.Context, request *rpc.DeleteProductRequest,
) (*rpc.DeleteProductResponse,
	error,
) {
	_, span := tracer.Start(ctx, "DeleteProduct")
	span.SetAttributes(attribute.String("request.id", request.String()))
	defer span.End()
	err := s.repo.DeleteProduct(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &rpc.DeleteProductResponse{}, nil
}

func (s catalogService) ListGroup(
	ctx context.Context, request *rpc.ListGroupRequest,
) (*rpc.ListGroupResponse, error) {
	productGroups, err := s.repo.ListGroup()

	if err != nil {
		return &rpc.ListGroupResponse{}, err
	}
	return ToPBListProductGroup(productGroups)
}

func (s catalogService) GetGroup(
	ctx context.Context, request *rpc.GetGroupRequest,
) (*rpc.GetGroupResponse, error) {
	_, span := tracer.Start(ctx, "GetGroup")
	span.SetAttributes(attribute.String("extra.key", "extra.value"))
	defer span.End()
	g, err := s.repo.FindGroupByID(ctx, request.GetId())
	if err != nil {
		fmt.Println(err)
	}
	p, err := s.repo.ReadProductByGroup(g.SN)
	if err != nil {
		fmt.Println(err)
	}
	return ToPBGetProductGroupResponse(g, p)
}

func (s catalogService) CreateGroup(
	ctx context.Context, request *rpc.CreateGroupRequest,
) (*rpc.CreateGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s catalogService) GetProductBySupplier(
	ctx context.Context, request *rpc.GetProductBySupplierRequest,
) (*rpc.GetProductBySupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}
