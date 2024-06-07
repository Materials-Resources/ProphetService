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
	return &rpc.Product{
		Uid:            p.Uid,
		Sn:             p.Sn,
		Name:           p.Name,
		Description:    p.Description,
		ProductGroupId: p.ProductGroupId,
	}
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

func CreateClerkReadGroupResponse(group *domain.ProductGroup) *rpc.ClerkReadGroupResponse {
	return &rpc.ClerkReadGroupResponse{
		ProductGroup: CreateProductGroup(group),
	}
}

func ConvertClerkCreateGroupRequestToDomain(r *rpc.ClerkCreateGroupRequest) *domain.ProductGroup {
	return &domain.ProductGroup{
		Sn:   r.Sn,
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
