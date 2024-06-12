package model

type AddressHistory struct {
	bun.BaseModel     `bun:"table:address_history"`
	AddressHistoryUid int32     `bun:"address_history_uid,type:int,pk"`
	CustomerName      string    `bun:"customer_name,type:varchar(50),nullzero"`
	ContactName       string    `bun:"contact_name,type:varchar(30),nullzero"`
	Address1          string    `bun:"address1,type:varchar(50),nullzero"`
	Address2          string    `bun:"address2,type:varchar(50),nullzero"`
	State             string    `bun:"state,type:varchar(50),nullzero"`
	PostalCode        string    `bun:"postal_code,type:char(10),nullzero"`
	Country           string    `bun:"country,type:varchar(50),nullzero"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	City              string    `bun:"city,type:varchar(50),nullzero"`
	Address3          string    `bun:"address3,type:varchar(50),nullzero"`
}
