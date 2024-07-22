package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type RmaReceiptHdr struct {
	bun.BaseModel     `bun:"table:rma_receipt_hdr"`
	RmaReceiptHdrUid  int32     `bun:"rma_receipt_hdr_uid,type:int,pk"`                              // Receipt unique internal ID number
	OeHdrUid          int32     `bun:"oe_hdr_uid,type:int"`                                          // FK to oe_hdr.oe_hdr_uid - link to associated RMA header record
	ReceiptNo         float64   `bun:"receipt_no,type:decimal(19,0)"`                                // Receipt unique external ID number
	ReceiptDate       time.Time `bun:"receipt_date,type:datetime,default:(getdate())"`               // Receipt entry date
	FreightOut        float64   `bun:"freight_out,type:decimal(19,9),default:(0)"`                   // Out freight associated w/receipt
	Period            float64   `bun:"period,type:decimal(3,0)"`                                     // GL period associated w/receipt
	YearForPeriod     float64   `bun:"year_for_period,type:decimal(4,0)"`                            // GL year associated w/receipt
	Confirmed         string    `bun:"confirmed,type:char(1),default:('N')"`                         // Specifies receipt confirmation status
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:(704)"`                       // specifies record active status
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	InvoiceNo         *string   `bun:"invoice_no,type:varchar(10),default:(null)"`                   // FK to invoice_hdr.invoice_no column - link to associated invoice header record
	FrontCounterRma   string    `bun:"front_counter_rma,type:char(1),default:('N')"`                 // Indicates whether this receipt was received at the front counter.  Used for taxing purposes.
	InvoiceBatchUid   int32     `bun:"invoice_batch_uid,type:int,default:(1)"`                       // FK to invoice_batch.invoice_batch_uid - link to associated invoice_batch record
	CreditPercentage  *float64  `bun:"credit_percentage,type:decimal(19,0)"`                         // A custom column which is used to calculate the Refund Amount
	SignatureFilename *string   `bun:"signature_filename,type:varchar(255)"`                         // Stores the filename of the signature capture
	RfnavTransNo      *int32    `bun:"rfnav_trans_no,type:int"`                                      // F73204: RF Navigator transaction number
}
