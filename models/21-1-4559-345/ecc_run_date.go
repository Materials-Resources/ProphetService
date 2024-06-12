package model

type EccRunDate struct {
	bun.BaseModel    `bun:"table:ecc_run_date"`
	EccRunDateUid    int32 `bun:"ecc_run_date_uid,type:int,identity"`
	TransactionType  `bun:"transaction_type,type:nvarchar(255)"`
	DateLastRun      time.Time `bun:"date_last_run,type:datetime"`
	LastSynFromDate  time.Time `bun:"last_syn_from_date,type:datetime,nullzero"`
	LastProcessedKey string    `bun:"last_processed_key,type:varchar,nullzero"`
	LastSynId        string    `bun:"last_syn_id,type:varchar(255),nullzero"`
}
