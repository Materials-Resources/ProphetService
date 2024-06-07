package api

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rpc "github.com/materials-resources/microservices-proto/golang/catalog"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
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

func (s CatalogApi) ClerkListGroups(ctx context.Context, request *rpc.ClerkListGroupRequest) (*rpc.ClerkListGroupResponse, error) {
	productGroup, err := s.service.ClerkListGroups(ctx)
	if err != nil {
		return nil, err

	}

	return CreateClerkListGroupsResponse(productGroup), nil

}

func (s CatalogApi) ClerkCreateGroup(ctx context.Context, request *rpc.ClerkCreateGroupRequest) (*rpc.ClerkCreateGroupResponse, error) {
	productGroup := ConvertClerkCreateGroupRequestToDomain(request)
	err := s.service.ClerkCreateGroup(ctx, productGroup)
	if err != nil {
		return nil, err

	}
	return &rpc.ClerkCreateGroupResponse{}, nil
}

func (s CatalogApi) ClerkReadGroup(ctx context.Context, request *rpc.ClerkReadGroupRequest) (*rpc.ClerkReadGroupResponse, error) {
	productGroup, err := s.service.ClerkReadGroup(ctx, request.GetSn())
	if err != nil {
		return nil, err

	}

	return CreateClerkReadGroupResponse(productGroup), nil
}

func (s CatalogApi) ClerkUpdateGroup(ctx context.Context, request *rpc.ClerkUpdateGroupRequest) (*rpc.ClerkUpdateGroupResponse, error) {
	productGroup := ConvertClerkUpdateGroupRequestToDomain(request)
	err := s.service.ClerkUpdateGroup(ctx, productGroup)
	if err != nil {
		return nil, err

	}
	return &rpc.ClerkUpdateGroupResponse{}, nil
}

func (s CatalogApi) ClerkDeleteGroup(ctx context.Context, request *rpc.ClerkDeleteGroupRequest) (*rpc.ClerkDeleteGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CatalogApi) CustomerGetProduct(ctx context.Context, request *rpc.CustomerGetProductRequest) (*rpc.CustomerGetProductResponse, error) {
	product, err := s.service.CustomerGetProduct(ctx, request.GetProductUid())
	if err != nil {
		return nil, err
	}
	return CreateCustomerGetProductResponse(product), nil
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
			Uid: p.Uid, Name: p.Name, Sn: p.Sn, Description: p.Description,
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
			Uid:            p.Uid,
			Sn:             p.Sn,
			Name:           p.Name,
			Description:    p.Description,
			StockQty:       p.StockQuantity,
			ProductGroupSn: p.ProductGroupId,
		}

	}
	return &rpc.ListProductsResponse{Products: rpcProducts}, nil
}

func (s CatalogApi) GetProduct(ctx context.Context, request *rpc.GetProductRequest) (
	*rpc.GetProductResponse, error,
) {
	product, err := s.service.GetProduct(ctx, request.GetUid())
	if err != nil {
		switch {
		case errors.Is(err, service.ErrorResourceNotFound):
			return nil, status.Error(codes.NotFound, "product not found")
		}

	}
	return &rpc.GetProductResponse{

		Uid:              product.Uid,
		Sn:               product.Sn,
		Name:             product.Name,
		Description:      product.Description,
		ProductGroupId:   product.ProductGroupId,
		ProductGroupName: product.ProductGroupName,
	}, nil
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
		Name:           request.GetProduct().GetName(),
		Description:    request.GetProduct().GetDescription(),
		ProductGroupId: request.GetProduct().GetProductGroupSn(),
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
	err := s.service.RequestDeleteProduct(ctx, string(request.GetUid()))
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
		Sn: product.Sn,
	}, nil
}
