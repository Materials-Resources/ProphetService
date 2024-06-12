package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CorpId struct {
	bun.BaseModel    `bun:"table:corp_id"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	AddressId        float64   `bun:"address_id,type:decimal(19,0),pk"`
	CreditLimit      float64   `bun:"credit_limit,type:decimal(19,4)"`
	CreditLimitUsed  float64   `bun:"credit_limit_used,type:decimal(19,4)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	AddressName      string    `bun:"address_name,type:varchar(255),nullzero"`
}
