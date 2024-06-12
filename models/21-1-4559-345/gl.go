package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Gl struct {
	bun.BaseModel              `bun:"table:gl"`
	GlUid                      int32     `bun:"gl_uid,type:int,pk"`
	CompanyNo                  string    `bun:"company_no,type:varchar(8)"`
	AccountNumber              string    `bun:"account_number,type:varchar(32)"`
	Period                     float64   `bun:"period,type:decimal(3,0)"`
	YearForPeriod              float64   `bun:"year_for_period,type:decimal(4,0)"`
	JournalId                  string    `bun:"journal_id,type:varchar(2)"`
	Amount                     float64   `bun:"amount,type:decimal(19,4)"`
	Source                     string    `bun:"source,type:varchar(255)"`
	Description                string    `bun:"description,type:varchar(255)"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0)"`
	ForeignAmount              float64   `bun:"foreign_amount,type:decimal(19,4)"`
	TransactionDate            time.Time `bun:"transaction_date,type:datetime"`
	TransactionNumber          float64   `bun:"transaction_number,type:decimal(19,0),nullzero"`
	Approved                   string    `bun:"approved,type:char,nullzero"`
	SequenceNumber             float64   `bun:"sequence_number,type:decimal(19,0)"`
	JobId                      string    `bun:"job_id,type:varchar(32),nullzero"`
	EncumberedAmount           float64   `bun:"encumbered_amount,type:decimal(19,4),nullzero"`
	ForeignEncumberedAmount    float64   `bun:"foreign_encumbered_amount,type:decimal(19,4),nullzero"`
	DateApprovedFromUnapproved time.Time `bun:"date_approved_from_unapproved,type:datetime,nullzero"`
	SourceTypeCd               int16     `bun:"source_type_cd,type:smallint,nullzero"`
	RecordTypeCd               int32     `bun:"record_type_cd,type:int,nullzero"`
	BankNo                     float64   `bun:"bank_no,type:decimal(19,0),nullzero"`
	ClearedBank                string    `bun:"cleared_bank,type:char,nullzero"`
	ClearedPeriod              int16     `bun:"cleared_period,type:smallint,nullzero"`
	ClearedYear                int16     `bun:"cleared_year,type:smallint,nullzero"`
	TransNoFromDeletion        float64   `bun:"trans_no_from_deletion,type:decimal(19,0),nullzero"`
	GroupNumber                int32     `bun:"group_number,type:int,nullzero"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	LinkedTransactionTypeCd    int32     `bun:"linked_transaction_type_cd,type:int,nullzero"`
	LinkedTransactionNumber    int32     `bun:"linked_transaction_number,type:int,nullzero"`
	LinkedTransactionLineId    int32     `bun:"linked_transaction_line_id,type:int,nullzero"`
	GlDimenTypeUid             int32     `bun:"gl_dimen_type_uid,type:int,nullzero"`
	GlDimensionId              string    `bun:"gl_dimension_id,type:varchar(255),nullzero"`
	GlDimensionDesc            string    `bun:"gl_dimension_desc,type:varchar(255),nullzero"`
	ExchangeRateManualEntry    float64   `bun:"exchange_rate_manual_entry,type:decimal(19,6),nullzero"`
}
