package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportTempTableXFile struct {
	bun.BaseModel             `bun:"table:import_temp_table_x_file"`
	ImportTempTableXFileUid   int32     `bun:"import_temp_table_x_file_uid,type:int,pk"`
	TempTableName             string    `bun:"temp_table_name,type:varchar(255)"`
	FileName                  string    `bun:"file_name,type:varchar(255)"`
	ImportDisplayName         string    `bun:"import_display_name,type:varchar(255)"`
	ImportTargetTable         string    `bun:"import_target_table,type:varchar(255)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int,default:(704)"`
	DatabaseVersion           string    `bun:"database_version,type:varchar(255),nullzero"`
	ImportFileDate            time.Time `bun:"import_file_date,type:datetime,nullzero"`
	DsDataobject              string    `bun:"ds_dataobject,type:varchar(255),nullzero"`
	LastVisibleColumnId       int32     `bun:"last_visible_column_id,type:int,nullzero"`
	ValStatusFlag             int32     `bun:"val_status_flag,type:int,nullzero"`
	SelectString              string    `bun:"select_string,type:text(2147483647),nullzero"`
	DecimalInitialValueString string    `bun:"decimal_initial_value_string,type:text(2147483647),nullzero"`
	LegacyTempTableName       string    `bun:"legacy_temp_table_name,type:varchar(255),nullzero"`
	ProcessTypeCd             int32     `bun:"process_type_cd,type:int,default:((707))"`
}
