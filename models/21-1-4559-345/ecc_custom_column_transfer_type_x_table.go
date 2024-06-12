package model

type EccCustomColumnTransferTypeXTable struct {
	bun.BaseModel                        `bun:"table:ecc_custom_column_transfer_type_x_table"`
	EccCustomColumnTransferTypeXTableUid int32 `bun:"ecc_custom_column_transfer_type_x_table_uid,type:int,pk,identity"`
	EccTransferType                      `bun:"ecc_transfer_type,type:nvarchar(255)"`
	UdTableName                          `bun:"ud_table_name,type:nvarchar(255)"`
	JoinSql                              `bun:"join_sql,type:nvarchar,nullzero"`
	DateCreated                          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified                     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy                     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SbUpdateJoin                         `bun:"sb_update_join,type:nvarchar,nullzero"`
	SbInsertSelect                       `bun:"sb_insert_select,type:nvarchar,nullzero"`
	ApplyTableJoins                      string `bun:"apply_table_joins,type:char,nullzero"`
	TriggerKeyColumn                     `bun:"trigger_key_column,type:nvarchar,nullzero"`
	TriggerDeletedJoin                   `bun:"trigger_deleted_join,type:nvarchar,nullzero"`
	ProcedureKeyColumn                   `bun:"procedure_key_column,type:nvarchar,nullzero"`
}
