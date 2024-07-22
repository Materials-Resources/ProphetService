package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Gl struct {
	bun.BaseModel              `bun:"table:gl"`
	GlUid                      int32      `bun:"gl_uid,type:int,pk"`                                            // Unique identifier of a row.
	CompanyNo                  string     `bun:"company_no,type:varchar(8)"`                                    // Unique code that identifies a company.
	AccountNumber              string     `bun:"account_number,type:varchar(32)"`                               // The account which the transaction is posted to.
	Period                     float64    `bun:"period,type:decimal(3,0)"`                                      // The period in which the transaction is posted.
	YearForPeriod              float64    `bun:"year_for_period,type:decimal(4,0)"`                             // The year in which the transaction is posted.
	JournalId                  string     `bun:"journal_id,type:varchar(2)"`                                    // The journal which the transaction is associated with.
	Amount                     float64    `bun:"amount,type:decimal(19,4)"`                                     // The amount of the transaction.
	Source                     string     `bun:"source,type:varchar(255)"`                                      // A document number associated with a transaction.  Sometimes, this is user-defined.
	Description                string     `bun:"description,type:varchar(255)"`                                 // A description of the transaction.
	DateCreated                time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified           time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy           string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	CurrencyId                 float64    `bun:"currency_id,type:decimal(19,0)"`                                // Type of currency of the transaction.
	ForeignAmount              float64    `bun:"foreign_amount,type:decimal(19,4)"`                             // Foreign amount of the transaction.
	TransactionDate            time.Time  `bun:"transaction_date,type:datetime"`                                // Date of the transaction.
	TransactionNumber          *float64   `bun:"transaction_number,type:decimal(19,0),unique"`                  // Number assigned to each transaction.
	Approved                   *string    `bun:"approved,type:char(1)"`                                         // Indicates whether the transaction is approved.
	SequenceNumber             float64    `bun:"sequence_number,type:decimal(19,0),unique"`                     // Indicates the sequence in which to process the loc
	JobId                      *string    `bun:"job_id,type:varchar(32)"`                                       // Alphanumeric ID that is used for tracking and grouping of transactions.
	EncumberedAmount           *float64   `bun:"encumbered_amount,type:decimal(19,4)"`                          // Amount of transaction posted to the encumbered account.
	ForeignEncumberedAmount    *float64   `bun:"foreign_encumbered_amount,type:decimal(19,4)"`                  // Foreign amount posted to the encumbered account.
	DateApprovedFromUnapproved *time.Time `bun:"date_approved_from_unapproved,type:datetime"`                   // Indicates when the transaction became approved.
	SourceTypeCd               *int16     `bun:"source_type_cd,type:smallint"`                                  // Indicates the source of the gl record.
	RecordTypeCd               *int32     `bun:"record_type_cd,type:int"`                                       // code representing a certain type of GL record
	BankNo                     *float64   `bun:"bank_no,type:decimal(19,0)"`                                    // Stores bank number for a posting to a bank acct
	ClearedBank                *string    `bun:"cleared_bank,type:char(1)"`                                     // Has this payment been cleared by the bank?
	ClearedPeriod              *int16     `bun:"cleared_period,type:smallint"`                                  // The period in which the payment cleared the bank.
	ClearedYear                *int16     `bun:"cleared_year,type:smallint"`                                    // The year in which the payment cleared the bank.
	TransNoFromDeletion        *float64   `bun:"trans_no_from_deletion,type:decimal(19,0)"`                     // Transaction created from deleting an approved transaction
	GroupNumber                *int32     `bun:"group_number,type:int"`                                         // Group GL records within a Transaction Number
	CreatedBy                  *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`          // User who created the record
	LinkedTransactionTypeCd    *int32     `bun:"linked_transaction_type_cd,type:int"`                           // Code Number that identifies the Linked Transaction Type
	LinkedTransactionNumber    *int32     `bun:"linked_transaction_number,type:int"`                            // Transaction Number that is linked to the GL
	LinkedTransactionLineId    *int32     `bun:"linked_transaction_line_id,type:int"`                           // Transaction Line ID that is linked to the GL
	GlDimenTypeUid             *int32     `bun:"gl_dimen_type_uid,type:int"`                                    // UID for GL dimension type
	GlDimensionId              *string    `bun:"gl_dimension_id,type:varchar(255)"`                             // Value for dimension type
	GlDimensionDesc            *string    `bun:"gl_dimension_desc,type:varchar(255)"`                           // Desc of dimension type
	ExchangeRateManualEntry    *float64   `bun:"exchange_rate_manual_entry,type:decimal(19,6)"`                 // rate when creating manual journal entries
}
