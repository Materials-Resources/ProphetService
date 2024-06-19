package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Country struct {
	bun.BaseModel    `bun:"table:country"`
	CountryUid       int32     `bun:"country_uid,type:int,pk"`                                      // Unique identifier for this record
	CountryNo        int32     `bun:"country_no,type:int,unique"`                                   // Numeric ISO country code
	TwoLetterCode    string    `bun:"two_letter_code,type:char(2)"`                                 // Two letter ISO country code
	ThreeLetterCode  string    `bun:"three_letter_code,type:char(3)"`                               // Three letter ISO country code
	CountryName      string    `bun:"country_name,type:varchar(255)"`                               // The name of the country, in English
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
