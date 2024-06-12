package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLink struct {
	bun.BaseModel    `bun:"table:document_link"`
	DocumentLinkUid  int32     `bun:"document_link_uid,type:int,pk"`
	SourceAreaCd     int32     `bun:"source_area_cd,type:int"`
	Key1Cd           string    `bun:"key1_cd,type:varchar(255)"`
	Key1Value        string    `bun:"key1_value,type:varchar(255)"`
	Key2Cd           string    `bun:"key2_cd,type:varchar(255),nullzero"`
	Key2Value        string    `bun:"key2_value,type:varchar(255),nullzero"`
	Key3Cd           string    `bun:"key3_cd,type:varchar(255),nullzero"`
	Key3Value        string    `bun:"key3_value,type:varchar(255),nullzero"`
	LinkName         string    `bun:"link_name,type:varchar(255)"`
	LinkPath         string    `bun:"link_path,type:varchar(4099)"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	OutsideUseFlag   string    `bun:"outside_use_flag,type:char,default:('N')"`
	MandatoryFlag    string    `bun:"mandatory_flag,type:char,default:('N')"`
}
