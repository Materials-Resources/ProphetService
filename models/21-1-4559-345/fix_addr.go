package prophet

import "github.com/uptrace/bun"

type FixAddr struct {
	bun.BaseModel `bun:"table:fix_addr"`
	Col001        *string `bun:"Col001,type:varchar(255)"`
	Col002        *string `bun:"Col002,type:varchar(255)"`
	Col003        *string `bun:"Col003,type:varchar(255)"`
	Col004        *string `bun:"Col004,type:varchar(255)"`
	Col005        *string `bun:"Col005,type:varchar(255)"`
	Col006        *string `bun:"Col006,type:varchar(255)"`
	Col007        *string `bun:"Col007,type:varchar(255)"`
	Col008        *string `bun:"Col008,type:varchar(255)"`
	Col009        *string `bun:"Col009,type:varchar(255)"`
	Col010        *string `bun:"Col010,type:varchar(255)"`
	Col011        *string `bun:"Col011,type:varchar(255)"`
	Col012        *string `bun:"Col012,type:varchar(255)"`
	Col013        *string `bun:"Col013,type:varchar(255)"`
}
