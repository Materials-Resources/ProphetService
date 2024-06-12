package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Country struct {
	bun.BaseModel    `bun:"table:country"`
	CountryUid       int32     `bun:"country_uid,type:int,pk"`
	CountryNo        int32     `bun:"country_no,type:int"`
	TwoLetterCode    string    `bun:"two_letter_code,type:char(2)"`
	ThreeLetterCode  string    `bun:"three_letter_code,type:char(3)"`
	CountryName      string    `bun:"country_name,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
