package prophet

import "github.com/uptrace/bun"

type Inventoryissuestestdesc struct {
	bun.BaseModel                     `bun:"table:InventoryIssuesTestDesc"`
	Testuid                           int32   `bun:"TestUID,type:int"`
	Testtype                          int32   `bun:"TestType,type:int"`
	Description                       string  `bun:"Description,type:varchar(255)"`
	Enabled                           bool    `bun:"Enabled,type:bit,default:((1))"`
	Config                            int32   `bun:"Config,type:int,default:((0))"`
	Bin                               *bool   `bun:"bin,type:bit"`
	Lot                               *bool   `bun:"Lot,type:bit"`
	LotBinXref                        *bool   `bun:"lot_bin_xref,type:bit"`
	Tag                               *bool   `bun:"tag,type:bit"`
	Serial                            *bool   `bun:"serial,type:bit"`
	Slab                              *bool   `bun:"slab,type:bit"`
	Rebuild                           *bool   `bun:"rebuild,type:bit"`
	Trans                             *bool   `bun:"trans,type:bit"`
	Allocation                        *bool   `bun:"allocation,type:bit"`
	OnHand                            *bool   `bun:"on_hand,type:bit"`
	Setup                             *bool   `bun:"setup,type:bit"`
	CustomTestNumber                  *int32  `bun:"custom_test_number,type:int"`
	TotalTestUid                      int32   `bun:"total_test_uid,type:int,autoincrement,identity"`
	InvLoc                            *bool   `bun:"inv_loc,type:bit,default:((0))"` // Track whether this test is part of the inv_loc test type.  (I.e. does the test include data from inv_loc.)
	DebugSql                          *string `bun:"debug_sql,type:varchar(max)"`
	DebugSqlDesc                      *string `bun:"debug_sql_desc,type:varchar(max)"`
	DebugSqlToken1                    *string `bun:"debug_sql_token1,type:varchar(max)"`
	DebugSqlToken2                    *string `bun:"debug_sql_token2,type:varchar(max)"`
	DebugSqlToken3                    *string `bun:"debug_sql_token3,type:varchar(max)"`
	DebugSqlToken4                    *string `bun:"debug_sql_token4,type:varchar(max)"`
	DebugSqlToken5                    *string `bun:"debug_sql_token5,type:varchar(max)"`
	DebugSqlToken6                    *string `bun:"debug_sql_token6,type:varchar(max)"`
	DebugSqlToken7                    *string `bun:"debug_sql_token7,type:varchar(max)"`
	DebugSqlToken8                    *string `bun:"debug_sql_token8,type:varchar(max)"`
	DebugSqlToken9                    *string `bun:"debug_sql_token9,type:varchar(max)"`
	DebugSqlToken10                   *string `bun:"debug_sql_token10,type:varchar(max)"`
	DebugSqlCustom                    *string `bun:"debug_sql_custom,type:varchar(max)"`
	DebugSqlDescCustom                *string `bun:"debug_sql_desc_custom,type:varchar(max)"`
	DebugSqlToken1Custom              *string `bun:"debug_sql_token1_custom,type:varchar(max)"`
	DebugSqlToken2Custom              *string `bun:"debug_sql_token2_custom,type:varchar(max)"`
	DebugSqlToken3Custom              *string `bun:"debug_sql_token3_custom,type:varchar(max)"`
	DebugSqlToken4Custom              *string `bun:"debug_sql_token4_custom,type:varchar(max)"`
	DebugSqlToken5Custom              *string `bun:"debug_sql_token5_custom,type:varchar(max)"`
	DebugSqlToken6Custom              *string `bun:"debug_sql_token6_custom,type:varchar(max)"`
	DebugSqlToken7Custom              *string `bun:"debug_sql_token7_custom,type:varchar(max)"`
	DebugSqlToken8Custom              *string `bun:"debug_sql_token8_custom,type:varchar(max)"`
	DebugSqlToken9Custom              *string `bun:"debug_sql_token9_custom,type:varchar(max)"`
	DebugSqlToken10Custom             *string `bun:"debug_sql_token10_custom,type:varchar(max)"`
	DebugSqlRebuild                   *string `bun:"debug_sql_rebuild,type:varchar(max)"`
	DebugSqlRebuildCustom             *string `bun:"debug_sql_rebuild_custom,type:varchar(max)"`
	DebugSqlReplaceSql                *string `bun:"debug_sql_replace_sql,type:varchar(max)"`
	DebugSqlCustomReplaceSql          *string `bun:"debug_sql_custom_replace_sql,type:varchar(max)"`
	DebugSqlRebuildReplaceSql         *string `bun:"debug_sql_rebuild_replace_sql,type:varchar(max)"`
	DebugSqlRebuildCustomReplaceSql   *string `bun:"debug_sql_rebuild_custom_replace_sql,type:varchar(max)"`
	DebugSqlP21ItemInfoReplaceSql     *string `bun:"debug_sql_p21_item_info_replace_sql,type:varchar(max)"`
	DebugSqlExtendedRebuildReplaceSql *string `bun:"debug_sql_extended_rebuild_replace_sql,type:varchar(max)"`
	AtomicLots                        *bool   `bun:"atomic_lots,type:bit"` // atomic_lots
}
