package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdrRma struct {
	bun.BaseModel          `bun:"table:oe_hdr_rma"`
	OeHdrRmaUid            int32     `bun:"oe_hdr_rma_uid,type:int,pk"`
	OeHdrUid               int32     `bun:"oe_hdr_uid,type:int"`
	ApplyCreditToInvoiceNo string    `bun:"apply_credit_to_invoice_no,type:char,default:('N')"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RetrievedByWms         string    `bun:"retrieved_by_wms,type:char,default:('N')"`
}
