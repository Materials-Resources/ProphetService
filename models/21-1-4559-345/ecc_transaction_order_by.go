package model

type EccTransactionOrderBy struct {
	bun.BaseModel            `bun:"table:ecc_transaction_order_by"`
	EccTransactionOrderByUid int32 `bun:"ecc_transaction_order_by_uid,type:int,pk,identity"`
	TransactionType          `bun:"transaction_type,type:nvarchar(255)"`
	OrderBy                  `bun:"order_by,type:nvarchar"`
}
