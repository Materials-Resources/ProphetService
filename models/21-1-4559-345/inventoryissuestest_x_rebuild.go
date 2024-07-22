package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryissuestestXRebuild struct {
	bun.BaseModel    `bun:"table:inventoryissuestest_x_rebuild"`
	Testuid          int32     `bun:"testuid,type:int"`                                             // TestUID from InventoryIssuesTests.
	Rebuilduid       int32     `bun:"rebuilduid,type:int"`                                          // rebuilduid from InventoryIssuesRebuilds.
	SequenceNo       int32     `bun:"sequence_no,type:int,default:((1))"`                           // Sequence number that this rebuild should run in for this test (if more than one apply).
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	TotalTestUid     *int32    `bun:"total_test_uid,type:int"`
}
