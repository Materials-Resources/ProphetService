package gen

import "github.com/uptrace/bun"

type TempCreditStatus struct {
	bun.BaseModel   `bun:"table:temp_credit_status"`
	CreditStatusId  string `bun:"credit_status_id,type:varchar(8)"`
	CreditStatusUid int32  `bun:"credit_status_uid,type:int,autoincrement,identity"`
}
