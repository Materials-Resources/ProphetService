package model

type CustomerEdiSetting struct {
	bun.BaseModel             `bun:"table:customer_edi_setting"`
	CustomerEdiSettingUid     int32     `bun:"customer_edi_setting_uid,type:int,pk,identity"`
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`
	CustomerId                float64   `bun:"customer_id,type:decimal(19,0)"`
	InterchgReceiverId        string    `bun:"interchg_receiver_id,type:varchar(255),nullzero"`
	IntlSan                   string    `bun:"intl_san,type:varchar(255),nullzero"`
	TradingPartnerName        string    `bun:"trading_partner_name,type:varchar(255),nullzero"`
	PassportCustomerId        string    `bun:"passport_customer_id,type:varchar(255),nullzero"`
	EdiInterchangeIdQualifier string    `bun:"edi_interchange_id_qualifier,type:varchar(255),nullzero"`
	EdiInterchangeId          string    `bun:"edi_interchange_id,type:varchar(255),nullzero"`
	ApplicationCode           string    `bun:"application_code,type:varchar(255),nullzero"`
	ElementSeparator          string    `bun:"element_separator,type:varchar(255),default:('|')"`
	SubElementSeparator       string    `bun:"sub_element_separator,type:varchar(255),default:('>')"`
	SegmentTerminator         string    `bun:"segment_terminator,type:varchar(255),default:('~')"`
	RepetitionSeparator       string    `bun:"repetition_separator,type:varchar(255),default:('<')"`
	AppendLineFeedFlag        string    `bun:"append_line_feed_flag,type:char,default:('N')"`
	FunctionalAckFlag         string    `bun:"functional_ack_flag,type:char,default:('Y')"`
	TestingModeFlag           string    `bun:"testing_mode_flag,type:char,default:('Y')"`
	EightyColumnLineBreakFlag string    `bun:"eighty_column_line_break_flag,type:char,default:('N')"`
	ValidateX12DocumentFlag   string    `bun:"validate_x12_document_flag,type:char,default:('Y')"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	Edi855ConsignmentOnlyFlag string    `bun:"edi855_consignment_only_flag,type:char,default:('N')"`
}
