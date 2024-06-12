package model

import (
	"github.com/uptrace/bun"
	"time"
)

type NoteArea struct {
	bun.BaseModel    `bun:"table:note_area"`
	NoteId           float64   `bun:"note_id,type:decimal(19,0),pk"`
	Area             string    `bun:"area,type:varchar(40),pk"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
