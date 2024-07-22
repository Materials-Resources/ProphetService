package prophet

import "github.com/uptrace/bun"

type P21InventoryIssuesRow struct {
	bun.BaseModel `bun:"table:p21_inventory_issues_row"`
	IssueCount    int32  `bun:"issue_count,type:int"`
	RowId         string `bun:"row_id,type:varchar(1)"`
}
