package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type WeboeFile struct {
	bun.BaseModel     `bun:"table:weboe_file"`
	WeboeFileUid      int32     `bun:"weboe_file_uid,type:int,autoincrement"`
	WeboeFileName     string    `bun:"weboe_file_name,type:varchar(255)"`
	WeboeOldFileName  string    `bun:"weboe_old_file_name,type:varchar(255)"`
	WeboeFileDesc     string    `bun:"weboe_file_desc,type:varchar(255)"`
	WeboeLastTransfer time.Time `bun:"weboe_last_transfer,type:datetime"`
	PartialTransfer   bool      `bun:"partial_transfer,type:bit,nullzero"`
	TransferOrder     int16     `bun:"transfer_order,type:smallint,nullzero"`
	FileSql           string    `bun:"file_sql,type:varchar(8000)"`
}
