package catalog

import (
	"context"
	"errors"
	"fmt"
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/catalog/service"
	"github.com/materials-resources/s_prophet/internal/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type catalogService struct {
	service  service.CatalogService
	tracer   trace.Tracer
	producer service.Producer
}

func (s catalogService) ListProductGroup(ctx context.Context, request *rpc.ListGroupRequest) (
	*rpc.ListGroupResponse, error,
) {
	_, span := s.tracer.Start(ctx, "ListGroup")
	defer span.End()
	productGroups, err := s.service.ListProductGroup(ctx)

	if err != nil {
		return &rpc.ListGroupResponse{}, err
	}

	res := domainToProductGroups(productGroups)
	return &rpc.ListGroupResponse{ProductGroups: res}, nil
}

func domainToProductGroups(productGroups []*domain.ProductGroup) []*rpc.ProductGroup {
	res := make([]*rpc.ProductGroup, len(productGroups))

	for i, pg := range productGroups {
		res[i] = &rpc.ProductGroup{
			Sn:   pg.SN,
			Name: pg.Name,
		}

	}
	return res
}

func (s catalogService) GetProductGroup(
	ctx context.Context, request *rpc.GetProductGroupRequest,
) (*rpc.GetProductGroupResponse, error) {
	productGroup, products, err := s.service.GetProductGroup(ctx, request.GetSn())
	if err != nil {
		fmt.Println(err)
	}

	res := &rpc.GetProductGroupResponse{
		ProductGroup: &rpc.ProductGroup{
			Sn: productGroup.SN, Name: productGroup.Name, Uid: productGroup.UID,
		},
	}

	res.Products = make([]*rpc.ProductDetail, len(products))
	for i, product := range products {
		res.Products[i] = &rpc.ProductDetail{
			Id:          product.UID,
			Name:        product.Name,
			Sn:          product.SN,
			Description: product.Description,
			StockQty:    product.StockQuantity,
		}

	}
	return res, nil
}

func (s catalogService) CreateProductGroup(
	ctx context.Context, request *rpc.CreateGroupRequest,
) (*rpc.CreateGroupResponse, error) {
	_, span := s.tracer.Start(ctx, "CreateGroup")
	defer span.End()

	productGroup := domain.ProductGroup{SN: request.GetProductGroup().GetSn(), Name: request.ProductGroup.GetName()}

	v := validator.New()

	domain.ValidateProductGroup(v, productGroup)

	if !v.Valid() {

		var validationErrors []*rpc.ValidationError
		for key, msg := range v.Errors {
			validationErrors = append(validationErrors, &rpc.ValidationError{Field: key, Message: msg})
		}
		return &rpc.CreateGroupResponse{ValidationErrors: validationErrors}, status.Error(
			codes.InvalidArgument,
			"validation error")
	}
	err := s.service.CreateProductGroup(ctx, &productGroup)
	if err != nil {
		return &rpc.CreateGroupResponse{}, err
	}
	return &rpc.CreateGroupResponse{}, nil
}

func (s catalogService) UpdateProductGroup(
	ctx context.Context,
	request *rpc.UpdateGroupRequest) (*rpc.UpdateGroupResponse, error) {
	productGroup := &domain.ProductGroup{
		SN: request.GetProductGroup().GetSn(), Name: request.GetProductGroup().GetName(),
	}
	v := validator.New()
	domain.ValidateProductGroup(v, *productGroup)

	if !v.Valid() {
		return nil, status.Error(codes.InvalidArgument, "validation error")

	}
	err := s.service.UpdateProductGroup(ctx, productGroup)
	if err != nil {
		return nil, err
	}
	return &rpc.UpdateGroupResponse{}, nil
}

func (s catalogService) GetProductPrice(
	ctx context.Context,
	request *rpc.GetProductPriceRequest,
) (*rpc.GetProductPriceResponse, error) {
	// dPP, err := s.repo.SelectProductPrice(ctx, request.GetProductUid())
	// if err != nil {
	// 	return &rpc.GetProductPriceResponse{}, err
	//
	// }
	//
	// productPrices := make([]*rpc.GetProductPriceResponse_ProductPrice, 0, len(dPP))
	// for _, pp := range dPP {
	// 	productPrices = append(
	// 		productPrices, &rpc.GetProductPriceResponse_ProductPrice{
	// 			ProductUid: pp.ProductUid, CustomerPrice: pp.CustomerPrice, ListPrice: pp.ListPrice,
	// 		})
	//
	// }

	// return &rpc.GetProductPriceResponse{ProductPrices: productPrices}, nil
	return nil, nil

}

func (s catalogService) GetProductSupplier(
	ctx context.Context,
	request *rpc.GetProductSupplierRequest,
) (*rpc.ProductSupplier, error) {
	// ps, err := s.repo.SelectProductSupplier(ctx, request.GetProductId(), request.GetSupplierId())
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return &rpc.ProductSupplier{
	// 	ProductId:     ps.ProductId,
	// 	SupplierId:    ps.SupplierId,
	// 	SupplierSn:    ps.SupplierSn,
	// 	ListPrice:     ps.ListPrice,
	// 	PurchasePrice: ps.PurchasePrice,
	// 	Delete:        false,
	// }, nil
	return nil, nil
}

func (s catalogService) UpdateProduct(
	ctx context.Context,
	request *rpc.UpdateProductRequest,
) (*rpc.UpdateProductResponse, error) {

	product := &domain.Product{
		UID:            request.GetProduct().GetId(),
		Name:           request.GetProduct().GetName(),
		Description:    request.GetProduct().GetDescription(),
		ProductGroupSn: request.GetProduct().GetProductGroupSn(),
	}

	s.producer.PublishUpdateProduct(ctx, product)

	// err := s.service.UpdateProduct(ctx, product, []float64{1001})
	// if err != nil {
	// 	return nil, err
	// }

	return &rpc.UpdateProductResponse{}, nil
}

func (s catalogService) CreateProductSupplier(
	ctx context.Context,
	request *rpc.CreateProductSupplierRequest,
) (*rpc.CreateProductSupplierResponse, error) {
	// d := &domain.ProductSupplier{
	// 	SupplierId: request.GetSupplierId(), SupplierSn: request.GetSupplierProductSn(),
	// 	ProductId: request.GetProductId(), ListPrice: request.GetListPrice(), PurchasePrice: request.GetPurchasePrice(),
	// }
	// vd, err := domain.NewValidatedProductSupplier(d)
	//
	// if err != nil {
	// 	return nil, err
	// }
	// s.repo.AddProductSupplier(ctx, vd)
	// return nil, nil
	return nil, nil
}

func (s catalogService) UpdateProductSupplier(
	ctx context.Context,
	request *rpc.UpdateProductSupplierRequest,
) (*rpc.UpdateProductSupplierResponse, error) {

	// var paths []string
	//
	// if !request.UpdateMask.IsValid(request.ProductSupplier) {
	// 	return nil, errors.New("invalid field mask")
	// }
	//
	// if fm := request.GetUpdateMask(); fm != nil && len(fm.Paths) > 0 {
	// 	paths = fm.Paths
	// }
	//
	// psp := domain.ProductSupplierPatch{}
	//
	// // TODO: Map specified paths to the domain.ProductSupplierPatch struct
	//
	// for _, p := range paths {
	// 	switch p {
	// 	case "list_price":
	// 		psp.ListPrice = &request.GetProductSupplier().ListPrice
	// 	case "purchase_price":
	// 		psp.PurchasePrice = &request.GetProductSupplier().PurchasePrice
	// 	case "supplier_sn":
	// 		psp.SupplierProductSn = &request.GetProductSupplier().SupplierSn
	// 	case "delete":
	// 		psp.Delete = &request.GetProductSupplier().Delete
	// 	}
	// }
	//
	// if err := s.repo.UpdateProductSupplier(
	// 	ctx, request.GetProductSupplier().ProductId,
	// 	request.GetProductSupplier().SupplierId, &psp,
	// ); err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (s catalogService) SetPrimaryProductSupplier(
	ctx context.Context,
	request *rpc.SetPrimaryProductSupplierRequest,
) (*rpc.SetPrimaryProductSupplierResponse, error) {
	// err := s.repo.SetPrimaryProductSupplier(ctx, request.GetProductId(), request.GetSupplierId())
	// if err != nil {
	// 	return nil, err
	// }
	// return &rpc.SetPrimaryProductSupplierResponse{}, nil
	return nil, nil
}

func (s catalogService) ListProduct(ctx context.Context, request *rpc.ListProductRequest) (
	*rpc.ListProductResponse,
	error,
) {
	filter := &domain.Filter{Cursor: int(request.GetCursor()), Limit: 200}
	products, err := s.service.ListProduct(ctx, *filter)

	if err != nil {
		return nil, err
	}

	rpcProducts := make([]*rpc.ProductDetail, len(products))

	for i, p := range products {
		rpcProducts[i] = &rpc.ProductDetail{
			Id:             p.UID,
			Sn:             p.SN,
			Name:           p.Name,
			Description:    p.Description,
			StockQty:       p.StockQuantity,
			ProductGroupSn: p.ProductGroupSn,
		}

	}
	return &rpc.ListProductResponse{Products: rpcProducts, NextCursor: 0}, nil
}

func (s catalogService) GetProduct(ctx context.Context, request *rpc.GetProductRequest) (
	*rpc.GetProductResponse, error,
) {
	_, span := s.tracer.Start(ctx, "getProduct")
	defer span.End()
	product, err := s.service.GetProductBySupplierPartNumber(ctx, "OBS50358", 103687)
	if err != nil {
		switch {
		case errors.Is(err, prophet_21_1_4559.ErrNotFound):
			return nil, status.Error(codes.NotFound, "product not found")
		}

	}
	return &rpc.GetProductResponse{
		Product: &rpc.ProductDetail{
			Id:          product.UID,
			Sn:          product.SN,
			Name:        product.Name,
			Description: product.Description,
		},
	}, nil
}

func (s catalogService) CreateProduct(
	ctx context.Context, request *rpc.CreateProductRequest,
) (
	*rpc.CreateProductResponse,
	error,
) {
	// TODO implement me
	panic("implement me")
}

func (s catalogService) DeleteProduct(
	ctx context.Context, request *rpc.DeleteProductRequest,
) (
	*rpc.DeleteProductResponse,
	error,
) {
	ctx, span := s.tracer.Start(ctx, "DeleteProduct", trace.WithSpanKind(trace.SpanKindServer))
	span.SetAttributes(attribute.String("request.id", request.String()))
	defer span.End()
	err := s.service.DeleteProduct(ctx, request.GetUid())
	if err != nil {
		return &rpc.DeleteProductResponse{}, err
	}

	// s.producer.PublishDelete(ctx, request.GetId())

	// err := s.repo.DeleteProduct(ctx, request.GetId())
	// if err != nil {
	//	span.RecordError(err)
	//	span.SetStatus(codes.Error, err.Error())
	//	return nil, err
	// }
	return &rpc.DeleteProductResponse{}, nil
}

func (s catalogService) GetProductBySupplier(
	ctx context.Context, request *rpc.GetProductBySupplierRequest,
) (*rpc.GetProductBySupplierResponse, error) {
	// TODO implement me
	panic("implement me")
}
