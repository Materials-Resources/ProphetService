package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CellDefinition struct {
	bun.BaseModel    `bun:"table:cell_definition"`
	FinReportId      string    `bun:"fin_report_id,type:varchar(15),pk"`                         // System-generated ID that identifies a financial report.
	Cell             string    `bun:"cell,type:varchar(10),pk"`                                  // Specifies a particular cell in an Excel or Lotus spreadsheet.
	Field            string    `bun:"field,type:char(1)"`                                        // Describes which amount relating to the account(s) should be used.
	ColumnNo         float64   `bun:"column_no,type:decimal(19,0)"`                              // Specifies the column number.
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	AccountMask      string    `bun:"account_mask,type:varchar(32),nullzero"`                    // The selected account range.  This value can be strictly the account mask if all accounts of a particular mask should be selected.
}
