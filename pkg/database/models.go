package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type Address struct {
	bun.BaseModel `bun:"table:address"`

	Id                  float64         `bun:"id,pk"`
	Name                string          `bun:"name"`
	MailAddress1        sql.NullString  `bun:"mail_address1"`
	MailAddress2        sql.NullString  `bun:"mail_address2"`
	MailCity            sql.NullString  `bun:"mail_city"`
	MailState           sql.NullString  `bun:"mail_state"`
	MailPostalCode      sql.NullString  `bun:"mail_postal_code"`
	MailCountry         sql.NullString  `bun:"mail_country"`
	PhysAddress1        sql.NullString  `bun:"phys_address1"`
	PhysAddress2        sql.NullString  `bun:"phys_address2"`
	PhysCity            sql.NullString  `bun:"phys_city"`
	PhysState           sql.NullString  `bun:"phys_state"`
	PhysPostalCode      sql.NullString  `bun:"phys_postal_code"`
	PhysCountry         sql.NullString  `bun:"phys_country"`
	CentralPhoneNo      sql.NullString  `bun:"central_phone_number"`
	CentralFaxNo        sql.NullString  `bun:"central_fax_number"`
	CorpAddressId       sql.NullFloat64 `bun:"corp_address_id"`
	ShippingAddressFlag sql.NullString  `bun:"shipping_address"`
}

type Contact struct {
	bun.BaseModel `bun:"table:contacts"`
	ID            string         `bun:"id,pk"`
	FirstName     string         `bun:"first_name"`
	MiddleInitial sql.NullString `bun:"mi"`
	LastName      string         `bun:"last_name"`
	Title         sql.NullString `bun:"title"`
	EmailAddress  sql.NullString `bun:"email_address"`
	DirectPhone   sql.NullString `bun:"direct_phone"`
}

type InventoryReceipt struct {
	bun.BaseModel `bun:"table:inventory_receipts_hdr"`
	Id            int32                  `bun:"receipt_number,pk"`
	Items         []InventoryReceiptLine `bun:"rel:has-many,join:receipt_number=receipt_number"`
}
type InventoryReceiptLine struct {
	bun.BaseModel `bun:"table:inventory_receipts_line"`

	ReceiptNumber     int32                   `bun:"receipt_number,pk"`
	LineNumber        int32                   `bun:"line_number,pk"`
	InvMastUid        int32                   `bun:"inv_mast_uid"`
	Qty               float32                 `bun:"unit_quantity"`
	Uom               string                  `bun:"unit_of_measure"`
	InventoryReceipt  InventoryReceipt        `bun:"rel:belongs-to,join:receipt_number=receipt_number"`
	OrderTransactions []*InventoryTransaction `bun:"rel:has-many,join:receipt_number=sub_document_no,join:inv_mast_uid=inv_mast_uid,join:type=trans_type,polymorphic:IRA"`
	InvMast           ProductMaster           `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

type InventoryTransaction struct {
	bun.BaseModel `bun:"table:inv_tran"`

	TransactionNumber float64 `bun:"transaction_number,pk"`
	SubDocumentNo     int32   `bun:"sub_document_no"`
	InvMastUid        int32   `bun:"inv_mast_uid"`
	QtyAllocated      float32 `bun:"qty_allocated"`
	DocumentNo        int32   `bun:"document_no"`
	TransType         string  `bun:"trans_type"`
}

type ProductMaster struct {
	bun.BaseModel `bun:"table:inv_mast"`

	InvMastUid          int32   `bun:"inv_mast_uid,pk"`
	ItemId              string  `bun:"item_id"`
	ItemDesc            string  `bun:"item_desc"`
	Price1              float64 `bun:"price1"`
	Price2              float64 `bun:"price2"`
	Price3              float64 `bun:"price3"`
	DefaultProductGroup string  `bun:"default_product_group"`
	ExtendedDesc        string  `bun:"extended_desc"`

	InvLoc            *InvLoc            `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
	InventorySupplier *InventorySupplier `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

type ProductGroup struct {
	bun.BaseModel             `bun:"table:product_group"`
	CompanyId                 string    `bun:"company_id,pk,notnull"`
	ProductGroupId            string    `bun:"product_group_id,pk,notnull"`
	ProductGroupDesc          string    `bun:"product_group_desc,notnull"`
	AssetAccountNo            string    `bun:"asset_account_no,notnull"`
	RevenueAccountNo          string    `bun:"revenue_account_no"`
	CosAccountNo              string    `bun:"cos_account_no"`
	DeleteFlag                string    `bun:"delete_flag,notnull"`
	DateCreated               time.Time `bun:"date_created,notnull"`
	DateLastModified          time.Time `bun:"date_last_modified,notnull"`
	LastMaintainedBy          string    `bun:"last_maintained_by,notnull"`
	ProductGroupUid           int32     `bun:"product_group_uid,unique"`
	EnvironmentalFeeFlag      string    `bun:"environmental_fee_flag,notnull,default:N"`
	AdminFleeFlag             string    `bun:"admin_fee_flag,notnull,default:N"`
	IncludeInSizeAnalysisFlag string    `bun:"include_in_size_analysis_flag,default:Y"`
	PriceRoundingFlag         string    `bun:"price_rounding_flag,notnull,default:N"`
	ApplyMinMarginRulesFlag   string    `bun:"apply_min_margin_rules_flag,notnull,default:N"`
	EnableLineProfitWarning   string    `bun:"enable_line_profit_warning,notnull,default:N"`
	OeSkipProfitCheckUnpriced string    `bun:"oe_skip_profit_check_unpriced,notnull,default:N"`
	OeSkipProfitCheckUncosted string    `bun:"oe_skip_profit_check_uncosted,notnull,default:N"`
	CompressorFlag            string    `bun:"compressor_flag,notnull,default:N"`
}

type PickTicket struct {
	bun.BaseModel `bun:"table:oe_pick_ticket"`

	PickTicketNo float64 `bun:"pick_ticket_no,pk"`
	OrderNo      string  `bun:"order_no"`
}
