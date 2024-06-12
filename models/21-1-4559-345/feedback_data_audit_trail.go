package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FeedbackDataAuditTrail struct {
	bun.BaseModel             `bun:"table:feedback_data_audit_trail"`
	FeedbackDataAuditTrailUid int32     `bun:"feedback_data_audit_trail_uid,type:int,pk,identity"`
	UsersId                   string    `bun:"users_id,type:varchar(40)"`
	CaseNo                    string    `bun:"case_no,type:varchar(255),nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SupportQueryCd            int32     `bun:"support_query_cd,type:int"`
	RetrievalArg1             string    `bun:"retrieval_arg1,type:varchar(255),nullzero"`
	RetrievalArg2             string    `bun:"retrieval_arg2,type:varchar(255),nullzero"`
	RetrievalArg3             string    `bun:"retrieval_arg3,type:varchar(255),nullzero"`
	RetrievalArg4             string    `bun:"retrieval_arg4,type:varchar(255),nullzero"`
	RetrievalArg5             string    `bun:"retrieval_arg5,type:varchar(255),nullzero"`
	RetrievalArg6             string    `bun:"retrieval_arg6,type:varchar(255),nullzero"`
	RetrievalArg7             string    `bun:"retrieval_arg7,type:varchar(255),nullzero"`
	RetrievalArg8             string    `bun:"retrieval_arg8,type:varchar(255),nullzero"`
	RetrievalArg9             string    `bun:"retrieval_arg9,type:varchar(255),nullzero"`
	RetrievalArg10            string    `bun:"retrieval_arg10,type:varchar(255),nullzero"`
	RetrievalArg11            string    `bun:"retrieval_arg11,type:varchar(255),nullzero"`
	RetrievalArg12            string    `bun:"retrieval_arg12,type:varchar(255),nullzero"`
	RetrievalArg13            string    `bun:"retrieval_arg13,type:varchar(255),nullzero"`
	RetrievalArg14            string    `bun:"retrieval_arg14,type:varchar(255),nullzero"`
	RetrievalArg15            string    `bun:"retrieval_arg15,type:varchar(255),nullzero"`
	ResultsFile               string    `bun:"results_file,type:varchar(8000),nullzero"`
	FeedbackActionCd          int32     `bun:"feedback_action_cd,type:int"`
}
