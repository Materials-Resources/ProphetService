package model

type Refactorlog struct {
	bun.BaseModel `bun:"table:__RefactorLog"`
	Operationkey  `bun:"operationkey,type:uniqueidentifier,pk"`
}
