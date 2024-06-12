package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Topic struct {
	bun.BaseModel    `bun:"table:topic"`
	TopicUid         int32     `bun:"topic_uid,type:int,pk"`
	TopicId          string    `bun:"topic_id,type:varchar(255),nullzero"`
	TopicDescription string    `bun:"topic_description,type:varchar(255),nullzero"`
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
}
