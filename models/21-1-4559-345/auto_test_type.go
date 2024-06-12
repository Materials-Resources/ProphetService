package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AutoTestType struct {
	bun.BaseModel    `bun:"table:auto_test_type"`
	AutoTestTypeUid  int32     `bun:"auto_test_type_uid,type:int,pk,identity"`
	AutoTestTypeCd   int32     `bun:"auto_test_type_cd,type:int"`
	ArgumentCount    int32     `bun:"argument_count,type:int"`
	DescArgument1    string    `bun:"desc_argument1,type:varchar(255),nullzero"`
	DescArgument2    string    `bun:"desc_argument2,type:varchar(255),nullzero"`
	DescArgument3    string    `bun:"desc_argument3,type:varchar(255),nullzero"`
	DescArgument4    string    `bun:"desc_argument4,type:varchar(255),nullzero"`
	DescArgument5    string    `bun:"desc_argument5,type:varchar(255),nullzero"`
	DescArgument6    string    `bun:"desc_argument6,type:varchar(255),nullzero"`
	DescArgument7    string    `bun:"desc_argument7,type:varchar(255),nullzero"`
	DescArgument8    string    `bun:"desc_argument8,type:varchar(255),nullzero"`
	DescArgument9    string    `bun:"desc_argument9,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
