package model

import (
	"github.com/uptrace/bun"
	"time"
)

type HelpTopic struct {
	bun.BaseModel    `bun:"table:help_topic"`
	Context          string    `bun:"context,type:varchar(50),pk"`
	ColumnName       string    `bun:"column_name,type:varchar(50),pk"`
	Keyword          string    `bun:"keyword,type:varchar(255)"`
	Topic            int32     `bun:"topic,type:int,nullzero"`
	AccessCounter    float64   `bun:"access_counter,type:decimal(19,0)"`
	Verified         string    `bun:"verified,type:char,nullzero"`
	TopicType        string    `bun:"topic_type,type:varchar(10),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
