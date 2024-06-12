package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TpcxOutboundDocument struct {
	bun.BaseModel             `bun:"table:tpcx_outbound_document"`
	TpcxOutboundDocumentUid   int32     `bun:"tpcx_outbound_document_uid,type:int,pk,identity"`
	ApplicationServerLocation string    `bun:"application_server_location,type:varchar(255),nullzero"`
	MessageType               string    `bun:"message_type,type:varchar(255)"`
	QueueName                 string    `bun:"queue_name,type:varchar(255)"`
	SendmessageReturnCode     int16     `bun:"sendmessage_return_code,type:smallint,nullzero"`
	ErrorMessage              string    `bun:"error_message,type:varchar(255),nullzero"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DocumentPreview           string    `bun:"document_preview,type:varchar(1024),default:('')"`
	Document                  string    `bun:"document,type:text(2147483647),nullzero"`
}
