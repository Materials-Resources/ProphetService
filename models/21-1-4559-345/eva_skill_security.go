package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type EvaSkillSecurity struct {
	bun.BaseModel       `bun:"table:eva_skill_security"`
	EvaSkillSecurityUid int32     `bun:"eva_skill_security_uid,type:int,autoincrement,identity,pk"`    // Unique identifier for this table
	SkillName           string    `bun:"skill_name,type:varchar(8000)"`                                // Skill name for this record
	Dynachange          *string   `bun:"dynachange,type:varchar(8000)"`                                // Dynachange information for this record
	DisableFlag         string    `bun:"disable_flag,type:char(1),default:('N')"`                      // Is this record disabled
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	SkillDesc           *string   `bun:"skill_desc,type:varchar(8000)"`                                // Skill description for this record
	Version             *string   `bun:"version,type:varchar(255),default:((1.0))"`                    // Version in which this skill was added
	DynachangeDesc      *string   `bun:"dynachange_desc,type:varchar(8000)"`                           // Description of dynachanges related to this record
	WindowName          *string   `bun:"window_name,type:varchar(255)"`                                // window name that controls skill dynachange security
	Tab                 *string   `bun:"tab,type:varchar(255)"`                                        // tab that controls skill dynachange security
	TabPage             *string   `bun:"tab_page,type:varchar(255)"`                                   // tab page that controls skill dynachange security
	ClassName           *string   `bun:"class_name,type:varchar(255)"`                                 // menu class name
}
