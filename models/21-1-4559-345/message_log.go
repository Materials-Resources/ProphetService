package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type MessageLog struct {
	bun.BaseModel    `bun:"table:message_log"`
	MessageLogUid    int32     `bun:"message_log_uid,type:int,pk"`                        // Unique Identifier of record
	MessageNo        int32     `bun:"message_no,type:int"`                                // Number for the message that will be displayed.
	MessageDate      time.Time `bun:"message_date,type:datetime"`                         // This column is unused.
	ApplicationName  string    `bun:"application_name,type:varchar(255)"`                 // This column is unused.
	UserId           string    `bun:"user_id,type:varchar(30),default:(user_name(null))"` // This column is unused.
	UserText         string    `bun:"user_text,type:varchar(4000)"`                       // This column is unused.
	TechnicalText    string    `bun:"technical_text,type:varchar(4000)"`                  // This column is unused.
	ButtonPressed    int32     `bun:"button_pressed,type:int"`                            // This column is unused.
	MsgSeverityLevel int32     `bun:"msg_severity_level,type:int"`                        // This column is unused.
	Stack            string    `bun:"stack,type:varchar(255)"`                            // This column is unused.
	UserTextExtended *string   `bun:"user_text_extended,type:text"`                       // This column is unused.
}
