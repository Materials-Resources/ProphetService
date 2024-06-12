package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdrNotepad struct {
	bun.BaseModel        `bun:"table:invoice_hdr_notepad"`
	InvoiceNo            string    `bun:"invoice_no,type:varchar(10),pk"`
	NoteId               float64   `bun:"note_id,type:decimal(19,0),pk"`
	Topic                string    `bun:"topic,type:varchar(30)"`
	Note                 string    `bun:"note,type:text(2147483647),nullzero"`
	ActivationDate       time.Time `bun:"activation_date,type:datetime"`
	ExpirationDate       time.Time `bun:"expiration_date,type:datetime"`
	EntryDate            time.Time `bun:"entry_date,type:datetime"`
	NotepadClassId       string    `bun:"notepad_class_id,type:varchar(8),nullzero"`
	Mandatory            string    `bun:"mandatory,type:char"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30)"`
	InvoiceHdrNotepadUid int32     `bun:"invoice_hdr_notepad_uid,type:int"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
