package model

type ImportSuspenseLine struct {
	bun.BaseModel         `bun:"table:import_suspense_line"`
	ImportSuspenseLineUid int32     `bun:"import_suspense_line_uid,type:int,pk"`
	ImportSuspenseHdrUid  int32     `bun:"import_suspense_hdr_uid,type:int"`
	SequenceNo            int32     `bun:"sequence_no,type:int"`
	ImportFilePath        string    `bun:"import_file_path,type:varchar(255)"`
	ImportFileName        string    `bun:"import_file_name,type:varchar(255)"`
	BoInstancename        string    `bun:"bo_instancename,type:varchar(255),nullzero"`
	DsDataobject          string    `bun:"ds_dataobject,type:varchar(255)"`
	DwDataobject          string    `bun:"dw_dataobject,type:varchar(255)"`
	ImportSetNo           string    `bun:"import_set_no,type:varchar(255)"`
	CompanyId             string    `bun:"company_id,type:varchar(8),nullzero"`
	CustomerId            float64   `bun:"customer_id,type:decimal(19,0),nullzero"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	KeyValue              string    `bun:"key_value,type:varchar(255),nullzero"`
	KeyValueTable         string    `bun:"key_value_table,type:varchar(255),nullzero"`
	ExportedStatus        int32     `bun:"exported_status,type:int,default:(0)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int,default:(1183)"`
	DisplayName           string    `bun:"display_name,type:varchar(255),default:('Unknown')"`
	SuspenseData          string    `bun:"suspense_data,type:text(2147483647)"`
	ErrorData             string    `bun:"error_data,type:text(2147483647)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	SourceId              string    `bun:"source_id,type:varchar(8),nullzero"`
}
