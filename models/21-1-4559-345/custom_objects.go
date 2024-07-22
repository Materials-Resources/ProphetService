package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomObjects struct {
	bun.BaseModel    `bun:"table:custom_objects"`
	CustomObjectsUid int32     `bun:"custom_objects_uid,type:int,pk"`                            // The unique row identifier for the table
	UsersId          string    `bun:"users_id,type:varchar(30)"`                                 // The user_id the version is assigned to
	Object           string    `bun:"object,type:varchar(128)"`                                  // The name of the object.
	ConfigurationId  int32     `bun:"configuration_id,type:int"`                                 // The customer configuration Id.
	ModString        string    `bun:"mod_string,type:text"`                                      // The object modification syntax
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	VersionId        string    `bun:"version_id,type:varchar(128)"`                              // The version ID of the DynaChange Personalizations/Customizations.
	VersionDesc      string    `bun:"version_desc,type:varchar(255)"`                            // Description of the version
	RoleId           int32     `bun:"role_id,type:int"`                                          // The role id the version is assigned to.
	Type             string    `bun:"type,type:char(1)"`                                         // Column to identify whether the version belongs to a user or a role.
	ObjectType       string    `bun:"object_type,type:char(1)"`                                  // The type of object( Data window,Menu,Tab)
	ApplyToAll       int32     `bun:"apply_to_all,type:int,default:(2)"`                         // Apply the version_id modifications to all objects
	DefaultValues    *string   `bun:"default_values,type:text"`                                  // default values for tab controls in dynachange designer and version manager
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`                  // Status of the version
	DesignUid        *int32    `bun:"design_uid,type:int"`                                       // Design UID
	MigratedToWeb    string    `bun:"migrated_to_web,type:char(1),default:('N')"`                // stores a flag indicating migration status
}
