package models

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type OeHdrSalesrep struct {
	bun.BaseModel              `bun:"table:oe_hdr_salesrep"`
	OrderNumber                string          `bun:"order_number,type:varchar(8),pk"`
	SalesrepId                 string          `bun:"salesrep_id,type:varchar(16),pk"`
	CommissionSplit            float64         `bun:"commission_split,type:decimal(5,2)"`
	DeleteFlag                 string          `bun:"delete_flag,type:char"`
	DateCreated                time.Time       `bun:"date_created,type:datetime"`
	DateLastModified           time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DatePaid                   sql.NullTime    `bun:"date_paid,type:datetime,nullzero"`
	PrimarySalesrep            sql.NullString  `bun:"primary_salesrep,type:char,nullzero"`
	ExcludeSplitValidationFlag string          `bun:"exclude_split_validation_flag,type:char,default:('Y')"`
	CommissionOverridePercent  sql.NullFloat64 `bun:"commission_override_percent,type:decimal(19,9),nullzero"`
}

var _ bun.BeforeAppendModelHook = (*OeHdrSalesrep)(nil)

func (m *OeHdrSalesrep) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type OeHdrSalesrepModel struct {
	bun bun.IDB
}

type OeHdrSalesrepParams struct {
	OrderNumber     string
	SalesrepId      string
	CommissionSplit float64
	PrimarySalesrep string
}

func (m *OeHdrSalesrepModel) Create(ctx context.Context, params OeHdrSalesrepParams) (*OeHdrSalesrep, error) {
	oeHdrSalesrep := &OeHdrSalesrep{
		OrderNumber:      params.OrderNumber,
		SalesrepId:       params.SalesrepId,
		CommissionSplit:  params.CommissionSplit,
		PrimarySalesrep:  sql.NullString{String: params.PrimarySalesrep, Valid: params.PrimarySalesrep != ""},
		LastMaintainedBy: "admin",
	}
	err := m.Insert(ctx, oeHdrSalesrep)
	return oeHdrSalesrep, err
}

func (m *OeHdrSalesrepModel) Insert(ctx context.Context, oeHdrSalesrep *OeHdrSalesrep) error {
	_, err := m.bun.NewInsert().Model(oeHdrSalesrep).Exec(ctx)
	return err
}
