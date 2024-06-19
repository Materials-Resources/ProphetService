package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CcPaymentTypeXProcessor struct {
	bun.BaseModel          `bun:"table:cc_payment_type_x_processor"`
	CcPaymentXProcessorUid int32     `bun:"cc_payment_x_processor_uid,type:int,pk"`                       // Unique Identifier for the table.
	PaymentTypeId          float64   `bun:"payment_type_id,type:decimal(19,0),unique"`                    // Identifies associated payment type ID for the record.
	CreditcardProcessorUid int32     `bun:"creditcard_processor_uid,type:int,unique"`                     // Identifies associated credit card processor for the record.
	OrderSourceTypeCd      int32     `bun:"order_source_type_cd,type:int,unique"`                         // Identifies associated order source type [All orders, Web orders only & Non-web orders only]
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`                                     // Identifies the status of each row in the table.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	LocationId             float64   `bun:"location_id,type:decimal(19,0),unique"`                        // ID for location
}
