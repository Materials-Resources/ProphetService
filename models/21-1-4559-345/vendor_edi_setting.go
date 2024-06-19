package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorEdiSetting struct {
	bun.BaseModel             `bun:"table:vendor_edi_setting"`
	VendorEdiSettingUid       int32     `bun:"vendor_edi_setting_uid,type:int,autoincrement,scanonly,pk"`    // Unique identifier for each record
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`                                   // Company ID for Vendor to which settings apply
	VendorId                  float64   `bun:"vendor_id,type:decimal(19,0)"`                                 // Vendor to which settings apply
	InterchgReceiverId        string    `bun:"interchg_receiver_id,type:varchar(255),nullzero"`              // Vendor standard address name for EDI transactions
	IntlSan                   string    `bun:"intl_san,type:varchar(255),nullzero"`                          // International standard address name
	TradingPartnerName        string    `bun:"trading_partner_name,type:varchar(255),nullzero"`              // Name used for TPCx transactions
	EdiInterchangeIdQualifier string    `bun:"edi_interchange_id_qualifier,type:varchar(255),nullzero"`      // Qualifier for EDI interchange ID
	EdiInterchangeId          string    `bun:"edi_interchange_id,type:varchar(255),nullzero"`                // EDI Interchange ID
	ApplicationCode           string    `bun:"application_code,type:varchar(255),nullzero"`                  // EDI application code
	ElementSeparator          string    `bun:"element_separator,type:varchar(255),default:('|')"`            // Up to two characters used as element separator
	SubElementSeparator       string    `bun:"sub_element_separator,type:varchar(255),default:('>')"`        // Up to two characters used as sub-element separator
	SegmentTerminator         string    `bun:"segment_terminator,type:varchar(255),default:('~')"`           // Up to two characters used as segment terminator
	RepetitionSeparator       string    `bun:"repetition_separator,type:varchar(255),default:('<')"`         // Up to two characters used as repetition terminator
	AppendLineFeedFlag        string    `bun:"append_line_feed_flag,type:char(1),default:('N')"`             // Indicates whether to append line feed
	FunctionalAckFlag         string    `bun:"functional_ack_flag,type:char(1),default:('Y')"`               // Indicates whether to request functional acknowledgement
	TestingModeFlag           string    `bun:"testing_mode_flag,type:char(1),default:('Y')"`                 // Indicates whether in testing mode
	EightyColumnLineBreakFlag string    `bun:"eighty_column_line_break_flag,type:char(1),default:('N')"`     // Indicates whether to use 80 column line break
	ValidateX12DocumentFlag   string    `bun:"validate_x12_document_flag,type:char(1),default:('Y')"`        // Indicates whether to validate X12 document
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	AdditionalFreight         float64   `bun:"additional_freight,type:decimal(19,9),default:((0))"`          // Custom Feature 57117: contains the amount that will be added as an extra charge to the freight
}
