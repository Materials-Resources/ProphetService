package model

import (
	"github.com/uptrace/bun"
	"time"
)

type RmaReceiptHdr struct {
	bun.BaseModel     `bun:"table:rma_receipt_hdr"`
	RmaReceiptHdrUid  int32     `bun:"rma_receipt_hdr_uid,type:int,pk"`
	OeHdrUid          int32     `bun:"oe_hdr_uid,type:int"`
	ReceiptNo         float64   `bun:"receipt_no,type:decimal(19,0)"`
	ReceiptDate       time.Time `bun:"receipt_date,type:datetime,default:(getdate())"`
	FreightOut        float64   `bun:"freight_out,type:decimal(19,9),default:(0)"`
	Period            float64   `bun:"period,type:decimal(3,0)"`
	YearForPeriod     float64   `bun:"year_for_period,type:decimal(4,0)"`
	Confirmed         string    `bun:"confirmed,type:char,default:('N')"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	InvoiceNo         string    `bun:"invoice_no,type:varchar(10),default:(null)"`
	FrontCounterRma   string    `bun:"front_counter_rma,type:char,default:('N')"`
	InvoiceBatchUid   int32     `bun:"invoice_batch_uid,type:int,default:(1)"`
	CreditPercentage  float64   `bun:"credit_percentage,type:decimal(19,0),nullzero"`
	SignatureFilename string    `bun:"signature_filename,type:varchar(255),nullzero"`
	RfnavTransNo      int32     `bun:"rfnav_trans_no,type:int,nullzero"`
}
