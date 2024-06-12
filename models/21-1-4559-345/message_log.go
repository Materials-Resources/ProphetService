package model

type MessageLog struct {
	bun.BaseModel    `bun:"table:message_log"`
	MessageLogUid    int32     `bun:"message_log_uid,type:int,pk"`
	MessageNo        int32     `bun:"message_no,type:int"`
	MessageDate      time.Time `bun:"message_date,type:datetime"`
	ApplicationName  string    `bun:"application_name,type:varchar(255)"`
	UserId           string    `bun:"user_id,type:varchar(30),default:(user_name(null))"`
	UserText         string    `bun:"user_text,type:varchar(4000)"`
	TechnicalText    string    `bun:"technical_text,type:varchar(4000)"`
	ButtonPressed    int32     `bun:"button_pressed,type:int"`
	MsgSeverityLevel int32     `bun:"msg_severity_level,type:int"`
	Stack            string    `bun:"stack,type:varchar(255)"`
	UserTextExtended string    `bun:"user_text_extended,type:text(2147483647),nullzero"`
}
