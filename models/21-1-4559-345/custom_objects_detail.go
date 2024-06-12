package model

type CustomObjectsDetail struct {
	bun.BaseModel          `bun:"table:custom_objects_detail"`
	CustomObjectsDetailUid int32     `bun:"custom_objects_detail_uid,type:int,pk"`
	CustomObjectsUid       int32     `bun:"custom_objects_uid,type:int"`
	SequenceNo             int32     `bun:"sequence_no,type:int"`
	ObjectName             string    `bun:"object_name,type:varchar(255)"`
	AttributeName          string    `bun:"attribute_name,type:varchar(255)"`
	AttributeValue         string    `bun:"attribute_value,type:varchar(2048)"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
