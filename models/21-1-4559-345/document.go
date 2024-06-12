package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Document struct {
	bun.BaseModel     `bun:"table:document"`
	DocumentId        string    `bun:"document_id,type:varchar(40),pk"`
	DocumentDesc      string    `bun:"document_desc,type:varchar(255)"`
	DeleteFlag        string    `bun:"delete_flag,type:char"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CustomerSend      int32     `bun:"customer_send,type:int,nullzero"`
	DocumentDependent int32     `bun:"document_dependent,type:int,nullzero"`
	PickTicketPrint   string    `bun:"pick_ticket_print,type:char,default:('N')"`
	InvoicePrint      string    `bun:"invoice_print,type:char,default:('N')"`
	PackingListPrint  string    `bun:"packing_list_print,type:char,default:('N')"`
	QuotePrint        string    `bun:"quote_print,type:char,default:('N')"`
	OrdAckPrint       string    `bun:"ord_ack_print,type:char,default:('N')"`
	PoPrint           string    `bun:"po_print,type:char,default:('N')"`
	DocumentType      int32     `bun:"document_type,type:int,nullzero"`
	Path              string    `bun:"path,type:varchar(255),nullzero"`
	RowStatusFlag     int32     `bun:"row_status_flag,type:int,default:((704))"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DocumentUid       int32     `bun:"document_uid,type:int,identity"`
}
