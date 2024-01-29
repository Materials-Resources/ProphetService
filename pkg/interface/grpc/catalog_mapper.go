package grpc

import (
	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
)

func ToPBGetProductGroupResponse(productGroup *entities.ValidatedProductGroup, products []*entities.ValidatedProduct,
) (*rpc.GetGroupResponse, error) {
	var pb = &rpc.GetGroupResponse{
		ProductGroup: &rpc.ProductGroup{
			Id:   productGroup.ID,
			Sn:   productGroup.SN,
			Name: productGroup.Name,
		},
	}
	for _, product := range products {
		pb.Products = append(pb.Products, &rpc.Product{
			Id:   product.ID,
			Sn:   product.SN,
			Name: product.Name,
		},
		)
	}
	return pb, nil
}

func ToPBListProductGroup(productGroups []*entities.ValidatedProductGroup) (*rpc.ListGroupResponse, error) {
	var pb = &rpc.ListGroupResponse{}

	for _, productGroup := range productGroups {
		pb.ProductGroups = append(pb.ProductGroups, &rpc.ProductGroup{
			Id:   productGroup.ID,
			Sn:   productGroup.SN,
			Name: productGroup.Name,
		},
		)
	}
	return pb, nil
}
