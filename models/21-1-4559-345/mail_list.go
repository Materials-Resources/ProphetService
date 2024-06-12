package model

type MailList struct {
	bun.BaseModel    `bun:"table:mail_list"`
	ListId           float64   `bun:"list_id,type:decimal(19,0),pk"`
	ListDesc         string    `bun:"list_desc,type:varchar(40),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
}
