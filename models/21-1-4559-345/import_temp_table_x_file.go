package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportTempTableXFile struct {
	bun.BaseModel             `bun:"table:import_temp_table_x_file"`
	ImportTempTableXFileUid   int32      `bun:"import_temp_table_x_file_uid,type:int,pk"`                     // Unique identifier for each record
	TempTableName             string     `bun:"temp_table_name,type:varchar(255)"`                            // name of temporary table holding import file data
	FileName                  string     `bun:"file_name,type:varchar(255)"`                                  // name of import file
	ImportDisplayName         string     `bun:"import_display_name,type:varchar(255)"`                        // display name of import type
	ImportTargetTable         string     `bun:"import_target_table,type:varchar(255)"`                        // CommerceCenter the import data will update
	DateCreated               time.Time  `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string     `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag             int32      `bun:"row_status_flag,type:int,default:(704)"`                       // import data temp table status
	DatabaseVersion           *string    `bun:"database_version,type:varchar(255)"`                           // system_setting database version when temp import table created
	ImportFileDate            *time.Time `bun:"import_file_date,type:datetime"`                               // date import file was last modified
	DsDataobject              *string    `bun:"ds_dataobject,type:varchar(255)"`                              // definition of import layout
	LastVisibleColumnId       *int32     `bun:"last_visible_column_id,type:int"`                              // column number of last import layout column
	ValStatusFlag             *int32     `bun:"val_status_flag,type:int"`                                     // pre-import validation status for file data
	SelectString              *string    `bun:"select_string,type:text"`                                      // SQL string used to retrieve import data from temp table
	DecimalInitialValueString *string    `bun:"decimal_initial_value_string,type:text"`                       // string to set initial values, used to create decimal default values
	LegacyTempTableName       *string    `bun:"legacy_temp_table_name,type:varchar(255)"`                     // temp table name for legacy import data
	ProcessTypeCd             int32      `bun:"process_type_cd,type:int,default:((707))"`                     // Process Type
}
