package models

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type ShipToSalesrep struct {
	bun.BaseModel        `bun:"table:ship_to_salesrep"`
	CompanyId            string         `bun:"company_id,type:varchar(8),pk"`
	ShipToId             float64        `bun:"ship_to_id,type:decimal(19,0),pk"`
	SalesrepId           string         `bun:"salesrep_id,type:varchar(16),pk"`
	DeleteFlag           string         `bun:"delete_flag,type:char"`
	DateCreated          time.Time      `bun:"date_created,type:datetime"`
	DateLastModified     time.Time      `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string         `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PrimarySalesrep      sql.NullString `bun:"primary_salesrep,type:char,nullzero"`
	CommissionPercentage float64        `bun:"commission_percentage,type:decimal(19,4),default:((100))"`
	PrimaryServiceRep    string         `bun:"primary_service_rep,type:char,default:('N')"`
}

type ShipToSalesrepModel struct {
	bun bun.IDB
}

func (m *ShipToSalesrepModel) GetByShipToId(
	ctx context.Context, companyId string, shipToId float64,
	primarySalesrep bool) (
	[]*ShipToSalesrep, error) {
	var shipToSalesrep []*ShipToSalesrep

	query := m.bun.NewSelect().Model(&shipToSalesrep).Where(
		"ship_to_id = ? AND company_id = ?", shipToId, companyId)

	if primarySalesrep {
		query = query.Where("primary_salesrep = 'Y'")
	}
	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return shipToSalesrep, nil
}
