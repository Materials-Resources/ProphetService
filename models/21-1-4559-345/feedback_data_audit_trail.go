package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FeedbackDataAuditTrail struct {
	bun.BaseModel             `bun:"table:feedback_data_audit_trail"`
	FeedbackDataAuditTrailUid int32     `bun:"feedback_data_audit_trail_uid,type:int,autoincrement,identity,pk"` // Unique identifier for rows of table.
	UsersId                   string    `bun:"users_id,type:varchar(40),unique"`                                 // ID for user running the window.
	CaseNo                    *string   `bun:"case_no,type:varchar(255),unique"`                                 // Support case entered when running window.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
	SupportQueryCd            int32     `bun:"support_query_cd,type:int,unique"`                                 // Code for the transaction type the window was run for.
	RetrievalArg1             *string   `bun:"retrieval_arg1,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg2             *string   `bun:"retrieval_arg2,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg3             *string   `bun:"retrieval_arg3,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg4             *string   `bun:"retrieval_arg4,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg5             *string   `bun:"retrieval_arg5,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg6             *string   `bun:"retrieval_arg6,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg7             *string   `bun:"retrieval_arg7,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg8             *string   `bun:"retrieval_arg8,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg9             *string   `bun:"retrieval_arg9,type:varchar(255)"`                                 // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg10            *string   `bun:"retrieval_arg10,type:varchar(255)"`                                // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg11            *string   `bun:"retrieval_arg11,type:varchar(255)"`                                // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg12            *string   `bun:"retrieval_arg12,type:varchar(255)"`                                // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg13            *string   `bun:"retrieval_arg13,type:varchar(255)"`                                // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg14            *string   `bun:"retrieval_arg14,type:varchar(255)"`                                // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	RetrievalArg15            *string   `bun:"retrieval_arg15,type:varchar(255)"`                                // Runtime argument (type determined at run time by support_query_cd) to determine what records to retrieve.
	ResultsFile               *string   `bun:"results_file,type:varchar(8000)"`                                  // File location of results.
	FeedbackActionCd          int32     `bun:"feedback_action_cd,type:int"`                                      // Feedback action.
}
