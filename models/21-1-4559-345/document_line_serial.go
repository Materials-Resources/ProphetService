package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLineSerial struct {
	bun.BaseModel         `bun:"table:document_line_serial"`
	DocumentNo            float64   `bun:"document_no,type:decimal(19,0)"`
	LineNo                int32     `bun:"line_no,type:int"`
	DocumentType          string    `bun:"document_type,type:char(2)"`
	SerialNumber          string    `bun:"serial_number,type:varchar(40)"`
	DocumentCd            string    `bun:"document_cd,type:varchar(10),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`
	DimensionTrackingKey  int32     `bun:"dimension_tracking_key,type:int,nullzero"`
	DocumentLineSerialUid int32     `bun:"document_line_serial_uid,type:int,pk"`
	RowStatus             int16     `bun:"row_status,type:tinyint"`
	SubLineNo             int32     `bun:"sub_line_no,type:int"`
	SequenceNo            int32     `bun:"sequence_no,type:int,default:(0)"`
	SerialNumberUid       int32     `bun:"serial_number_uid,type:int,nullzero"`
	QtyFromTags           float64   `bun:"qty_from_tags,type:decimal(19,9),default:((-1))"`
	LotUid                int32     `bun:"lot_uid,type:int,nullzero"`
	RfBinUid              int32     `bun:"rf_bin_uid,type:int,nullzero"`
	DepositBinUid         int32     `bun:"deposit_bin_uid,type:int,nullzero"`
}
