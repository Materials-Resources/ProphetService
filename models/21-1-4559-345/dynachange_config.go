package model

type DynachangeConfig struct {
	bun.BaseModel    `bun:"table:dynachange_config"`
	ConfigurationId  int32     `bun:"configuration_id,type:int,pk"`
	DynachangeId     float64   `bun:"dynachange_id,type:decimal(6,0),pk"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30)"`
}
