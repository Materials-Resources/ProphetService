package model

type CopyTableDataXProcess struct {
	bun.BaseModel            `bun:"table:copy_table_data_x_process"`
	CopyTableDataXProcessUid int32     `bun:"copy_table_data_x_process_uid,type:int,pk,identity"`
	ProcessTypeCd            int32     `bun:"process_type_cd,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
