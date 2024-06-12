package model

type State struct {
	bun.BaseModel              `bun:"table:state"`
	StateUid                   int32     `bun:"state_uid,type:int,pk"`
	CountryUid                 int32     `bun:"country_uid,type:int"`
	TwoLetterCode              string    `bun:"two_letter_code,type:char(2)"`
	StateName                  string    `bun:"state_name,type:varchar(255)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CombinedFederalState1099No int32     `bun:"combined_federal_state_1099_no,type:int,nullzero"`
	TelecheckStateCode         int32     `bun:"telecheck_state_code,type:int,nullzero"`
}
