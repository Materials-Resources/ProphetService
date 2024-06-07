package domain

import (
	"github.com/materials-resources/s-prophet/internal/validator"
	"time"
)

type OrderType int

const (
	OrderTypeUnknown OrderType = iota
	OrderTypeQuote
	OrderTypeOrder
)

type Order struct {
	Id                   string
	CustomerBranchId     string
	CustomerId           string
	CustomerName         string
	ContactId            string
	ContactName          string
	PurchaseOrder        string
	DeliveryInstructions string
	Taker                string
	OrderDate            time.Time
	Items                []*OrderItem

	ShippingAddressId         string
	ShippingAddressName       string
	ShippingAddressLineOne    string
	ShippingAddressLineTwo    string
	ShippingAddressCity       string
	ShippingAddressState      string
	ShippingAddressPostalCode string
	ShippingAddressCountry    string
}

type OrderStatus struct {
	Approved  bool
	Cancelled bool
}

func ValidateOrder(v *validator.Validator, order *Order) {
	v.Check(order.CustomerId != "", "customer.id", "must be set")
	v.Check(order.CustomerBranchId != "", "address_id", "must be set")
	v.Check(len(order.PurchaseOrder) <= 50, "purchase_order", "must be less than 50 characters")
}

type OrderItem struct {
	Id                  string
	ProductUid          string
	ProductSn           string
	ProductName         string
	CustomerProductSn   string
	OrderQuantity       float64
	OrderQuantityUnit   string
	PriceUnit           string
	Price               float64
	TotalPrice          float64
	ShippedQuantity     float64
	BackOrderedQuantity float64
}

type Contact struct {
	Id          string
	Name        string
	PhoneNumber string
	Title       string
}

type Customer struct {
	Id          float64
	Name        string
	PhoneNumber string
	Title       string
}
