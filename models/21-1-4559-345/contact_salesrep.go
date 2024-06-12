package model

type ContactSalesrep struct {
	bun.BaseModel      `bun:"table:contact_salesrep"`
	ContactId          string    `bun:"contact_id,type:varchar(16)"`
	SalesrepId         string    `bun:"salesrep_id,type:varchar(16)"`
	DeleteFlag         string    `bun:"delete_flag,type:char"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ContactSalesrepUid int32     `bun:"contact_salesrep_uid,type:int,pk"`
}
