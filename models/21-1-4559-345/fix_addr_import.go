package gen

import "github.com/uptrace/bun"

type FixAddrImport struct {
	bun.BaseModel `bun:"table:fix_addr_import"`
	Col001        string `bun:"Col001,type:varchar(255),nullzero"`
	Col002        string `bun:"Col002,type:varchar(255),nullzero"`
	Col003        string `bun:"Col003,type:varchar(255),nullzero"`
	Col004        string `bun:"Col004,type:varchar(255),nullzero"`
	Col005        string `bun:"Col005,type:varchar(255),nullzero"`
	Col006        string `bun:"Col006,type:varchar(255),nullzero"`
	Col007        string `bun:"Col007,type:varchar(255),nullzero"`
	Col008        string `bun:"Col008,type:varchar(255),nullzero"`
	Col009        string `bun:"Col009,type:varchar(255),nullzero"`
	Col010        string `bun:"Col010,type:varchar(255),nullzero"`
	Col011        string `bun:"Col011,type:varchar(255),nullzero"`
	Col012        string `bun:"Col012,type:varchar(255),nullzero"`
	Col013        string `bun:"Col013,type:varchar(255),nullzero"`
}
