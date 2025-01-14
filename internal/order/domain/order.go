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

type OrderStatus int

const (
	OrderStatusUnknown OrderStatus = iota
	OrderStatusPendingApproval
	OrderStatusApproved
	OrderStatusCancelled
	OrderStatusCompleted
)

type Order struct {
	Id            string
	BranchId      string
	Customer      Customer
	Contact       Contact
	ContactId     string
	ContactName   string
	PurchaseOrder string
	Status        OrderStatus

	Taker         string
	DateCreated   time.Time
	DateRequested time.Time

	ShippingAddress      Address
	DeliveryInstructions string

	Items []*OrderItem
}

type OrderSummary struct {
	Id            string
	ContactId     string
	BranchId      string
	PurchaseOrder string
	Status        OrderStatus
	DateCreated   time.Time
	DateRequested time.Time
}

func ValidateOrder(v *validator.Validator, order *Order) {
	//v.Check(order.CustomerId != "", "customer.id", "must be set")
	v.Check(order.BranchId != "", "branch_id", "must be set")
	v.Check(len(order.PurchaseOrder) <= 50, "purchase_order", "must be less than 50 characters")
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
	ProductId         string
	ProductSn         string
	ProductName       string
	CustomerProductSn string
	OrderedQuantity   float64
	ShippedQuantity   float64
	UnitType          string
	UnitPrice         float64
	TotalPrice        float64
}

func (oi *OrderItem) CalculateBackOrderedQuantity() float64 {
	return oi.OrderedQuantity - oi.ShippedQuantity
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
