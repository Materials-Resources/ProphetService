package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLinkWindowXKey struct {
	bun.BaseModel             `bun:"table:document_link_window_x_key"`
	DocumentLinkWindowXKeyUid int32     `bun:"document_link_window_x_key_uid,type:int,pk"`
	DocumentLinkWindowUid     int32     `bun:"document_link_window_uid,type:int"`
	DocumentLinkKeyUid        int32     `bun:"document_link_key_uid,type:int"`
	TabClass                  string    `bun:"tab_class,type:varchar(255)"`
	DwName                    string    `bun:"dw_name,type:varchar(255)"`
	Key1Column                *string   `bun:"key1_column,type:varchar(255)"`
	Key1Type                  *string   `bun:"key1_type,type:varchar(5)"`
	Key1Source                *int16    `bun:"key1_source,type:tinyint,default:(0)"`
	Key2Column                *string   `bun:"key2_column,type:varchar(255)"`
	Key2Type                  *string   `bun:"key2_type,type:varchar(5)"`
	Key2Source                *int16    `bun:"key2_source,type:tinyint,default:(0)"`
	Key3Column                *string   `bun:"key3_column,type:varchar(255)"`
	Key3Type                  *string   `bun:"key3_type,type:varchar(5)"`
	Key3Source                *int16    `bun:"key3_source,type:tinyint,default:(0)"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DisplayOutsideUseFlag     *string   `bun:"display_outside_use_flag,type:char(1),default:('N')"`
	DisplayMandatoryFlag      *string   `bun:"display_mandatory_flag,type:char(1),default:('N')"`
	DisplayPrintFlag          *string   `bun:"display_print_flag,type:char(1),default:('N')"`
	DisplayFaxFlag            *string   `bun:"display_fax_flag,type:char(1),default:('N')"`
	DisplayEmailFlag          *string   `bun:"display_email_flag,type:char(1),default:('N')"`
	DisplayAllSourceLinksFlag *string   `bun:"display_all_source_links_flag,type:char(1),default:('N')"`
	TransmitLotLinkFlag       *string   `bun:"transmit_lot_link_flag,type:char(1),default:('N')"` // Indicates whether to automatically transmit outside use lot document links to ship to
	PrintItemLinkFlag         *string   `bun:"print_item_link_flag,type:char(1),default:('N')"`   // Indicates whether to automatically print mandatory item document links
	IncludeForImportRole      *string   `bun:"include_for_import_role,type:varchar(255)"`         // Import role for which document_link functionality should be included
}
