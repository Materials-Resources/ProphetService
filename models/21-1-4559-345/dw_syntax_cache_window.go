package model

type DwSyntaxCacheWindow struct {
	bun.BaseModel          `bun:"table:dw_syntax_cache_window"`
	DwSyntaxCacheWindowUid int32     `bun:"dw_syntax_cache_window_uid,type:int,pk,identity"`
	WindowName             string    `bun:"window_name,type:varchar(255)"`
	EnableCachingFlag      string    `bun:"enable_caching_flag,type:char"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
