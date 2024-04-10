package catalog

import (
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	rpc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
)

// ToPBProduct Converts domain.Product to rpc.Product
func ToPBProduct(d *domain.Product) (*rpc.ProductDetail, error) {
	return &rpc.ProductDetail{Id: d.ID, Name: d.Name, Sn: d.SN, Description: d.Description,
		ProductGroupSn: d.ProductGroupId, StockQty: d.StockQuantity,
	}, nil
}

// ToPBListProductResponse Converts a slice of domain.Product to rpc.ListProductResponse
func ToPBListProductResponse(dProducts []*domain.Product, nc int32) (*rpc.ListProductResponse, error) {
	pb := &rpc.ListProductResponse{NextCursor: nc}
	for _, d := range dProducts {
		pbProduct, err := ToPBProduct(d)
		if err != nil {
			return nil, err
		}
		pb.Products = append(pb.Products, pbProduct)
	}
	return pb, nil
}

func ToPBGetProductPriceResponse(d []*domain.ProductPrice) (*rpc.GetProductPriceResponse, error) {

	pb := &rpc.GetProductPriceResponse{}

	for _, price := range d {
		pb.ProductPrices = append(pb.ProductPrices, &rpc.GetProductPriceResponse_ProductPrice{
			ProductUid: price.ProductUid,
			ListPrice:  price.ListPrice,
		},
		)
	}

	return pb, nil

}
func ToPBGetProductGroupResponse(
	productGroup *domain.ProductGroup, products []*domain.Product,
) (*rpc.GetGroupResponse, error) {
	var pb = &rpc.GetGroupResponse{
		ProductGroup: &rpc.ProductGroup{
			Sn:   productGroup.SN,
			Name: productGroup.Name,
		},
	}
	for _, product := range products {

		pbProduct, err := ToPBProduct(product)
		if err != nil {
			return nil, err
		}
		pb.Products = append(pb.Products, pbProduct)
	}
	return pb, nil
}

// ToPBProductGroup Converts domain.ProductGroup to rpc.ProductGroup
func ToPBProductGroup(productGroup *domain.ProductGroup) (*rpc.ProductGroup, error) {
	return &rpc.ProductGroup{
		Sn:   productGroup.SN,
		Name: productGroup.Name,
		Id:   productGroup.ID,
	}, nil
}

// ToPBListProductGroup Converts a slice of domain.ProductGroup to rpc.ListGroupResponse
func ToPBListProductGroup(dProductGroups []*domain.ProductGroup) (*rpc.ListGroupResponse, error) {
	var pb = &rpc.ListGroupResponse{}

	for _, productGroup := range dProductGroups {
		pbProductGroup, err := ToPBProductGroup(productGroup)
		if err != nil {
			return nil, err
		}
		pb.ProductGroups = append(pb.ProductGroups, pbProductGroup)

	}
	return pb, nil
}
