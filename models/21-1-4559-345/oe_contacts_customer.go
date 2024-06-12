package model

type OeContactsCustomer struct {
	bun.BaseModel            `bun:"table:oe_contacts_customer"`
	CompanyId                string    `bun:"company_id,type:varchar(8),pk"`
	CustomerId               float64   `bun:"customer_id,type:decimal(19,0),pk"`
	ContactId                string    `bun:"contact_id,type:varchar(16),pk"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ShopperUid               int32     `bun:"shopper_uid,type:int,default:(0)"`
	StatementContact         string    `bun:"statement_contact,type:char,default:('N')"`
	CreditCardNo             string    `bun:"credit_card_no,type:varchar(255),nullzero"`
	CreditCardType           string    `bun:"credit_card_type,type:varchar(255),nullzero"`
	CreditCardName           string    `bun:"credit_card_name,type:varchar(255),nullzero"`
	CreditCardExpirationDate time.Time `bun:"credit_card_expiration_date,type:datetime,nullzero"`
	PedigreeContact          string    `bun:"pedigree_contact,type:char,default:('N')"`
}
