package model

type CopyTableDataXCounter struct {
	bun.BaseModel            `bun:"table:copy_table_data_x_counter"`
	CopyTableDataXCounterUid int32     `bun:"copy_table_data_x_counter_uid,type:int,pk,identity"`
	CopyTableDataXTableUid   int32     `bun:"copy_table_data_x_table_uid,type:int"`
	CounterId                string    `bun:"counter_id,type:varchar(64)"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
