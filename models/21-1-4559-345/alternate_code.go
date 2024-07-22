package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type AlternateCode struct {
	bun.BaseModel      `bun:"table:alternate_code"`
	AlternateCode      string    `bun:"alternate_code,type:varchar(40),pk"`                        // What is the alternate code for this item?
	DeleteFlag         string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated        time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	InvMastUid         int32     `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	AlternateCodeDesc  *string   `bun:"alternate_code_desc,type:varchar(40)"`                      // Description of the alternate code item
	AlternateCodeUid   int32     `bun:"alternate_code_uid,type:int"`                               // Unique identifier for the alternate code record
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`      // User who created the record
	SourceTypeCd       int32     `bun:"source_type_cd,type:int,default:((2051))"`                  // Code that tells how the Alt Code was created (via Main or Wireless application)
	DefaultUom         *string   `bun:"default_uom,type:varchar(8)"`                               // The default UOM to use when an item is entered using this record's alternate code
	ExcludeFromB2bFlag *string   `bun:"exclude_from_b2b_flag,type:char(1)"`                        // Determines if we should exclude this alternate code from the B2B DTS.
}
