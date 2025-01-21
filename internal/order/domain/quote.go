package domain

import "time"

type QuoteStatus int

const (
	QuoteStatusUnknown QuoteStatus = iota
	QuoteStatusPendingApproval
	QuoteStatusApproved
	QuoteStatusCancelled
	QuoteStatusExpired
)

type QuoteSummary struct {
	Id            string
	ContactId     string
	BranchId      string
	PurchaseOrder string
	DateCreated   time.Time
	DateRequested time.Time
	DateExpires   time.Time

	Status QuoteStatus
}

type Quote struct {
	Id            string
	Branch        *Branch
	Customer      Customer
	Contact       Contact
	PurchaseOrder string
	DateCreated   time.Time
	DateRequested time.Time
	DateExpires   time.Time

	Status QuoteStatus
	Items  []*QuoteItem
}

type QuoteItem struct {
	Id                string
	ProductId         string
	ProductSn         string
	ProductName       string
	CustomerProductSn string
	OrderedQuantity   float64
	UnitType          string
	UnitPrice         float64
	TotalPrice        float64
}
