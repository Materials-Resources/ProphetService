package models

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type OeLinePromiseDate struct {
	bun.BaseModel                   `bun:"table:oe_line_promise_date"`
	OeLineUid                       int32          `bun:"oe_line_uid,type:int"`
	OriginalPromiseDate             sql.NullTime   `bun:"original_promise_date,type:datetime,nullzero"`
	PromiseDate                     time.Time      `bun:"promise_date,type:datetime"`
	EditedFlag                      string         `bun:"edited_flag,type:char"`
	AroDays                         sql.NullInt32  `bun:"aro_days,type:int,nullzero"`
	DateCreated                     time.Time      `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                       string         `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified                time.Time      `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy                string         `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	OeLinePromiseDateUid            int32          `bun:"oe_line_promise_date_uid,type:int,pk"`
	PromiseDateEditedDate           sql.NullTime   `bun:"promise_date_edited_date,type:datetime,nullzero"`
	ExtendedDesc                    sql.NullString `bun:"extended_desc,type:varchar(8000),nullzero"`
	Alert                           sql.NullString `bun:"alert,type:varchar,nullzero"`
	PreviousPromiseDate             sql.NullTime   `bun:"previous_promise_date,type:datetime,nullzero"`
	InsufficientDataToCalculateFlag sql.NullString `bun:"insufficient_data_to_calculate_flag,type:char,nullzero"`
	Edi855sResendPromiseDate        sql.NullTime   `bun:"edi855s_resend_promise_date,type:datetime,nullzero"`
}
