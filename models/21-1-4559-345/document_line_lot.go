package model

type DocumentLineLot struct {
	bun.BaseModel      `bun:"table:document_line_lot"`
	DocumentNo         int32     `bun:"document_no,type:int"`
	LineNo             int32     `bun:"line_no,type:int"`
	DocumentType       string    `bun:"document_type,type:char(2)"`
	LotCd              string    `bun:"lot_cd,type:varchar(40)"`
	QtyToChange        float64   `bun:"qty_to_change,type:decimal(19,9)"`
	QtyApplied         float64   `bun:"qty_applied,type:decimal(19,9)"`
	UnitQuantity       float64   `bun:"unit_quantity,type:decimal(19,9)"`
	UnitOfMeasure      string    `bun:"unit_of_measure,type:varchar(8)"`
	DocumentCd         string    `bun:"document_cd,type:varchar(10),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`
	DocumentLineLotUid int32     `bun:"document_line_lot_uid,type:int,pk"`
	UnitSize           float64   `bun:"unit_size,type:decimal(19,4)"`
	SubLineNo          int32     `bun:"sub_line_no,type:int"`
	SequenceNo         int32     `bun:"sequence_no,type:int,default:(0)"`
	LotUid             int32     `bun:"lot_uid,type:int,nullzero"`
	QtyFromTags        float64   `bun:"qty_from_tags,type:decimal(19,9),default:((-1))"`
	ItemRevisionUid    int32     `bun:"item_revision_uid,type:int,nullzero"`
	CurrentbinUid      int32     `bun:"currentbin_uid,type:int,nullzero"`
	QaStatus           string    `bun:"qa_status,type:varchar(255),nullzero"`
	SubInvoiceNo       string    `bun:"sub_invoice_no,type:varchar(40),nullzero"`
}
