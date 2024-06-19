package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Token struct {
	bun.BaseModel    `bun:"table:token"`
	TokenUid         int32     `bun:"token_uid,type:int,pk"`                                     // Unique Identifier of token table.
	Name             string    `bun:"name,type:varchar(40),unique"`                              // Name of the token.
	Description      string    `bun:"description,type:varchar(80),unique"`                       // Description of the token.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DataTypeCd       int32     `bun:"data_type_cd,type:int,nullzero"`                            // Store a code that will let the program know how to build a where clause.
	AvailableAreas   int32     `bun:"available_areas,type:int,unique,nullzero"`                  // Available areas in the system for this token
	CodeGroupNo      int16     `bun:"code_group_no,type:smallint,nullzero"`                      // System assigned code for the code group that relates to this token.
	UserLookupFlag   string    `bun:"user_lookup_flag,type:char(1),nullzero"`                    // Indicates that the token represents a field that will be used as part of a query to find other associated data.
}
