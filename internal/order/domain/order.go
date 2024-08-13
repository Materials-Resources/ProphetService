package domain

import (
	"fmt"
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
	Id               string
	CustomerBranchId string
	Customer         Customer
	Contact          Contact
	ContactId        string
	ContactName      string
	PurchaseOrder    *string

	Taker         *string
	DateCreated   time.Time
	DateRequested *time.Time
	Items         []*OrderItem

	ShippingAddress      Address
	DeliveryInstructions *string
}

type OrderStatus struct {
	Approved  bool
	Cancelled bool
}

func ValidateOrder(v *validator.Validator, order *Order) {
	//v.Check(order.CustomerId != "", "customer.id", "must be set")
	v.Check(order.CustomerBranchId != "", "address_id", "must be set")
	if order.PurchaseOrder != nil {
		v.Check(len(*order.PurchaseOrder) <= 50, "purchase_order", "must be less than 50 characters")
	}
}

type OrderInput struct {
	PurchaseOrder *string
}

type OrderItemInput struct {
	ProductUid        string
	Quantity          float64
	UnitPrice         float64
	CustomerProductSn string
}

type OrderItem struct {
	Id                string
	ProductUid        string
	ProductSn         string
	ProductName       string
	CustomerProductSn string
	OrderQuantity     float64
	OrderQuantityUnit string
	UnitPrice         float64
	TotalPrice        float64
	ShippedQuantity   float64
}

func (oi *OrderItem) CalculateBackOrderedQuantity() float64 {
	return oi.OrderQuantity - oi.ShippedQuantity
}

type Contact struct {
	Id          string
	FirstName   string
	LastName    string
	PhoneNumber string
	Title       *string
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

type Customer struct {
	Id   string
	Name *string

	PhoneNumber string
}
