package model

type Salutation struct {
	bun.BaseModel    `bun:"table:salutation"`
	Salutation       string    `bun:"salutation,type:varchar(4),pk"`
	FullName         string    `bun:"full_name,type:varchar(30)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint,default:(704)"`
}
