package model

type ScanPack struct {
	bun.BaseModel            `bun:"table:scan_pack"`
	ScanPackUid              int32     `bun:"scan_pack_uid,type:int,pk"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SourceCd                 int32     `bun:"source_cd,type:int,default:((3612))"`
	PrntUcc128LblDisplayFlag string    `bun:"prnt_ucc128_lbl_display_flag,type:char,default:('N')"`
	LockFlag                 string    `bun:"lock_flag,type:char,default:('N')"`
}
