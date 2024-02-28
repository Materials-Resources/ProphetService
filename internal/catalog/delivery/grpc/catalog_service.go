package catalog

import (
	"context"
	"fmt"

	"github.com/materials-resources/s_prophet/internal/catalog"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	"github.com/materials-resources/s_prophet/internal/catalog/kafka"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type catalogService struct {
	repo     catalog.Repository
	tracer   trace.Tracer
	producer kafka.Producer
}

func (s catalogService) UpdateProduct(
	ctx context.Context,
	request *rpc.UpdateProductRequest) (*rpc.UpdateProductResponse, error) {

	return nil, nil
}

func (s catalogService) CreateProductSupplier(
	ctx context.Context,
	request *rpc.CreateProductSupplierRequest) (*rpc.CreateProductSupplierResponse, error) {
	d := &domain.ProductSupplier{SupplierId: request.GetSupplierId(), SupplierSn: request.GetSupplierProductSn(),
		ProductId: request.GetProductId(), ListPrice: request.GetListPrice(), PurchasePrice: request.GetPurchasePrice(),
	}
	vd, err := domain.NewValidatedProductSupplier(d)
	if err != nil {
		return nil, err
	}
	s.repo.AddProductSupplier(ctx, vd)
	return nil, nil
}

func (s catalogService) UpdateProductSupplier(
	ctx context.Context,
	request *rpc.UpdateProductSupplierRequest) (*rpc.UpdateProductSupplierResponse, error) {
	p := &domain.ProductSupplier{ProductId: request.GetProductSupplier().GetProductId(),
		SupplierId: request.GetProductSupplier().SupplierId, Delete: request.GetProductSupplier().GetDelete(),
	}

	vp, err := domain.NewValidatedProductSupplier(p)
	if err != nil {
		return nil, err
	}
	if err := s.repo.UpdateProductSupplier(ctx, vp); err != nil {
		return nil, err
	}
	return &rpc.UpdateProductSupplierResponse{}, nil
}

func (s catalogService) SetPrimaryProductSupplier(
	ctx context.Context,
	request *rpc.SetPrimaryProductSupplierRequest) (*rpc.SetPrimaryProductSupplierResponse, error) {
	err := s.repo.SetPrimaryProductSupplier(ctx, request.GetProductId(), request.GetSupplierId())
	if err != nil {
		return nil, err
	}
	return &rpc.SetPrimaryProductSupplierResponse{}, nil
}

func (s catalogService) ListProduct(ctx context.Context, request *rpc.ListProductRequest) (*rpc.ListProductResponse,
	error,
) {
	res, nc, err := s.repo.ListProduct(ctx, request.GetCursor(), 100)
	if err != nil {
		return nil, err
	}
	return ToPBListProductResponse(res, int32(nc))
}

func (s catalogService) UpdateGroup(ctx context.Context, request *rpc.UpdateGroupRequest) (*rpc.UpdateGroupResponse,
	error,
) {
	d := &domain.ProductGroup{SN: request.GetProductGroup().GetSn(), Name: request.GetProductGroup().GetName()}

	err := s.repo.UpdateGroup(ctx, d)
	if err != nil {
		return nil, err
	}
	return &rpc.UpdateGroupResponse{}, nil
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

	s.producer.PublishDelete(ctx, request.GetId())

	//err := s.repo.DeleteProduct(ctx, request.GetId())
	//if err != nil {
	//	span.RecordError(err)
	//	span.SetStatus(codes.Error, err.Error())
	//	return nil, err
	//}
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
	_, span := s.tracer.Start(ctx, "CreateGroup")
	defer span.End()
	d := domain.NewProductGroup(request.GetProductGroup().GetName(), request.GetProductGroup().GetSn())
	vd, err := domain.NewValidatedProductGroup(d)
	if err != nil {
		return nil, err
	}
	s.repo.CreateGroup(ctx, vd)
	return nil, nil
}

func (s catalogService) GetProductBySupplier(
	ctx context.Context, request *rpc.GetProductBySupplierRequest,
) (*rpc.GetProductBySupplierResponse, error) {
	//TODO implement me
	panic("implement me")
}
