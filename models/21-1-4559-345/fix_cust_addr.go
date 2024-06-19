package gen

import "github.com/uptrace/bun"

type FixCustAddr struct {
	bun.BaseModel `bun:"table:fix_cust_addr"`
	Col001        string `bun:"Col001,type:varchar(255),nullzero"`
	Col002        string `bun:"Col002,type:varchar(255),nullzero"`
	Col003        string `bun:"Col003,type:varchar(255),nullzero"`
	Col004        string `bun:"Col004,type:varchar(255),nullzero"`
	Col005        string `bun:"Col005,type:varchar(255),nullzero"`
	Col006        string `bun:"Col006,type:varchar(255),nullzero"`
	Col007        string `bun:"Col007,type:varchar(255),nullzero"`
	Col008        string `bun:"Col008,type:varchar(255),nullzero"`
	Col009        string `bun:"Col009,type:varchar(255),nullzero"`
}
