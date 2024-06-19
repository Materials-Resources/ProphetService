package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportSuspenseLine struct {
	bun.BaseModel         `bun:"table:import_suspense_line"`
	ImportSuspenseLineUid int32     `bun:"import_suspense_line_uid,type:int,pk"`                      // Unique Identifier of table
	ImportSuspenseHdrUid  int32     `bun:"import_suspense_hdr_uid,type:int"`                          // Unique Identifier of import_suspense_hdr table
	SequenceNo            int32     `bun:"sequence_no,type:int"`                                      // Number indicating the sequence in which each file is imported - used within the application
	ImportFilePath        string    `bun:"import_file_path,type:varchar(255)"`                        // The path of the file containing the import record being suspended, for scheduled imports, this will be the polling path
	ImportFileName        string    `bun:"import_file_name,type:varchar(255)"`                        // Name of the file containing the import record being suspended, without the path
	BoInstancename        string    `bun:"bo_instancename,type:varchar(255),nullzero"`                // Internal use - identifies the business objects associated with the file- often the same as the table being updated
	DsDataobject          string    `bun:"ds_dataobject,type:varchar(255)"`                           // Internal use - identifies the data container definition for the import layout
	DwDataobject          string    `bun:"dw_dataobject,type:varchar(255)"`                           // Internal use - identifies the data container definition for the displayed import files
	ImportSetNo           string    `bun:"import_set_no,type:varchar(255)"`                           // The Import Set No associated with the suspended record
	CompanyId             string    `bun:"company_id,type:varchar(8),nullzero"`                       // Unique code that identifies a company.
	CustomerId            float64   `bun:"customer_id,type:decimal(19,0),nullzero"`                   // used for sales order imports, the customer id of the order header record
	LocationId            float64   `bun:"location_id,type:decimal(19,0),nullzero"`                   // used for sales order imports, the sales location id of the order header record
	KeyValue              string    `bun:"key_value,type:varchar(255),nullzero"`                      // Used for sales order imports, identifies the order number of an imported order for which a line item failed
	KeyValueTable         string    `bun:"key_value_table,type:varchar(255),nullzero"`                // Used for sales order imports, identifies the table associated with the key_value
	ExportedStatus        int32     `bun:"exported_status,type:int,default:(0)"`                      // Used during the editing of suspended sales order records, identifies whether an order acknowledgement has been sent; reset to No when data is edited.
	RowStatusFlag         int32     `bun:"row_status_flag,type:int,default:(1183)"`                   // Indicates current record status.
	DisplayName           string    `bun:"display_name,type:varchar(255),default:('Unknown')"`        // Text used for the tab which displays the suspense record
	SuspenseData          string    `bun:"suspense_data,type:text"`                                   // the tab-delimited import record
	ErrorData             string    `bun:"error_data,type:text"`                                      // the tab-delimited error record(s) associated with the suspense record
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	SourceId              string    `bun:"source_id,type:varchar(8),nullzero"`                        // web reference number of B2BSeller Sales Order imports
}
