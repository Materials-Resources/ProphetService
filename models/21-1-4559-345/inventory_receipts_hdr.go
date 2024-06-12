package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReceiptsHdr struct {
	bun.BaseModel           `bun:"table:inventory_receipts_hdr"`
	ReceiptNumber           float64   `bun:"receipt_number,type:decimal(19,0),pk"`
	PoNumber                float64   `bun:"po_number,type:decimal(19,0)"`
	Approved                string    `bun:"approved,type:char"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	VouchComplete           string    `bun:"vouch_complete,type:char"`
	Period                  float64   `bun:"period,type:decimal(3,0)"`
	YearForPeriod           float64   `bun:"year_for_period,type:decimal(4,0)"`
	FreightAmount           float64   `bun:"freight_amount,type:decimal(19,4),nullzero"`
	FreightAmountDisplay    float64   `bun:"freight_amount_display,type:decimal(19,4),nullzero"`
	CurrencyId              float64   `bun:"currency_id,type:decimal(19,0)"`
	ReceiptType             string    `bun:"receipt_type,type:char"`
	PackingSlipNumber       string    `bun:"packing_slip_number,type:varchar(255),nullzero"`
	FreightCodeUid          int32     `bun:"freight_code_uid,type:int,nullzero"`
	ReceiverNumber          string    `bun:"receiver_number,type:varchar(255),nullzero"`
	ExternalReferenceNo     string    `bun:"external_reference_no,type:varchar(255),nullzero"`
	PoAsnHdrUid             int32     `bun:"po_asn_hdr_uid,type:int,nullzero"`
	ContainerReceiptsHdrUid int32     `bun:"container_receipts_hdr_uid,type:int,nullzero"`
	ShipmentDate            time.Time `bun:"shipment_date,type:datetime,nullzero"`
	SourceTypeCd            int32     `bun:"source_type_cd,type:int,nullzero"`
	InvReceiptsClearingAcct string    `bun:"inv_receipts_clearing_acct,type:varchar(32),nullzero"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	GroupPoHdrUid           int32     `bun:"group_po_hdr_uid,type:int,nullzero"`
	RfnavTransNo            int32     `bun:"rfnav_trans_no,type:int,nullzero"`
}
