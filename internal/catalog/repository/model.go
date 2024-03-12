package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

type inventorySupplier prophet_21_1_4559.InventorySupplier

// FromDomain maps domain.ProductSupplier to inventorySupplier
func (i *inventorySupplier) FromDomain(d *domain.ProductSupplier) error {
	if i.DivisionId == 0 {
		if id, err := strconv.ParseFloat(d.SupplierId, 64); err != nil {
			return errors.New("could not convert supplier id")
		} else {
			i.DivisionId = id
		}
	}
	if d.Delete {
		i.DeleteFlag = "Y"
	} else {
		i.DeleteFlag = "N"
	}

	i.ListPrice = d.ListPrice
	i.Cost = d.PurchasePrice
	i.SupplierPartNo = sql.NullString{String: d.SupplierSn, Valid: true}
	if id, err := strconv.ParseFloat(d.SupplierId, 64); err != nil {
		return errors.New("could not convert supplier id")
	} else {
		i.SupplierId = id
	}
	if id, err := strconv.Atoi(d.ProductId); err != nil {
		return errors.New("could not convert product id")
	} else {
		i.InvMastUid = int32(id)
	}
	return nil
}

// FromDomainPatch maps domain.ProductSupplierPatch to inventorySupplier
func (i *inventorySupplier) FromDomainPatch(p *domain.ProductSupplierPatch) {
	if p.SupplierProductSn != nil {
		i.SupplierPartNo = sql.NullString{String: *p.SupplierProductSn, Valid: true}
	}
	if p.PurchasePrice != nil {
		i.Cost = *p.PurchasePrice
	}
	if p.ListPrice != nil {
		i.ListPrice = *p.ListPrice
	}
	if p.Delete != nil {
	}
}

// WriteToDomain maps inventorySupplier data to domain.ProductSupplier
func (i *inventorySupplier) WriteToDomain(d *domain.ProductSupplier) {
	d.ProductId = strconv.Itoa(int(i.InvMastUid))
	d.SupplierId = strconv.FormatFloat(i.SupplierId, 'f', -1, 64)
	d.ListPrice = i.ListPrice
	d.PurchasePrice = i.Cost
	d.SupplierSn = i.SupplierPartNo.String
}

type productGroup prophet_21_1_4559.ProductGroup

func (p *productGroup) WithDefaults() *productGroup {
	p.CompanyId = "MRS"
	p.AssetAccountNo = "121001"
	p.DeleteFlag = "N"
	p.RevenueAccountNo = sql.NullString{String: "401001", Valid: true}
	p.CosAccountNo = sql.NullString{String: "501001", Valid: true}
	p.LastMaintainedBy = "admin"

	return p
}

func (p *productGroup) FromDomain(d *domain.ProductGroup) *productGroup {
	if d.SN != "" {
		p.ProductGroupId = d.SN
	}
	if d.Name != "" {
		p.ProductGroupId = d.SN
	}
	p.ProductGroupDesc = d.Name
	return p
}

// WriteToDomain maps productGroup data to domain.ProductGroup
func (p *productGroup) WriteToDomain(d *domain.ProductGroup) {
	d.Name = p.ProductGroupDesc
	d.SN = p.ProductGroupId
	d.ID = strconv.Itoa(int(p.ProductGroupUid))
}

type invMast prophet_21_1_4559.InvMast

func (i *invMast) LimitColumns() []string {
	return []string{"item_id", "item_desc", "extended_desc", "inv_mast_uid"}
}

// WriteToDomain maps invMast data to domain.Product
func (i *invMast) WriteToDomain(d *domain.Product) {
	d.Name = i.ItemDesc
	d.SN = i.ItemId
	d.Description = i.ExtendedDesc.String
	d.ID = strconv.Itoa(int(i.InvMastUid))
}

type invLoc prophet_21_1_4559.InvLoc

// LimitColumns Get slice of columns to be selected
func (i *invLoc) LimitColumns() []string {
	return []string{"inv_mast_uid", "product_group_id"}

}

// WriteToDomain maps invLoc data to domain.Product
func (i *invLoc) WriteToDomain(d *domain.Product) {
	d.Name = i.InvMast.ItemDesc
	d.Description = i.InvMast.ExtendedDesc.String
	d.ID = strconv.Itoa(int(i.InvMastUid))
	d.SN = i.InvMast.ItemId
	d.ProductGroupName = i.ProductGroup.ProductGroupDesc
}

type assemblyHdr prophet_21_1_4559.AssemblyHdr
