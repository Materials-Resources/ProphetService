package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type LinkQuantity struct {
	bun.BaseModel    `bun:"table:link_quantity"`
	LinkQuantityUid  int32     `bun:"link_quantity_uid,type:int,autoincrement,scanonly,pk"`         // Unique Identifier for table
	FromUid          int32     `bun:"from_uid,type:int"`                                            // Transaction UID this is linked from
	FromTypeCd       int32     `bun:"from_type_cd,type:int"`                                        // Type of transaction this is linked from
	ToUid            int32     `bun:"to_uid,type:int"`                                              // Transaction UID this is linked to
	ToTypeCd         int32     `bun:"to_type_cd,type:int"`                                          // Type of transaction this is linked to
	Quantity         float64   `bun:"quantity,type:decimal(19,9)"`                                  // Amount of Quantity linked
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`                                     // Indicates the status of the record
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
