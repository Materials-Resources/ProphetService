package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CodeP21 struct {
	bun.BaseModel      `bun:"table:code_p21"`
	CodeUid            int32     `bun:"code_uid,type:int,pk"`
	CodeNo             int32     `bun:"code_no,type:int"`
	LanguageId         string    `bun:"language_id,type:varchar(8)"`
	CodeDescription    string    `bun:"code_description,type:varchar(255)"`
	RowStatusFlag      string    `bun:"row_status_flag,type:char(1),default:('A')"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	CodeSubDescription string    `bun:"code_sub_description,type:varchar(255),nullzero"`
}
