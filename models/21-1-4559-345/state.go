package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type State struct {
	bun.BaseModel              `bun:"table:state"`
	StateUid                   int32     `bun:"state_uid,type:int,pk"`                                        // UID for the state
	CountryUid                 int32     `bun:"country_uid,type:int,unique"`                                  // Country for this state/province
	TwoLetterCode              string    `bun:"two_letter_code,type:char(2),unique"`                          // Two letter ISO state code
	StateName                  string    `bun:"state_name,type:varchar(255)"`                                 // State name
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	CombinedFederalState1099No *int32    `bun:"combined_federal_state_1099_no,type:int"`                      // State code for combined federal/state 1099 filing
	TelecheckStateCode         *int32    `bun:"telecheck_state_code,type:int"`                                // Telecheck (protobase integration) has a special reference number for each state.
}
