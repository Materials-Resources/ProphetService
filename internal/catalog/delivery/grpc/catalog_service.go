package catalog

import (
	"context"
	"fmt"

	"github.com/materials-resources/s_prophet/internal/catalog"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type catalogService struct {
	repo   catalog.Repository
	tracer trace.Tracer
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
	ctx, span := s.tracer.Start(ctx, "DeleteProduct", trace.WithSpanKind(trace.SpanKindServer))
	span.SetAttributes(attribute.String("request.id", request.String()))
	defer span.End()

	err := s.repo.DeleteProduct(ctx, request.GetId())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
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
	_, span := s.tracer.Start(ctx, "GetGroup")
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
