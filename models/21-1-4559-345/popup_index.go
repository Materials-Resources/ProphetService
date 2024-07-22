package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PopupIndex struct {
	bun.BaseModel    `bun:"table:popup_index"`
	PopupIndexUid    int32     `bun:"popup_index_uid,type:int,autoincrement,identity,pk"`           // Popup Index Id
	Dwcontrol        *string   `bun:"dwcontrol,type:varchar(50),unique"`                            // DW control associated to popup
	Window           *string   `bun:"window,type:varchar(50),unique"`                               // Window associated to popup
	Dwfield          *string   `bun:"dwfield,type:varchar(50),unique"`                              // Field associated to popup
	Role             *string   `bun:"role,type:varchar(50),unique"`                                 // Role for the popup
	Condition        *string   `bun:"condition,type:varchar(255),unique"`                           // Additional condition to popup
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	PopupDetailUid   int32     `bun:"popup_detail_uid,type:int"`                                    // popup_detail.popup_detail_uid foreign key
	UserRole         *string   `bun:"user_role,type:varchar(50),unique"`                            // user role
	Dataobject       *string   `bun:"dataobject,type:varchar(255),unique"`                          // Datawindow object
	ConfigurationId  int32     `bun:"configuration_id,type:int,default:((0))"`                      // Used when this popup is used for specific configuration.
	Action           *string   `bun:"action,type:varchar(50),unique"`                               // Defines if the popup should be opened only when a certain action happens.
	UserId           *string   `bun:"user_id,type:varchar(30),unique"`                              // Defines if the popup should be opened only by an user.
	Shortcut         *string   `bun:"shortcut,type:varchar(30),unique"`                             // Defines if the popup should be opened only when using a certain shortcut.
	CoreFlag         bool      `bun:"core_flag,type:bit,default:((0)),unique"`                      // Save if  the index is core
	DynachangeId     *float64  `bun:"dynachange_id,type:decimal(6,0),unique"`                       // Determines the identifier of the dynachange associated to this popup
}
