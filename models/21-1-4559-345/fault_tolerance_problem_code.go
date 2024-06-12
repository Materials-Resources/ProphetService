package model

type FaultToleranceProblemCode struct {
	bun.BaseModel `bun:"table:fault_tolerance_problem_code"`
	FtpUid        int32  `bun:"ftp_uid,type:int,pk"`
	FtpCode       string `bun:"ftp_code,type:varchar(255)"`
}
