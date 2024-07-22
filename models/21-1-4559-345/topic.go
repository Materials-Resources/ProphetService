package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Topic struct {
	bun.BaseModel    `bun:"table:topic"`
	TopicUid         int32     `bun:"topic_uid,type:int,pk"`               // Unique identifier for the topic
	TopicId          *string   `bun:"topic_id,type:varchar(255),unique"`   // This is a brief description of the topic of the call that allows you to quickly identify the purpose of a particular call.
	TopicDescription *string   `bun:"topic_description,type:varchar(255)"` // The brief description of a topic ID with which it is associated.  If a valid topic ID is entered, then the corresponding topic description will be displayed.
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint"`       // Indicates current record status.
	DateCreated      time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
}
