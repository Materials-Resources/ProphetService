package model

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessRuleEventKey struct {
	bun.BaseModel           `bun:"table:business_rule_event_key"`
	BusinessRuleEventKeyUid int32     `bun:"business_rule_event_key_uid,type:int,pk,identity"`
	BusinessRuleEventUid    int32     `bun:"business_rule_event_uid,type:int"`
	KeyValue                string    `bun:"key_value,type:varchar(255)"`
	DisplayValue            string    `bun:"display_value,type:varchar(255)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
