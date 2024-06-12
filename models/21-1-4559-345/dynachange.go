package model

type Dynachange struct {
	bun.BaseModel     `bun:"table:dynachange"`
	DynachangeId      float64   `bun:"dynachange_id,type:decimal(6,0),pk"`
	BaseClass         string    `bun:"base_class,type:varchar(255)"`
	PersonalizedClass string    `bun:"personalized_class,type:varchar(255)"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
