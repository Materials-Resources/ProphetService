package model

type PcVersion struct {
	bun.BaseModel  `bun:"table:pc_version"`
	ProductCode    string `bun:"product_code,type:varchar(10),pk"`
	ProductVersion string `bun:"product_version,type:varchar(100),nullzero"`
	ScriptName     string `bun:"script_name,type:varchar(20),nullzero"`
	DbVersion      string `bun:"db_version,type:varchar(20),nullzero"`
}
