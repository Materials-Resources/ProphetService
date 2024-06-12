package model

type OePickTicketDetail struct {
	bun.BaseModel               `bun:"table:oe_pick_ticket_detail"`
	CompanyId                   string    `bun:"company_id,type:varchar(8)"`
	PickTicketNo                float64   `bun:"pick_ticket_no,type:decimal(19,0),pk"`
	LineNumber                  float64   `bun:"line_number,type:decimal(19,0),pk"`
	PrintQuantity               float64   `bun:"print_quantity,type:decimal(19,9),nullzero"`
	ShipQuantity                float64   `bun:"ship_quantity,type:decimal(19,9),nullzero"`
	FreightIn                   float64   `bun:"freight_in,type:decimal(19,4),nullzero"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	UnitOfMeasure               string    `bun:"unit_of_measure,type:varchar(8)"`
	UnitSize                    float64   `bun:"unit_size,type:decimal(19,4)"`
	UnitQuantity                float64   `bun:"unit_quantity,type:decimal(19,9)"`
	OeLineNo                    float64   `bun:"oe_line_no,type:decimal(19,0)"`
	QtyRequested                float64   `bun:"qty_requested,type:decimal(19,9),nullzero"`
	QtyToPick                   float64   `bun:"qty_to_pick,type:decimal(19,9),nullzero"`
	Staged                      string    `bun:"staged,type:char,nullzero"`
	ReleaseNo                   float64   `bun:"release_no,type:decimal(6,0),nullzero"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int"`
	InvoiceLineUid              int32     `bun:"invoice_line_uid,type:int,nullzero"`
	FreightCodeUid              int32     `bun:"freight_code_uid,type:int,nullzero"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	UnitQuantityEdited          string    `bun:"unit_quantity_edited,type:char,default:('N')"`
	EnvironmentalFeeFlag        string    `bun:"environmental_fee_flag,type:char,nullzero"`
	AdminFeeFlag                string    `bun:"admin_fee_flag,type:char,nullzero"`
	InspectedFlag               string    `bun:"inspected_flag,type:char,default:('N')"`
	OriginalFreightIn           float64   `bun:"original_freight_in,type:decimal(19,9),nullzero"`
	QtyScanned                  float64   `bun:"qty_scanned,type:decimal(19,9),nullzero"`
	FreightInMinAppliedFlag     string    `bun:"freight_in_min_applied_flag,type:char,nullzero"`
	BoxNumber                   string    `bun:"box_number,type:varchar(255),nullzero"`
	OriginalQtyToPick           float64   `bun:"original_qty_to_pick,type:decimal(19,9),nullzero"`
	CountryOfOrigin             string    `bun:"country_of_origin,type:varchar(8),nullzero"`
	SubPickTicketNo             int32     `bun:"sub_pick_ticket_no,type:int,nullzero"`
	ShippingCountryCodeUid      int32     `bun:"shipping_country_code_uid,type:int,nullzero"`
	BoxLpn                      string    `bun:"box_lpn,type:varchar(255),nullzero"`
	PrintPartNo                 string    `bun:"print_part_no,type:varchar(40),nullzero"`
	NewJobPriceShipControlNoUid int32     `bun:"new_job_price_ship_control_no_uid,type:int,nullzero"`
	ReturnedCylinderQuantity    int32     `bun:"returned_cylinder_quantity,type:int,nullzero"`
	RfnavPickStatusCd           int32     `bun:"rfnav_pick_status_cd,type:int,nullzero"`
}
