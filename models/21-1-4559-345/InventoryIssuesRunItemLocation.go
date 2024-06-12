package model

import "github.com/uptrace/bun"

type Inventoryissuesrunitemlocation struct {
	bun.BaseModel `bun:"table:InventoryIssuesRunItemLocation"`
	Run           int32 `bun:"Run,type:int"`
	InvMastUid    int32 `bun:"inv_mast_uid,type:int"`
	LocationId    int32 `bun:"location_id,type:int"`
	Instances     int32 `bun:"instances,type:int,default:((1))"`
	Testuid       int32 `bun:"TestUID,type:int,nullzero"`
	TableUid      int32 `bun:"table_uid,type:int,identity"`
	TotalTestUid  int32 `bun:"total_test_uid,type:int,nullzero"`
}
