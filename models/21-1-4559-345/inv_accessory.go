package model

type InvAccessory struct {
	bun.BaseModel         `bun:"table:inv_accessory"`
	InvAccessoryUid       int32     `bun:"inv_accessory_uid,type:int,pk"`
	ParentInvMastUid      int32     `bun:"parent_inv_mast_uid,type:int"`
	ChildInvMastUid       int32     `bun:"child_inv_mast_uid,type:int"`
	AutoPopulateQuantity  string    `bun:"auto_populate_quantity,type:char,default:('N')"`
	ScaleToParentQuantity string    `bun:"scale_to_parent_quantity,type:char,default:('Y')"`
	ChildUnitQuantity     float64   `bun:"child_unit_quantity,type:decimal(19,9),default:(0)"`
	ChildUomCodeNo        int32     `bun:"child_uom_code_no,type:int"`
	ParentUnitQuantity    float64   `bun:"parent_unit_quantity,type:decimal(19,9),default:(0)"`
	ParentUomCodeNo       int32     `bun:"parent_uom_code_no,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LinkQtyToParentFlag   string    `bun:"link_qty_to_parent_flag,type:char,default:('N')"`
	ShipTogetherFlag      string    `bun:"ship_together_flag,type:char,default:('N')"`
	MandatoryFlag         string    `bun:"mandatory_flag,type:char,default:('N')"`
}
