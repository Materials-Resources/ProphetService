package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FormDestinationHierarchy struct {
	bun.BaseModel               `bun:"table:form_destination_hierarchy"`
	FormDestinationHierarchyUid int32     `bun:"form_destination_hierarchy_uid,type:int,pk"`
	FormDestinationUid          int32     `bun:"form_destination_uid,type:int"`
	SequenceNo                  int32     `bun:"sequence_no,type:int"`
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
