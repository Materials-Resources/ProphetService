package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type LabelDefinitionXLoc struct {
	bun.BaseModel            `bun:"table:label_definition_x_loc"`
	LabelDefinitionXLocUid   int32     `bun:"label_definition_x_loc_uid,type:int,pk"`                       // Unique ID for this record
	LabelDefinitionUid       int32     `bun:"label_definition_uid,type:int"`                                // UID of the label definition record
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                                   // Company ID for this location
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`                               // Location ID for this record
	LabelProcessTypeCd       int32     `bun:"label_process_type_cd,type:int"`                               // Label Printing Processes code
	LabelTemplateFilename    string    `bun:"label_template_filename,type:varchar(255),nullzero"`           // Path to label template file (if null, just the label definition ID will be passed to the forms package, leaving the onus for finding the actual file path on the forms processing software)
	DefaultPrintOptionCd     int32     `bun:"default_print_option_cd,type:int"`                             // Default option code for printing this label (Yes/No/Prompt)
	DefaultPrinter           string    `bun:"default_printer,type:varchar(255),nullzero"`                   // Default printer for this label at the location
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                     // Status of this row
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DefaultForProcXLabelType string    `bun:"default_for_proc_x_label_type,type:char(1),default:('N')"`     // The label definition defaulted in WWMS label printing
}
