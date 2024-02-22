package repository

import (
	"database/sql"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

type ProductGroup struct {
	model *prophet_19_1_3668.ProductGroup
}

func NewProductGroup() *ProductGroup {
	return &ProductGroup{model: new(prophet_19_1_3668.ProductGroup)}
}

func (p *ProductGroup) WithDefaults() *ProductGroup {
	p.model.CompanyId = "MRS"
	p.model.AssetAccountNo = "121001"
	p.model.DeleteFlag = "N"
	p.model.RevenueAccountNo = sql.NullString{String: "401001", Valid: true}
	p.model.CosAccountNo = sql.NullString{String: "501001", Valid: true}
	p.model.LastMaintainedBy = "admin"

	return p
}

func (p *ProductGroup) FromDomain(d *domain.ProductGroup) *ProductGroup {
	if d.SN != "" {
		p.model.ProductGroupId = d.SN
	}
	if d.Name != "" {
		p.model.ProductGroupId = d.SN
	}
	p.model.ProductGroupDesc = d.Name
	return p
}

func (p *ProductGroup) GetModel() *prophet_19_1_3668.ProductGroup {
	return p.model
}

type InvMast struct {
	model *prophet_19_1_3668.InvMast
}

func NewInvMast() *InvMast {
	return &InvMast{model: new(prophet_19_1_3668.InvMast)}
}

func (i *InvMast) GetModel() *prophet_19_1_3668.InvMast {
	return i.model
}

type InvLoc struct {
	model *prophet_19_1_3668.InvLoc
}

func NewInvLoc() *InvLoc {
	return &InvLoc{model: new(prophet_19_1_3668.InvLoc)}
}

func (i *InvLoc) GetModel() *prophet_19_1_3668.InvLoc {
	return i.model
}
