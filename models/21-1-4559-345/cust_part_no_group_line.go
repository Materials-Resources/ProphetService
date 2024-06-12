package model

type CustPartNoGroupLine struct {
	bun.BaseModel          `bun:"table:cust_part_no_group_line"`
	CustPartNoGroupLineUid int32     `bun:"cust_part_no_group_line_uid,type:int,pk"`
	CustPartNoGroupHdrUid  int32     `bun:"cust_part_no_group_hdr_uid,type:int"`
	CustomerId             float64   `bun:"customer_id,type:decimal(19,0)"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
