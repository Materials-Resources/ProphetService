package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type TpcxOutboundDocument struct {
	bun.BaseModel             `bun:"table:tpcx_outbound_document"`
	TpcxOutboundDocumentUid   int32     `bun:"tpcx_outbound_document_uid,type:int,autoincrement,scanonly,pk"` // UID
	ApplicationServerLocation string    `bun:"application_server_location,type:varchar(255),nullzero"`        // Jaguar server
	MessageType               string    `bun:"message_type,type:varchar(255)"`                                // Type of document, GETPA, EDI, etc
	QueueName                 string    `bun:"queue_name,type:varchar(255)"`                                  // Name of the message queue the document will be transitted on.
	SendmessageReturnCode     int16     `bun:"sendmessage_return_code,type:smallint,nullzero"`                // Return code from the QueueMessage component
	ErrorMessage              string    `bun:"error_message,type:varchar(255),nullzero"`                      // Any information about why the message could not be transmitted.
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`                                      // Will store whether the message was successfully transmitted or not.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`  // User who last changed the record
	DocumentPreview           string    `bun:"document_preview,type:varchar(1024),default:('')"`              // Displays partial document in the logging maint window in the application
	Document                  string    `bun:"document,type:text,nullzero"`                                   // XML document being sent (IMAGE NOT TEXT)
}
