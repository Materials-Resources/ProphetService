package model

type Inventoryissuestestdesc struct {
	bun.BaseModel                     `bun:"table:InventoryIssuesTestDesc"`
	Testuid                           int32  `bun:"TestUID,type:int"`
	Testtype                          int32  `bun:"TestType,type:int"`
	Description                       string `bun:"Description,type:varchar(255)"`
	Enabled                           `bun:"Enabled,type:bit,default:((1))"`
	Config                            int32 `bun:"Config,type:int,default:((0))"`
	Bin                               `bun:"bin,type:bit,nullzero"`
	Lot                               `bun:"Lot,type:bit,nullzero"`
	LotBinXref                        `bun:"lot_bin_xref,type:bit,nullzero"`
	Tag                               `bun:"tag,type:bit,nullzero"`
	Serial                            `bun:"serial,type:bit,nullzero"`
	Slab                              `bun:"slab,type:bit,nullzero"`
	Rebuild                           `bun:"rebuild,type:bit,nullzero"`
	Trans                             `bun:"trans,type:bit,nullzero"`
	Allocation                        `bun:"allocation,type:bit,nullzero"`
	OnHand                            `bun:"on_hand,type:bit,nullzero"`
	Setup                             `bun:"setup,type:bit,nullzero"`
	CustomTestNumber                  int32 `bun:"custom_test_number,type:int,nullzero"`
	TotalTestUid                      int32 `bun:"total_test_uid,type:int,identity"`
	InvLoc                            `bun:"inv_loc,type:bit,default:((0))"`
	DebugSql                          string `bun:"debug_sql,type:varchar,nullzero"`
	DebugSqlDesc                      string `bun:"debug_sql_desc,type:varchar,nullzero"`
	DebugSqlToken1                    string `bun:"debug_sql_token1,type:varchar,nullzero"`
	DebugSqlToken2                    string `bun:"debug_sql_token2,type:varchar,nullzero"`
	DebugSqlToken3                    string `bun:"debug_sql_token3,type:varchar,nullzero"`
	DebugSqlToken4                    string `bun:"debug_sql_token4,type:varchar,nullzero"`
	DebugSqlToken5                    string `bun:"debug_sql_token5,type:varchar,nullzero"`
	DebugSqlToken6                    string `bun:"debug_sql_token6,type:varchar,nullzero"`
	DebugSqlToken7                    string `bun:"debug_sql_token7,type:varchar,nullzero"`
	DebugSqlToken8                    string `bun:"debug_sql_token8,type:varchar,nullzero"`
	DebugSqlToken9                    string `bun:"debug_sql_token9,type:varchar,nullzero"`
	DebugSqlToken10                   string `bun:"debug_sql_token10,type:varchar,nullzero"`
	DebugSqlCustom                    string `bun:"debug_sql_custom,type:varchar,nullzero"`
	DebugSqlDescCustom                string `bun:"debug_sql_desc_custom,type:varchar,nullzero"`
	DebugSqlToken1Custom              string `bun:"debug_sql_token1_custom,type:varchar,nullzero"`
	DebugSqlToken2Custom              string `bun:"debug_sql_token2_custom,type:varchar,nullzero"`
	DebugSqlToken3Custom              string `bun:"debug_sql_token3_custom,type:varchar,nullzero"`
	DebugSqlToken4Custom              string `bun:"debug_sql_token4_custom,type:varchar,nullzero"`
	DebugSqlToken5Custom              string `bun:"debug_sql_token5_custom,type:varchar,nullzero"`
	DebugSqlToken6Custom              string `bun:"debug_sql_token6_custom,type:varchar,nullzero"`
	DebugSqlToken7Custom              string `bun:"debug_sql_token7_custom,type:varchar,nullzero"`
	DebugSqlToken8Custom              string `bun:"debug_sql_token8_custom,type:varchar,nullzero"`
	DebugSqlToken9Custom              string `bun:"debug_sql_token9_custom,type:varchar,nullzero"`
	DebugSqlToken10Custom             string `bun:"debug_sql_token10_custom,type:varchar,nullzero"`
	DebugSqlRebuild                   string `bun:"debug_sql_rebuild,type:varchar,nullzero"`
	DebugSqlRebuildCustom             string `bun:"debug_sql_rebuild_custom,type:varchar,nullzero"`
	DebugSqlReplaceSql                string `bun:"debug_sql_replace_sql,type:varchar,nullzero"`
	DebugSqlCustomReplaceSql          string `bun:"debug_sql_custom_replace_sql,type:varchar,nullzero"`
	DebugSqlRebuildReplaceSql         string `bun:"debug_sql_rebuild_replace_sql,type:varchar,nullzero"`
	DebugSqlRebuildCustomReplaceSql   string `bun:"debug_sql_rebuild_custom_replace_sql,type:varchar,nullzero"`
	DebugSqlP21ItemInfoReplaceSql     string `bun:"debug_sql_p21_item_info_replace_sql,type:varchar,nullzero"`
	DebugSqlExtendedRebuildReplaceSql string `bun:"debug_sql_extended_rebuild_replace_sql,type:varchar,nullzero"`
	AtomicLots                        `bun:"atomic_lots,type:bit,nullzero"`
}
