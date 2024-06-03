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
	ShippingAddress      Address
	Contact              Contact
	Customer             Customer
	CustomerBranchId     float64 // This is ship_to_id in the database
	CustomerId           string
	CustomerName         string
	ContactId            string
	ContactName          string
	AddressId            float64
	PurchaseOrder        string
	DeliveryInstructions string
	Taker                string
	Status               OrderStatus
	OrderDate            time.Time
	RequestedDate        time.Time
	OrderType            OrderType
	Completed            bool
	Items                []*OrderItem
}

type OrderStatus struct {
	Approved  bool
	Cancelled bool
}

func ValidateOrder(v *validator.Validator, order *Order) {
	v.Check(order.Customer.Id != 0, "customer.id", "must be set")
	v.Check(order.AddressId != 0, "address_id", "must be set")
	v.Check(len(order.PurchaseOrder) <= 50, "purchase_order", "must be less than 50 characters")
}

type OrderItem struct {
	ProductUid          int32
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
