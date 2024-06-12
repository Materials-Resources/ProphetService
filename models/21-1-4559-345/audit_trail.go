package model

type AuditTrail struct {
	bun.BaseModel     `bun:"table:audit_trail"`
	AuditTrailUid     int32     `bun:"audit_trail_uid,type:int,pk,identity"`
	SourceAreaCd      int32     `bun:"source_area_cd,type:int"`
	ColumnChanged     string    `bun:"column_changed,type:varchar(255)"`
	Key1Cd            string    `bun:"key1_cd,type:varchar(255)"`
	Key1Value         string    `bun:"key1_value,type:varchar(255)"`
	Key2Cd            string    `bun:"key2_cd,type:varchar(255),nullzero"`
	Key2Value         string    `bun:"key2_value,type:varchar(255),nullzero"`
	Key3Cd            string    `bun:"key3_cd,type:varchar(255),nullzero"`
	Key3Value         string    `bun:"key3_value,type:varchar(255),nullzero"`
	LineNo            int32     `bun:"line_no,type:int,nullzero"`
	InvMastUid        int32     `bun:"inv_mast_uid,type:int,nullzero"`
	OldValue          string    `bun:"old_value,type:varchar(255),nullzero"`
	NewValue          string    `bun:"new_value,type:varchar(255),nullzero"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TableChanged      string    `bun:"table_changed,type:varchar(255),nullzero"`
	UidValue          int32     `bun:"uid_value,type:int,nullzero"`
	ColumnDescription string    `bun:"column_description,type:varchar(255),nullzero"`
	AuxiliaryValue    string    `bun:"auxiliary_value,type:varchar(255),nullzero"`
}
