package api

import (
	rpc "github.com/materials-resources/microservices-proto/golang/catalog"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
)

func CreateCustomerGetProductResponse(p *domain.Product) *rpc.CustomerGetProductResponse {
	return &rpc.CustomerGetProductResponse{
		Product: CreateProduct(p),
	}
}

func CreateClerkListGroupsResponse(groups []*domain.ProductGroup) *rpc.ClerkListGroupResponse {
	return &rpc.ClerkListGroupResponse{
		ProductGroups: CreateProductGroupSlice(groups),
	}
}

func CreateProductGroupSlice(groups []*domain.ProductGroup) []*rpc.ProductGroup {
	var result []*rpc.ProductGroup
	for _, g := range groups {
		result = append(result, CreateProductGroup(g))
	}
	return result
}

func CreateProduct(p *domain.Product) *rpc.Product {

	r := &rpc.Product{
		Uid:              p.Uid,
		ProductGroupName: p.ProductGroupName,
	}

	if p.Sn != nil {
		r.Sn = *p.Sn
	}

	if p.Name != nil {
		r.Name = *p.Name
	}

	if p.Description != nil {
		r.Description = *p.Description

	}
	if p.ProductGroupSn != nil {
		r.ProductGroupSn = *p.ProductGroupSn

	}
	return r
}

func CreateProductSlice(products []*domain.Product) []*rpc.Product {
	var result []*rpc.Product
	for _, p := range products {
		result = append(result, CreateProduct(p))
	}
	return result
}

func CreateProductGroup(p *domain.ProductGroup) *rpc.ProductGroup {
	return &rpc.ProductGroup{
		Sn:   p.Sn,
		Name: *p.Name,
	}
}

func CreateCursorMetadata(p *domain.PaginationMetadata) *rpc.CursorMetadata {
	return &rpc.CursorMetadata{
		PreviousCursor: int64(p.PreviousCursor),
		NextCursor:     int64(p.NextCursor),
		Total:          p.Total,
	}

}

func CreateClerkReadGroupResponse(group *domain.ProductGroup) *rpc.ClerkReadGroupResponse {
	return &rpc.ClerkReadGroupResponse{
		ProductGroup: CreateProductGroup(group),
	}
}

func CreateClerkReadProductResponse(p *domain.Product) *rpc.ClerkReadProductResponse {
	return &rpc.ClerkReadProductResponse{
		Product: CreateProduct(p),
	}
}

func CreateClerkListProductsResponse(products []*domain.Product, pagination *domain.PaginationMetadata) *rpc.ClerkListProductsResponse {
	res := &rpc.ClerkListProductsResponse{
		CursorMetadata: CreateCursorMetadata(pagination),
	}

	for _, p := range products {
		res.Products = append(res.Products, &rpc.ClerkListProductsResponse_ProductDetails{
			Product:       CreateProduct(p),
			StockQuantity: float32(p.StockQuantity),
		})

	}
	return res
}

func ConvertClerkCreateGroupRequestToDomain(r *rpc.ClerkCreateGroupRequest) *domain.ProductGroup {
	return &domain.ProductGroup{
		Sn:   r.GetSn(),
		Name: &r.Name,
	}
}

func ConvertClerkUpdateGroupRequestToDomain(r *rpc.ClerkUpdateGroupRequest) *domain.ProductGroup {
	productGroupDomain := &domain.ProductGroup{
		Sn: r.Sn,
	}
	paths := r.GetProductGroup().GetUpdateMask()

	for _, path := range paths.GetPaths() {
		if path == "name" {
			productGroupDomain.Name = &r.ProductGroup.Name
		}
	}
	return productGroupDomain
}

func ConvertClerkUpdateProductRequestToDomain(r *rpc.ClerkUpdateProductRequest) *domain.Product {
	productDomain := &domain.Product{
		Uid: r.GetProductUid(),
	}

	paths := r.GetProduct().GetUpdateMask()

	for _, path := range paths.GetPaths() {
		if path == "name" {
			productDomain.Name = &r.GetProduct().Name
		}
		if path == "description" {
			productDomain.Description = &r.GetProduct().Description
		}
		if path == "product_group_sn" {
			productDomain.ProductGroupSn = &r.GetProduct().ProductGroupSn
		}
	}
	return productDomain
}

func ConvertClerkListProductsRequestToDomain(r *rpc.ClerkListProductsRequest) *domain.Pagination {
	pagination := &domain.Pagination{
		Size: r.GetCursor().GetSize(),
	}

	switch r.GetCursor().GetDirection().(type) {
	case *rpc.Cursor_SelectPrevious:
		pagination.Direction = domain.CursorDirectionPrevious
		pagination.Cursor = int32(r.GetCursor().GetSelectPrevious())
	case *rpc.Cursor_SelectNext:
		pagination.Direction = domain.CursorDirectionNext
		pagination.Cursor = int32(r.GetCursor().GetSelectNext())
	}
	return pagination
}
