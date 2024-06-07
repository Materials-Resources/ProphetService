package api

import (
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)
import rpc "github.com/materials-resources/microservices-proto/golang/order"

func CreateClerkGetOrderResponse(o *domain.Order) *rpc.ClerkGetOrderResponse {
	res := &rpc.ClerkGetOrderResponse{
		Id:            o.Id,
		PurchaseOrder: o.PurchaseOrder,
		//ShippingAddressId:         o.ShippingAddressId,
		//ShippingAddressName:       o.ShippingAddressName,
		//ShippingAddressLineOne:    o.ShippingAddressLineOne,
		//ShippingAddressLineTwo:    o.ShippingAddressLineTwo,
		//ShippingAddressCity:       o.ShippingAddressCity,
		//ShippingAddressState:      o.ShippingAddressState,
		//ShippingAddressPostalCode: o.ShippingAddressPostalCode,
		//ShippingAddressCountry:    o.ShippingAddressCountry,
		DeliveryInstructions: o.DeliveryInstructions,
		DatePlaced:           timestamppb.New(o.OrderDate),
		Taker:                o.Taker,
		CustomerId:           o.CustomerId,
		CustomerName:         o.CustomerName,
		ContactId:            o.ContactId,
		ContactName:          o.ContactName,
	}

	if o.Items != nil {
		for _, item := range o.Items {
			res.Items = append(res.Items, &rpc.ClerkGetOrderResponse_Item{
				Base: CreateOrderItem(item),
			})
		}
	}
	return res
}
func CreateCustomerGetOrderResponse(o *domain.Order) *rpc.CustomerGetOrderResponse {
	res := &rpc.CustomerGetOrderResponse{
		Id:                        o.Id,
		PurchaseOrder:             o.PurchaseOrder,
		ShippingAddressId:         o.ShippingAddressId,
		ShippingAddressName:       o.ShippingAddressName,
		ShippingAddressLineOne:    o.ShippingAddressLineOne,
		ShippingAddressLineTwo:    o.ShippingAddressLineTwo,
		ShippingAddressCity:       o.ShippingAddressCity,
		ShippingAddressState:      o.ShippingAddressState,
		ShippingAddressPostalCode: o.ShippingAddressPostalCode,
		ShippingAddressCountry:    o.ShippingAddressCountry,
		DeliveryInstructions:      o.DeliveryInstructions,
		DatePlaced:                timestamppb.New(o.OrderDate),
		Taker:                     o.Taker,
		CustomerId:                o.CustomerId,
		CustomerName:              o.CustomerName,
		ContactId:                 o.ContactId,
		ContactName:               o.ContactName,
	}

	if o.Items != nil {
		for _, item := range o.Items {
			res.Items = append(res.Items, &rpc.CustomerGetOrderResponse_Item{
				Base: CreateOrderItem(item),
			})
		}
	}
	return res
}

func CreateCustomerGetShipmentsForOrderResponse(shipments []*domain.Shipment) *rpc.CustomerGetShipmentsForOrderResponse {
	res := &rpc.CustomerGetShipmentsForOrderResponse{}
	for _, s := range shipments {
		res.Shipments = append(res.Shipments, &rpc.Shipment{
			Id:              s.Id,
			CarrierName:     s.CarrierName,
			CarrierTracking: s.CarrierTracking,
		})
	}
	return res
}

func CreateOrderItem(item *domain.OrderItem) *rpc.OrderItem {
	return &rpc.OrderItem{
		Id:                  item.Id,
		ProductUid:          item.ProductUid,
		ProductSn:           item.ProductSn,
		ProductName:         item.ProductName,
		OrderedAsSn:         item.CustomerProductSn,
		QuantityOrdered:     item.OrderQuantity,
		QuantityUnit:        item.OrderQuantityUnit,
		TotalPrice:          item.TotalPrice,
		ShippedQuantity:     item.ShippedQuantity,
		BackOrderedQuantity: item.BackOrderedQuantity,
	}
}

func CreateOrder(o *domain.Order) *rpc.Order {
	res := &rpc.Order{
		Id:            o.Id,
		PurchaseOrder: o.PurchaseOrder,
		DatePlaced:    timestamppb.New(o.OrderDate),
	}
	return res
}

func CreateCustomerListOrdersResponse(orders []*domain.Order) *rpc.CustomerListOrdersResponse {
	res := &rpc.CustomerListOrdersResponse{}
	for _, o := range orders {
		res.Orders = append(res.Orders, CreateOrder(o))
	}
	return res
}

func CreateClerkGetShipmentResponse(shipment *domain.Shipment) *rpc.ClerkGetShipmentResponse {
	return &rpc.ClerkGetShipmentResponse{
		BaseShipment:              CreateShipment(shipment),
		ShippingAddressId:         shipment.ShippingAddressId,
		ShippingAddressName:       shipment.ShippingAddressName,
		ShippingAddressLineOne:    shipment.ShippingAddressLineOne,
		ShippingAddressLineTwo:    shipment.ShippingAddressLineTwo,
		ShippingAddressCity:       shipment.ShippingAddressCity,
		ShippingAddressState:      shipment.ShippingAddressState,
		ShippingAddressPostalCode: shipment.ShippingAddressPostalCode,
		ShippingAddressCountry:    shipment.ShippingAddressCountry,
		DeliveryInstructions:      shipment.DeliveryInstructions,
		OrderId:                   shipment.OrderId,
		OrderPurchaseOrder:        shipment.OrderPurchaseOrder,
		ContactId:                 shipment.ContactId,
		ContactName:               shipment.ContactName,
	}
}

func CreateShipment(s *domain.Shipment) *rpc.Shipment {
	return &rpc.Shipment{
		Id:              s.Id,
		CarrierName:     s.CarrierName,
		CarrierTracking: s.CarrierTracking,
	}
}
