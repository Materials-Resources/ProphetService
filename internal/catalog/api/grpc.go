package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/catalog/service"
	"github.com/materials-resources/s_prophet/internal/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewCatalogApi(
	service service.CatalogService, tracer trace.Tracer, producer service.Producer,
) CatalogApi {
	return CatalogApi{service: service, tracer: tracer, producer: producer}
}

type CatalogApi struct {
	service  service.CatalogService
	tracer   trace.Tracer
	producer service.Producer
}

func (s CatalogApi) GetBasicProductDetails(
	ctx context.Context, request *rpc.GetBasicProductDetailsRequest) (*rpc.GetBasicProductDetailsResponse, error) {
	product, err := s.service.GetBasicProductDetails(ctx, request.GetProductUid())
	if err != nil {
		return nil, err
	}
	var res rpc.GetBasicProductDetailsResponse
	res.BasicProductDetails = make([]*rpc.GetBasicProductDetailsResponse_BasicProductDetail, len(product))
	for i, p := range product {
		res.BasicProductDetails[i] = &rpc.GetBasicProductDetailsResponse_BasicProductDetail{
			ProductUid: p.UID, Name: p.Name, Sn: p.SN,
		}

	}
	return &res, nil
}

func (s CatalogApi) ListSuppliers(ctx context.Context, request *rpc.ListSuppliersRequest) (
	*rpc.ListSuppliersResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s CatalogApi) ListProducts(ctx context.Context, request *rpc.ListProductsRequest) (
	*rpc.ListProductsResponse,
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
	return &rpc.ListProductsResponse{Products: rpcProducts, NextCursor: 0}, nil
}

func (s CatalogApi) GetProduct(ctx context.Context, request *rpc.GetProductRequest) (
	*rpc.GetProductResponse, error,
) {
	product, err := s.service.GetProduct(ctx, request.GetId())
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return nil, status.Error(codes.NotFound, "product not found")
		}

	}
	return &rpc.GetProductResponse{
		Product: &rpc.ProductDetail{
			Id:             product.UID,
			Sn:             product.SN,
			Name:           product.Name,
			Description:    product.Description,
			ProductGroupSn: product.ProductGroupSn,
		},
	}, nil
}

func (s CatalogApi) ListProductGroups(ctx context.Context, request *rpc.ListProductGroupsRequest) (
	*rpc.ListProductGroupsResponse, error,
) {
	_, span := s.tracer.Start(ctx, "ListGroup")
	defer span.End()
	productGroups, err := s.service.ListProductGroup(ctx)

	if err != nil {
		return &rpc.ListProductGroupsResponse{}, err
	}

	res := domainToProductGroups(productGroups)
	return &rpc.ListProductGroupsResponse{ProductGroups: res}, nil
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

func (s CatalogApi) GetProductGroup(
	ctx context.Context, request *rpc.GetProductGroupRequest,
) (*rpc.GetProductGroupResponse, error) {
	productGroup, productUids, err := s.service.GetProductGroup(ctx, request.GetSn())
	if err != nil {
		fmt.Println(err)
	}

	res := &rpc.GetProductGroupResponse{
		ProductGroup: &rpc.ProductGroup{
			Sn: productGroup.SN, Name: productGroup.Name, Uid: productGroup.UID,
		},
		ProductUids: productUids,
	}
	return res, nil
}

func (s CatalogApi) CreateProductGroup(
	ctx context.Context, request *rpc.CreateProductGroupRequest,
) (*rpc.CreateProductGroupResponse, error) {
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
		return &rpc.CreateProductGroupResponse{ValidationErrors: validationErrors}, status.Error(
			codes.InvalidArgument,
			"validation error")
	}
	err := s.service.CreateProductGroup(ctx, &productGroup)
	if err != nil {
		return &rpc.CreateProductGroupResponse{}, err
	}
	return &rpc.CreateProductGroupResponse{}, nil
}

func (s CatalogApi) UpdateProductGroup(
	ctx context.Context,
	request *rpc.UpdateProductGroupRequest) (*rpc.UpdateProductGroupResponse, error) {
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
	return &rpc.UpdateProductGroupResponse{}, nil
}

func (s CatalogApi) GetProductPrice(
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

func (s CatalogApi) GetSupplier(
	ctx context.Context,
	request *rpc.GetSupplierRequest,
) (*rpc.GetSupplierResponse, error) {
	// ps, err := s.repo.SelectProductSupplier(ctx, request.GetProductId(), request.GetSupplierId())
	// if err != nil {
	// 	return nil, err
	// }
	//
	// return &rpc.ProductSupplier{
	// 	ProductUid:     ps.ProductUid,
	// 	SupplierId:    ps.SupplierId,
	// 	SupplierSn:    ps.SupplierSn,
	// 	ListPrice:     ps.ListPrice,
	// 	PurchasePrice: ps.PurchasePrice,
	// 	Delete:        false,
	// }, nil
	return nil, nil
}

func (s CatalogApi) UpdateProduct(
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

func (s CatalogApi) CreateSupplier(
	ctx context.Context,
	request *rpc.CreateSupplierRequest,
) (*rpc.CreateSupplierResponse, error) {
	// d := &domain.ProductSupplier{
	// 	SupplierId: request.GetSupplierId(), SupplierSn: request.GetSupplierProductSn(),
	// 	ProductUid: request.GetProductId(), ListPrice: request.GetListPrice(), PurchasePrice: request.GetPurchasePrice(),
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

func (s CatalogApi) UpdateSupplier(
	ctx context.Context,
	request *rpc.UpdateSupplierRequest,
) (*rpc.UpdateSupplierResponse, error) {

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
	// 	ctx, request.GetProductSupplier().ProductUid,
	// 	request.GetProductSupplier().SupplierId, &psp,
	// ); err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (s CatalogApi) SetPrimarySupplier(
	ctx context.Context,
	request *rpc.SetPrimarySupplierRequest,
) (*rpc.SetPrimarySupplierResponse, error) {

	err := s.service.SetPrimaryProductSupplier(
		ctx, request.GetProductUid(), request.GetLocationId(), request.GetSupplierId(), request.GetDivisionId())
	if err != nil {
		return nil, err
	}
	return &rpc.SetPrimarySupplierResponse{}, nil
}

func (s CatalogApi) CreateProduct(
	ctx context.Context, request *rpc.CreateProductRequest,
) (
	*rpc.CreateProductResponse,
	error,
) {
	// TODO implement me
	panic("implement me")
}

func (s CatalogApi) DeleteProduct(
	ctx context.Context, request *rpc.DeleteProductRequest,
) (
	*rpc.DeleteProductResponse,
	error,
) {
	ctx, span := s.tracer.Start(ctx, "DeleteProduct", trace.WithSpanKind(trace.SpanKindServer))
	span.SetAttributes(attribute.String("request.id", request.String()))
	defer span.End()
	err := s.service.RequestDeleteProduct(ctx, request.GetUid())
	if err != nil {
		return &rpc.DeleteProductResponse{}, err
	}
	return &rpc.DeleteProductResponse{}, nil
}

func (s CatalogApi) GetProductBySupplier(
	ctx context.Context, request *rpc.GetProductBySupplierRequest,
) (*rpc.GetProductBySupplierResponse, error) {
	product, err := s.service.GetProductBySupplierPartNumber(ctx, request.GetSupplierPartNo(), request.GetSupplierId())
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			return nil, status.Error(codes.NotFound, "product not found")
		}

	}

	return &rpc.GetProductBySupplierResponse{
		Sn: product.SN,
	}, nil
}
