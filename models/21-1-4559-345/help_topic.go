package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type HelpTopic struct {
	bun.BaseModel    `bun:"table:help_topic"`
	Context          string    `bun:"context,type:varchar(50),pk"`                               // The context of the help topic.  Can be a table -  datawindow -  or window
	ColumnName       string    `bun:"column_name,type:varchar(50),pk"`                           // table column name -  datawindow column name -  tab name
	Keyword          string    `bun:"keyword,type:varchar(255)"`                                 // Keyword to be used when accessing help
	Topic            int32     `bun:"topic,type:int,nullzero"`                                   // The topic of the note for the referenced area.
	AccessCounter    float64   `bun:"access_counter,type:decimal(19,0)"`                         // Is used to count how many times a specific keyword is used to lookup help
	Verified         string    `bun:"verified,type:char(1),nullzero"`                            // Shows status of record verification
	TopicType        string    `bun:"topic_type,type:varchar(10),nullzero"`                      // Type of Topic (TABLE -  DW -  TAB -  WINDOW)
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
