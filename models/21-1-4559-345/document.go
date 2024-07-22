package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Document struct {
	bun.BaseModel     `bun:"table:document"`
	DocumentId        string    `bun:"document_id,type:varchar(40),pk"`                           // Unique Identifier of record
	DocumentDesc      string    `bun:"document_desc,type:varchar(255)"`                           // What is this document?
	DeleteFlag        string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CustomerSend      *int32    `bun:"customer_send,type:int"`                                    // Indicates if the document should be sent to ALL/SOME/NONE - code will be used
	DocumentDependent *int32    `bun:"document_dependent,type:int"`                               // Indicates if the document is dependent on the transaction, item or lot level
	PickTicketPrint   string    `bun:"pick_ticket_print,type:char(1),default:('N')"`              // Determines if the document should be printed with the Pick Ticket
	InvoicePrint      string    `bun:"invoice_print,type:char(1),default:('N')"`                  // Determines if the document should be printed with the Invoice
	PackingListPrint  string    `bun:"packing_list_print,type:char(1),default:('N')"`             // Determines if the document should be printed with the Packing List
	QuotePrint        string    `bun:"quote_print,type:char(1),default:('N')"`                    // Determines if the document should be printed with the Quote
	OrdAckPrint       string    `bun:"ord_ack_print,type:char(1),default:('N')"`                  // Determines if the document should be printed with the Order Acknowledgement
	PoPrint           string    `bun:"po_print,type:char(1),default:('N')"`                       // Determines if the document should be printed with the Printed PO and Printed PORG report
	DocumentType      *int32    `bun:"document_type,type:int"`                                    // Indicates the type of document (command line, Network, Word template) - selection affects the way the Path is defined
	Path              *string   `bun:"path,type:varchar(255)"`                                    // Path used to retrieve/create the document
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:((704))"`                  // Indicates the status of the record (active, inactive)
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // User who intially creates the record via the application
	DocumentUid       int32     `bun:"document_uid,type:int,autoincrement,unique,identity"`       // Unique Identifier for the record - will be unique and used for foreign key relationships
}
