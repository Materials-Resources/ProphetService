package model

type LabelDefinitionXLoc struct {
	bun.BaseModel            `bun:"table:label_definition_x_loc"`
	LabelDefinitionXLocUid   int32     `bun:"label_definition_x_loc_uid,type:int,pk"`
	LabelDefinitionUid       int32     `bun:"label_definition_uid,type:int"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	LabelProcessTypeCd       int32     `bun:"label_process_type_cd,type:int"`
	LabelTemplateFilename    string    `bun:"label_template_filename,type:varchar(255),nullzero"`
	DefaultPrintOptionCd     int32     `bun:"default_print_option_cd,type:int"`
	DefaultPrinter           string    `bun:"default_printer,type:varchar(255),nullzero"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DefaultForProcXLabelType string    `bun:"default_for_proc_x_label_type,type:char,default:('N')"`
}
