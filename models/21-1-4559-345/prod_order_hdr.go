package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderHdr struct {
	bun.BaseModel            `bun:"table:prod_order_hdr"`
	ProdOrderNumber          float64   `bun:"prod_order_number,type:decimal(19,0),pk"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	SourceLocationId         float64   `bun:"source_location_id,type:decimal(19,0)"`
	OrderDate                time.Time `bun:"order_date,type:datetime"`
	RequiredDate             time.Time `bun:"required_date,type:datetime"`
	ExpectedCompletionDate   time.Time `bun:"expected_completion_date,type:datetime"`
	Approved                 string    `bun:"approved,type:char"`
	Complete                 string    `bun:"complete,type:char"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	Cancel                   string    `bun:"cancel,type:char"`
	Comment                  string    `bun:"comment,type:varchar(255),nullzero"`
	EnteredBy                string    `bun:"entered_by,type:varchar(30)"`
	Period                   float64   `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod            float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`
	AssembleDisassemble      string    `bun:"assemble_disassemble,type:char"`
	Printed                  string    `bun:"printed,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	CreateAutoTransfer       string    `bun:"create_auto_transfer,type:char,default:('N')"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	RetrievedByWms           string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	ExpediteDate             time.Time `bun:"expedite_date,type:datetime,nullzero"`
	EstimatedFreightCost     float64   `bun:"estimated_freight_cost,type:decimal(19,9),nullzero"`
	ActualFreightCost        float64   `bun:"actual_freight_cost,type:decimal(19,9),nullzero"`
	ManualFreightOverideFlag string    `bun:"manual_freight_overide_flag,type:char,default:('N')"`
	ReleaseDate              time.Time `bun:"release_date,type:datetime,nullzero"`
	DrawingsComplete         string    `bun:"drawings_complete,type:char,default:('N')"`
	InventoryLocationId      float64   `bun:"inventory_location_id,type:decimal(19,0),nullzero"`
}
