package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FieldChooserInfo struct {
	bun.BaseModel         `bun:"table:field_chooser_info"`
	FieldChooserInfoUid   int32     `bun:"field_chooser_info_uid,type:int,pk,identity"`
	Area                  int32     `bun:"area,type:int"`
	FieldName             string    `bun:"field_name,type:varchar(255)"`
	DisplayName           string    `bun:"display_name,type:varchar(255),nullzero"`
	Editable              string    `bun:"editable,type:char,nullzero"`
	FieldTypeCd           int32     `bun:"field_type_cd,type:int,nullzero"`
	FieldInfo             string    `bun:"field_info,type:varchar(255),nullzero"`
	ExtraText             string    `bun:"extra_text,type:varchar(255),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ConfigurationId       int32     `bun:"configuration_id,type:int,nullzero"`
	AutoPopulateFlag      string    `bun:"auto_populate_flag,type:char,default:('N')"`
	ValueSourceTable      string    `bun:"value_source_table,type:varchar(255),nullzero"`
	ValueSourceColumn     string    `bun:"value_source_column,type:varchar(255),nullzero"`
	ValueSourceExpression string    `bun:"value_source_expression,type:varchar(8000),nullzero"`
}
