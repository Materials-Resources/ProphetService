package model

type OeSchedule struct {
	bun.BaseModel     `bun:"table:oe_schedule"`
	OrderNumber       string    `bun:"order_number,type:varchar(8),pk"`
	ReleaseType       string    `bun:"release_type,type:varchar(30)"`
	RoundType         string    `bun:"round_type,type:varchar(30)"`
	FirstDate         time.Time `bun:"first_date,type:datetime,nullzero"`
	TotalReleases     float64   `bun:"total_releases,type:decimal(19,0),nullzero"`
	FrequencyValue    float64   `bun:"frequency_value,type:decimal(19,0),nullzero"`
	FrequencyType     string    `bun:"frequency_type,type:varchar(8),nullzero"`
	ExpediteValue     float64   `bun:"expedite_value,type:decimal(19,0),nullzero"`
	ExpediteType      string    `bun:"expedite_type,type:varchar(8),nullzero"`
	PickValue         float64   `bun:"pick_value,type:decimal(19,0),nullzero"`
	PickType          string    `bun:"pick_type,type:varchar(8),nullzero"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DefaultToAll      string    `bun:"default_to_all,type:char"`
	ApplyToLineMethod int32     `bun:"apply_to_line_method,type:int,nullzero"`
}
