package model

type OutputAuditTrail struct {
	bun.BaseModel       `bun:"table:output_audit_trail"`
	OutputAuditTrailUid int32     `bun:"output_audit_trail_uid,type:int,pk,identity"`
	DocumentNumber      string    `bun:"document_number,type:varchar(255)"`
	DocumentType        string    `bun:"document_type,type:varchar(255)"`
	OutputType          string    `bun:"output_type,type:varchar(255)"`
	FileName            string    `bun:"file_name,type:varchar(8000)"`
	FilePath            string    `bun:"file_path,type:varchar(8000)"`
	ClientName          string    `bun:"client_name,type:varchar(255),nullzero"`
	PrinterName         string    `bun:"printer_name,type:varchar(255),nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
