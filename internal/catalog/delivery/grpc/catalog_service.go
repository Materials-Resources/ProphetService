package catalog

import (
	"context"
	"errors"
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

func (s catalogService) GetProductPrice(
	ctx context.Context,
	request *rpc.GetProductPriceRequest) (*rpc.GetProductPriceResponse, error) {
	dPP, err := s.repo.SelectProductPrice(ctx, request.GetProductUid())
	if err != nil {
		return &rpc.GetProductPriceResponse{}, err

	}

	return ToPBGetProductPriceResponse(dPP)

}

func (s catalogService) GetProductSupplier(
	ctx context.Context,
	request *rpc.GetProductSupplierRequest) (*rpc.ProductSupplier, error) {
	ps, err := s.repo.SelectProductSupplier(ctx, request.GetProductId(), request.GetSupplierId())
	if err != nil {
		return nil, err
	}

	return &rpc.ProductSupplier{
		ProductId:     ps.ProductId,
		SupplierId:    ps.SupplierId,
		SupplierSn:    ps.SupplierSn,
		ListPrice:     ps.ListPrice,
		PurchasePrice: ps.PurchasePrice,
		Delete:        false,
	}, nil
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

	var paths []string

	if !request.UpdateMask.IsValid(request.ProductSupplier) {
		return nil, errors.New("invalid field mask")
	}

	if fm := request.GetUpdateMask(); fm != nil && len(fm.Paths) > 0 {
		paths = fm.Paths
	}

	psp := domain.ProductSupplierPatch{}

	// TODO: Map specified paths to the domain.ProductSupplierPatch struct

	for _, p := range paths {
		switch p {
		case "list_price":
			psp.ListPrice = &request.GetProductSupplier().ListPrice
		case "purchase_price":
			psp.PurchasePrice = &request.GetProductSupplier().PurchasePrice
		case "supplier_sn":
			psp.SupplierProductSn = &request.GetProductSupplier().SupplierSn
		case "delete":
			psp.Delete = &request.GetProductSupplier().Delete
		}
	}

	if err := s.repo.UpdateProductSupplier(ctx, request.GetProductSupplier().ProductId,
		request.GetProductSupplier().SupplierId, &psp,
	); err != nil {
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
	res, nextCursor, err := s.repo.ListProducts(ctx, &domain.ProductFilter{Limit: 1000, Cursor: request.GetCursor()})
	if err != nil {
		return nil, err
	}
	return ToPBListProductResponse(res, nextCursor)
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
	_, span := s.tracer.Start(ctx, "getProduct")
	defer span.End()
	dProduct, err := s.repo.SelectProduct(ctx, request.GetId())
	if err != nil {
		return &rpc.GetProductResponse{}, err
	}

	pbProduct, err := ToPBProduct(dProduct)
	return &rpc.GetProductResponse{Product: pbProduct}, nil
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
	_, span := s.tracer.Start(ctx, "ListGroup")
	defer span.End()
	productGroups, err := s.repo.ListGroup()

	if err != nil {
		return &rpc.ListGroupResponse{}, err
	}
	return ToPBListProductGroup(productGroups)
}

func (s catalogService) GetGroup(
	ctx context.Context, request *rpc.GetGroupRequest,
) (*rpc.GetGroupResponse, error) {
	g, err := s.repo.FindGroupByID(ctx, request.GetId())
	if err != nil {
		fmt.Println(err)
	}
	p, _, err := s.repo.ListProducts(ctx, &domain.ProductFilter{GroupID: request.GetId()})
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
