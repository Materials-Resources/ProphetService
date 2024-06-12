package model

import (
	"github.com/uptrace/bun"
	"time"
)

type BinType struct {
	bun.BaseModel    `bun:"table:bin_type"`
	BinTypeUid       int32     `bun:"bin_type_uid,type:int,pk"`
	BinType          string    `bun:"bin_type,type:varchar(255)"`
	BinTypeDesc      string    `bun:"bin_type_desc,type:varchar(255)"`
	CompanyId        string    `bun:"company_id,type:varchar(8)"`
	PutableFlag      string    `bun:"putable_flag,type:char"`
	PickableFlag     string    `bun:"pickable_flag,type:char"`
	QuarantineFlag   string    `bun:"quarantine_flag,type:char"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	WeighStationFlag string    `bun:"weigh_station_flag,type:char,default:('N')"`
	FrontCounterFlag string    `bun:"front_counter_flag,type:char,nullzero"`
	BackflushFlag    string    `bun:"backflush_flag,type:char,default:('N')"`
	FullSkidOnlyFlag string    `bun:"full_skid_only_flag,type:char,nullzero"`
	ShippableFlag    string    `bun:"shippable_flag,type:char,nullzero"`
}
