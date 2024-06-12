package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportSuspenseFkerror struct {
	bun.BaseModel            `bun:"table:import_suspense_fkerror"`
	ImportSuspenseFkerrorUid int32     `bun:"import_suspense_fkerror_uid,type:int,pk"`
	ImportSuspenseHdrUid     int32     `bun:"import_suspense_hdr_uid,type:int"`
	FkTablename              string    `bun:"fk_tablename,type:varchar(255)"`
	FkError                  string    `bun:"fk_error,type:text(2147483647)"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
