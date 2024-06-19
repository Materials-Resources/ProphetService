package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type EccRunDate struct {
	bun.BaseModel    `bun:"table:ecc_run_date"`
	EccRunDateUid    int32     `bun:"ecc_run_date_uid,type:int,autoincrement"`
	TransactionType  string    `bun:"transaction_type,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	DateLastRun      time.Time `bun:"date_last_run,type:datetime"`
	LastSynFromDate  time.Time `bun:"last_syn_from_date,type:datetime,nullzero"`
	LastProcessedKey string    `bun:"last_processed_key,type:varchar(max),nullzero"`
	LastSynId        string    `bun:"last_syn_id,type:varchar(255),nullzero"`
}
