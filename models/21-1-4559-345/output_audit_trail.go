package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OutputAuditTrail struct {
	bun.BaseModel       `bun:"table:output_audit_trail"`
	OutputAuditTrailUid int32     `bun:"output_audit_trail_uid,type:int,autoincrement,identity,pk"`    // Unique identifier of the table
	DocumentNumber      string    `bun:"document_number,type:varchar(255)"`                            // Document Number of the record been Printed, Emailed or Faxed
	DocumentType        string    `bun:"document_type,type:varchar(255)"`                              // Document Type of the record been Printed, Emailed or Faxed
	OutputType          string    `bun:"output_type,type:varchar(255)"`                                // Whether document it is been Printed, Emailed or Faxed
	FileName            string    `bun:"file_name,type:varchar(8000)"`                                 // File created that contains the transaction information
	FilePath            string    `bun:"file_path,type:varchar(8000)"`                                 // Location where the file was saved
	ClientName          *string   `bun:"client_name,type:varchar(255)"`                                // Computer name that generated the Printed, Emailed or Faxed request
	PrinterName         *string   `bun:"printer_name,type:varchar(255)"`                               // Printer name
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
