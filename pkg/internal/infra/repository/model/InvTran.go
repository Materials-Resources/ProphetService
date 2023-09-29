package model

import (
	"github.com/uptrace/bun"
)

type InvTran struct {
	bun.BaseModel     `bun:"table:inv_tran"`
	TransactionNumber float64 `bun:"transaction_number,type:decimal(19),pk"`
	InvMastUid        int32   `bun:"inv_mast_uid,type:int"`
	QtyAllocated      float64 `bun:"qty_allocated,type:decimal(19,9)"`
	DocumentNo        float64 `bun:"document_no,type:decimal(19)"`
	SubDocumentNo     float64 `bun:"sub_document_no,type:decimal(19)"`
	TransType         string  `bun:"trans_type,type:varchar(5)"`
}
