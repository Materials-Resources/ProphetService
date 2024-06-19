package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Messages struct {
	bun.BaseModel    `bun:"table:messages"`
	MessageNo        int32     `bun:"message_no,type:int,pk"`                                        // Message number that is displayed.
	LanguageId       string    `bun:"language_id,type:varchar(8)"`                                   // What language is this message in?
	MessageTitle     string    `bun:"message_title,type:varchar(50)"`                                // Title of the message
	UserText         string    `bun:"user_text,type:varchar(255)"`                                   // The text displayed to the user.
	TechnicalText    string    `bun:"technical_text,type:varchar(255)"`                              // The technical text written to the message log table.
	Icon             string    `bun:"icon,type:varchar(20)"`                                         // Icon that is associated with the message.
	Button           string    `bun:"button,type:varchar(30)"`                                       // What button should appear in the dialog box contai
	DefaultButton    float64   `bun:"default_button,type:decimal(2,0)"`                              // The default button if there are multiple. Ex. OK/Cancel
	MsgSeverityLevel int32     `bun:"msg_severity_level,type:int"`                                   // Level of severity of the message.
	PrintIndicator   string    `bun:"print_indicator,type:char(1)"`                                  // Indicates if this message can be printed.
	UserInput        string    `bun:"user_input,type:char(1)"`                                       // Indicates that the user can add notes to the message before copying/printing.
	WindowFlag       string    `bun:"window_flag,type:char(1)"`                                      // Used internally to determine whether to render the message as a message box or response window.
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	UseLongMessage   string    `bun:"use_long_message,type:char(1)"`                                 // Should this message be written to the user_text or
	CopyIndicator    string    `bun:"copy_indicator,type:char(1)"`                                   // This column will be used to signify whether or not a copy button should be displayed on an error message window.  It will function similiarily to the other buttons.
}
