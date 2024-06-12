package model

type FormDestination struct {
	bun.BaseModel      `bun:"table:form_destination"`
	FormDestinationUid int32     `bun:"form_destination_uid,type:int,pk"`
	FormCd             int32     `bun:"form_cd,type:int"`
	OutputCd           int32     `bun:"output_cd,type:int"`
	SourceCd           int32     `bun:"source_cd,type:int"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int"`
}
