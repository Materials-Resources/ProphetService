package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvMastLinks struct {
	bun.BaseModel    `bun:"table:inv_mast_links"`
	InvMastLinksUid  int32     `bun:"inv_mast_links_uid,type:int,pk"`                          // Unique identifier for this inv_mast_links record
	InvMastUid       int32     `bun:"inv_mast_uid,type:int"`                                   // Unique identifier for the item id.
	LinkName         string    `bun:"link_name,type:varchar(30)"`                              // User defined name for inv_mast_link
	LinkPath         string    `bun:"link_path,type:varchar(255)"`                             // Path designating file to associate with this inv_mast_uid
	LinkArea         int32     `bun:"link_area,type:int"`                                      // Areas in the system where this link will be accessible
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(704)"`                  // Indicates current record status.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`          // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`    // Indicates the date/time this record was last modified.
	LastModifiedBy   string    `bun:"last_modified_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
