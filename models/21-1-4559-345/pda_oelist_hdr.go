package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PdaOelistHdr struct {
	bun.BaseModel         `bun:"table:pda_oelist_hdr"`
	PdaOelistHdrUid       int32     `bun:"pda_oelist_hdr_uid,type:int,pk"`            // Unique ID of the record
	OelistListId          string    `bun:"oelist_list_id,type:varchar(255)"`          // User defined ID for the list, Ex Monday
	OelistListDescription *string   `bun:"oelist_list_description,type:varchar(255)"` // Extended description of the list. Ex All the customers I go to every Monday.
	ListDate              time.Time `bun:"list_date,type:datetime"`                   // Date the list was last generated.
	CompanyId             string    `bun:"company_id,type:varchar(8)"`                // Unique code that identifies a company.
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`            // Location whose items will be sold.
	SalesrepId            string    `bun:"salesrep_id,type:varchar(16)"`              // Salesrep whose is doing the selling.
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`                  // Indicates current record status.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`       // ID of the user who last maintained this record
}
