package model

type HelpAccess struct {
	bun.BaseModel       `bun:"table:help_access"`
	HelpLevel           string    `bun:"help_level,type:varchar(25)"`
	WindowName          string    `bun:"window_name,type:varchar(100)"`
	ObjectName          string    `bun:"object_name,type:varchar(100)"`
	KeywordSource       string    `bun:"keyword_source,type:varchar(50)"`
	Keyword             string    `bun:"keyword,type:varchar(255)"`
	HelpTopicContext    string    `bun:"help_topic_context,type:varchar(50),nullzero"`
	HelpTopicColumnName string    `bun:"help_topic_column_name,type:varchar(50),nullzero"`
	HelpFileName        string    `bun:"help_file_name,type:varchar(255),nullzero"`
	HelpFileDate        time.Time `bun:"help_file_date,type:datetime,nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
	HelpAccessUid       int32     `bun:"help_access_uid,type:int,pk,identity"`
}
