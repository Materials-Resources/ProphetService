package model

type CustPartNoGroupHdr struct {
	bun.BaseModel         `bun:"table:cust_part_no_group_hdr"`
	CustPartNoGroupHdrUid int32     `bun:"cust_part_no_group_hdr_uid,type:int,pk"`
	CustPartNoGroupNo     string    `bun:"cust_part_no_group_no,type:varchar(255)"`
	CustPartNoGroupDesc   string    `bun:"cust_part_no_group_desc,type:varchar(255)"`
	CompanyId             string    `bun:"company_id,type:varchar(8)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
