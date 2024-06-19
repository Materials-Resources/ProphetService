package gen

import "github.com/uptrace/bun"

type Refactorlog struct {
	bun.BaseModel `bun:"table:__RefactorLog"`
	Operationkey  string `bun:"operationkey,type:uniqueidentifier,pk"` // OperationKey
}
