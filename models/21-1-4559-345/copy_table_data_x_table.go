package model

type CopyTableDataXTable struct {
	bun.BaseModel            `bun:"table:copy_table_data_x_table"`
	CopyTableDataXTableUid   int32     `bun:"copy_table_data_x_table_uid,type:int,pk,identity"`
	CopyTableDataXProcessUid int32     `bun:"copy_table_data_x_process_uid,type:int"`
	TableName                string    `bun:"table_name,type:varchar(255)"`
	SourceTableName          string    `bun:"source_table_name,type:varchar(255),nullzero"`
	SequenceNumber           int32     `bun:"sequence_number,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseReadUncommitted       string    `bun:"use_read_uncommitted,type:char,default:('N')"`
}
