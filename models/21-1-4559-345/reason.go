package model

type Reason struct {
	bun.BaseModel           `bun:"table:reason"`
	Id                      string    `bun:"id,type:varchar(5),pk"`
	Reason                  string    `bun:"reason,type:varchar(40)"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	GlAccountNo             string    `bun:"gl_account_no,type:varchar(32),nullzero"`
	CompanyNo               string    `bun:"company_no,type:varchar(8),nullzero"`
	AffectActualUsageFlag   string    `bun:"affect_actual_usage_flag,type:char,default:('N')"`
	DeliveryReasonFlag      string    `bun:"delivery_reason_flag,type:char,default:('N')"`
	UseBranchAccFlag        string    `bun:"use_branch_acc_flag,type:char,default:('Y')"`
	ArCashReasonFlag        string    `bun:"ar_cash_reason_flag,type:char,default:('N')"`
	MacUpdateReasonFlag     string    `bun:"mac_update_reason_flag,type:char,default:('N')"`
	RmaReasonFlag           string    `bun:"rma_reason_flag,type:char,default:('N')"`
	FieldDestroyReasonFlag  string    `bun:"field_destroy_reason_flag,type:char,default:('N')"`
	InvAdjItemLevelFlag     string    `bun:"inv_adj_item_level_flag,type:char,default:('N')"`
	InvAdjItemLevelEditFlag string    `bun:"inv_adj_item_level_edit_flag,type:char,default:('N')"`
	ReturnToStockFlag       string    `bun:"return_to_stock_flag,type:char,default:('N')"`
	ReturnAction            int32     `bun:"return_action,type:int,nullzero"`
	MafSurchargeReasonFlag  string    `bun:"maf_surcharge_reason_flag,type:char,default:('N')"`
}
