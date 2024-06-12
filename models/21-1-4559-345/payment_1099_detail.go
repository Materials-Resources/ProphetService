package model

type Payment1099Detail struct {
	bun.BaseModel        `bun:"table:payment_1099_detail"`
	CompanyId            string    `bun:"company_id,type:varchar(8)"`
	VendorId             float64   `bun:"vendor_id,type:decimal(19,0)"`
	Year                 float64   `bun:"year,type:decimal(4,0)"`
	BankNo               float64   `bun:"bank_no,type:decimal(19,0)"`
	CheckNo              string    `bun:"check_no,type:varchar(255)"`
	Ten99Type            float64   `bun:"ten99_type,type:decimal(19,0)"`
	Amount               float64   `bun:"amount,type:decimal(19,4)"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Payment1099DetailUid int32     `bun:"payment_1099_detail_uid,type:int,pk"`
}
