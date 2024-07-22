package prophet

import "github.com/uptrace/bun"

type ApinvHdrNotDelete struct {
	bun.BaseModel `bun:"table:apinv_hdr_not_delete"`
	VoucherNo     string `bun:"voucher_no,type:varchar(20)"`
	Reason        string `bun:"reason,type:varchar(255)"`
}
