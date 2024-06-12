package model

type Messages struct {
	bun.BaseModel    `bun:"table:messages"`
	MessageNo        int32     `bun:"message_no,type:int,pk"`
	LanguageId       string    `bun:"language_id,type:varchar(8)"`
	MessageTitle     string    `bun:"message_title,type:varchar(50)"`
	UserText         string    `bun:"user_text,type:varchar(255)"`
	TechnicalText    string    `bun:"technical_text,type:varchar(255)"`
	Icon             string    `bun:"icon,type:varchar(20)"`
	Button           string    `bun:"button,type:varchar(30)"`
	DefaultButton    float64   `bun:"default_button,type:decimal(2,0)"`
	MsgSeverityLevel int32     `bun:"msg_severity_level,type:int"`
	PrintIndicator   string    `bun:"print_indicator,type:char"`
	UserInput        string    `bun:"user_input,type:char"`
	WindowFlag       string    `bun:"window_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	UseLongMessage   string    `bun:"use_long_message,type:char"`
	CopyIndicator    string    `bun:"copy_indicator,type:char"`
}
