package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RfKeysDefault struct {
	bun.BaseModel    `bun:"table:rf_keys_default"`
	KeyhelpUid       int32     `bun:"keyhelp_uid,type:int,autoincrement,scanonly,pk"`           // The primary key for the table
	KeyhelpCode      int32     `bun:"keyhelp_code,type:int"`                                    // This column  identifies the help key. The keyhelp values are defined in n_cst_constant_rf. They define a particular category of action like deposit, drill, skip.
	WwmsWindowCode   int32     `bun:"wwms_window_code,type:int"`                                // This column determines the window where the key is to be displayed.
	WwmsScreenCode   int32     `bun:"wwms_screen_code,type:int"`                                // This column determines the screen where the key is to be displayed.
	Keystroke        string    `bun:"keystroke,type:varchar(255)"`                              // This column has the key combination to press for the function key.
	Description      string    `bun:"description,type:varchar(255)"`                            // This column has the description of action performed on key press.
	EventName        string    `bun:"event_name,type:varchar(255),nullzero"`                    // This column has the name of the event which will be  triggered on key press
	KeystrokeDesc    string    `bun:"keystroke_desc,type:varchar(255),nullzero"`                // This column has the key apart from modifier(shift, ctrl).
	KeystrokeMod     string    `bun:"keystroke_mod,type:varchar(255),nullzero"`                 // This column has the  modifier like shift, ctrl to be used with key.
	ImageFile        string    `bun:"image_file,type:varchar(255)"`                             // This column has the name of the image file to be shown for the function or event
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`           // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:('P21_DBA')"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`     // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:('P21_DBA')"` // User who last changed the record
}
