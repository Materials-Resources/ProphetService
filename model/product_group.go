package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/materials-resources/s_prophet/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type ProductGroup struct {
	bun.BaseModel             `bun:"table:product_group"`
	CompanyId                 string         `bun:"company_id,type:varchar(8),pk"`
	ProductGroupId            string         `bun:"product_group_id,type:varchar(8),pk"`
	ProductGroupDesc          string         `bun:"product_group_desc,type:varchar(40)"`
	AssetAccountNo            string         `bun:"asset_account_no,type:varchar(80)"`
	RevenueAccountNo          sql.NullString `bun:"revenue_account_no,type:varchar(80)"`
	CosAccountNo              sql.NullString `bun:"cos_account_no,type:varchar(80)"`
	DeleteFlag                string         `bun:"delete_flag,type:char"`
	DateCreated               time.Time      `bun:"date_created,type:datetime"`
	DateLastModified          time.Time      `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string         `bun:"last_maintained_by,type:varchar(30)"`
	ProductGroupUid           int32          `bun:"product_group_uid,type:int,autoincrement"`
	EnvironmentalFeeFlag      string         `bun:"environmental_fee_flag,type:char,default:'N'"`
	AdminFeeFlag              string         `bun:"admin_fee_flag,type:char,default:'N'"`
	IncludeInSizeAnalysisFlag sql.NullString `bun:"include_in_size_analysis_flag,type:char,default:'Y'"`
	PriceRoundingFlag         string         `bun:"price_rounding_flag,type:char,default:'N'"`
	ApplyMinMarginRulesFlag   string         `bun:"apply_min_margin_rules_flag,type:char,default:'N'"`
	EnableLineProfitWarning   string         `bun:"enable_line_profit_warning,type:char,default:'N'"`
	OeSkipProfitCheckUnpriced string         `bun:"oe_skip_profit_check_unpriced,type:char,default:'N'"`
	OeSkipProfitCheckUncosted string         `bun:"oe_skip_profit_check_uncosted,type:char,default:'N'"`
	CompressorFlag            string         `bun:"compressor_flag,type:char,default:'N'"`
}

var _ bun.BeforeAppendModelHook = (*ProductGroup)(nil)

func (m *ProductGroup) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CompanyId = config.Configuration.App.Defaults.ProductGroup.CompanyId
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
		m.AssetAccountNo = config.Configuration.App.Defaults.ProductGroup.AssetAccountNo
		m.RevenueAccountNo = sql.NullString{
			String: config.Configuration.App.Defaults.ProductGroup.RevenueAccountNo,
			Valid:  true,
		}
		m.CosAccountNo = sql.NullString{
			String: config.Configuration.App.Defaults.ProductGroup.CosAccountNo,
			Valid:  true,
		}
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}
