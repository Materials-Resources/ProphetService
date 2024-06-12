package model

import (
	"github.com/uptrace/bun"
	"time"
)

type EvaSkillSecurity struct {
	bun.BaseModel       `bun:"table:eva_skill_security"`
	EvaSkillSecurityUid int32     `bun:"eva_skill_security_uid,type:int,pk,identity"`
	SkillName           string    `bun:"skill_name,type:varchar(8000)"`
	Dynachange          string    `bun:"dynachange,type:varchar(8000),nullzero"`
	DisableFlag         string    `bun:"disable_flag,type:char,default:('N')"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SkillDesc           string    `bun:"skill_desc,type:varchar(8000),nullzero"`
	Version             string    `bun:"version,type:varchar(255),default:((1.0))"`
	DynachangeDesc      string    `bun:"dynachange_desc,type:varchar(8000),nullzero"`
	WindowName          string    `bun:"window_name,type:varchar(255),nullzero"`
	Tab                 string    `bun:"tab,type:varchar(255),nullzero"`
	TabPage             string    `bun:"tab_page,type:varchar(255),nullzero"`
	ClassName           string    `bun:"class_name,type:varchar(255),nullzero"`
}
