package model

type Inventoryissuesrun struct {
	bun.BaseModel    `bun:"table:InventoryIssuesRun"`
	Run              int32     `bun:"Run,type:int,pk"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int,nullzero"`
	LocationId       int32     `bun:"location_id,type:int,nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255)"`
	RunByUser        string    `bun:"run_by_user,type:varchar(255),nullzero"`
	DateRunEnded     time.Time `bun:"date_run_ended,type:datetime,nullzero"`
	DateRunStarted   time.Time `bun:"date_run_started,type:datetime,nullzero"`
	Server           string    `bun:"server,type:varchar(255),nullzero"`
	DbName           string    `bun:"db_name,type:varchar(255),nullzero"`
	DbVersion        string    `bun:"db_version,type:varchar(255),nullzero"`
	Testuid          int32     `bun:"TestUID,type:int,nullzero"`
	Instances        int32     `bun:"instances,type:int,nullzero"`
	DateDiffMs       int32     `bun:"date_diff_ms,type:int,nullzero"`
	TotalTestUid     int32     `bun:"total_test_uid,type:int,nullzero"`
}
