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
