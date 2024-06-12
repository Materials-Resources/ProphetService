package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PoScheduleRule struct {
	bun.BaseModel    `bun:"table:po_schedule_rule"`
	PoHdrUid         int32     `bun:"po_hdr_uid,type:int,pk"`
	ReleaseType      int32     `bun:"release_type,type:int"`
	RoundType        int32     `bun:"round_type,type:int"`
	FirstDate        time.Time `bun:"first_date,type:datetime,nullzero"`
	TotalReleases    int32     `bun:"total_releases,type:int,nullzero"`
	FrequencyValue   int32     `bun:"frequency_value,type:int,nullzero"`
	FrequencyType    string    `bun:"frequency_type,type:varchar(8),nullzero"`
	DefaultToAll     string    `bun:"default_to_all,type:char,default:('N')"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
