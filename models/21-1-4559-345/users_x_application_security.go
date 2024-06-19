package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type UsersXApplicationSecurity struct {
	bun.BaseModel          `bun:"table:users_x_application_security"`
	UsersXAppSecurityUid   int32     `bun:"users_x_app_security_uid,type:int,autoincrement,identity,pk"`  // Unique ID, primary key
	UsersId                string    `bun:"users_id,type:varchar(30)"`                                    // ID from user table, FK
	ApplicationSecurityUid int32     `bun:"application_security_uid,type:int"`                            // UID from application_security table, FK
	CodeValue              int32     `bun:"code_value,type:int,nullzero"`                                 // Code (int) value for this record. Will contain only codes for code_p21
	DecimalValue           float64   `bun:"decimal_value,type:decimal(19,9),nullzero"`                    // Decimal value for this record.
	StringValue            string    `bun:"string_value,type:varchar(255),nullzero"`                      // String value for this record.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
