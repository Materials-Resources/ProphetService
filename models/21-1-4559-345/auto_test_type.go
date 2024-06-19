package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AutoTestType struct {
	bun.BaseModel    `bun:"table:auto_test_type"`
	AutoTestTypeUid  int32     `bun:"auto_test_type_uid,type:int,autoincrement,identity,pk"`        // Unique identifier for auto_test_type table.
	AutoTestTypeCd   int32     `bun:"auto_test_type_cd,type:int,unique"`                            // System code value describing a test action type.
	ArgumentCount    int32     `bun:"argument_count,type:int"`                                      // Number of logical arguments needed to execute this type of test action.
	DescArgument1    string    `bun:"desc_argument1,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument2    string    `bun:"desc_argument2,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument3    string    `bun:"desc_argument3,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument4    string    `bun:"desc_argument4,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument5    string    `bun:"desc_argument5,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument6    string    `bun:"desc_argument6,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument7    string    `bun:"desc_argument7,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument8    string    `bun:"desc_argument8,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DescArgument9    string    `bun:"desc_argument9,type:varchar(255),nullzero"`                    // Description of what the argument is for the test action type (if applicable).
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
