package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type NoteDisplayArea struct {
	bun.BaseModel       `bun:"table:note_display_area"`
	NoteDisplayAreasUid int32     `bun:"note_display_areas_uid,type:int,autoincrement,identity,pk"`    // Unique identifier for a record
	TopicId             int32     `bun:"topic_id,type:int"`                                            // Numeric identifier to group all areas (records) where a given note will display
	DisplayArea         string    `bun:"display_area,type:varchar(40)"`                                // Name of the Area where a note will be displayed
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
