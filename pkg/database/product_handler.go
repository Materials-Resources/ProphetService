package database

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type ProductHandler struct {
	db *bun.DB
}

type DefaultValues struct {
	CompanyId string
}

func NewProductHandler(db *bun.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) SelectProductGroups(ctx context.Context) (*[]ProductGroup, error) {
	var productGroups []ProductGroup

	err := h.db.NewSelect().Model(&productGroups).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query product groups: %v", err)
	}

	return &productGroups, nil

}

func (h *ProductHandler) SelectProductGroup(ctx context.Context) {
}

func (h *ProductHandler) SelectProductsInProductGroup(ctx context.Context, id string) (*[]ProductMaster, error) {
	var products []ProductMaster

	err := h.db.NewSelect().Model(&products).Where("default_product_group = ?", id).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query for products in product group %s: %v", id, err)
	}

	return &products, nil
}

type InsertProductGroupInput struct {
	ProductGroupId   string
	ProductGroupDesc string
}

func (h *ProductHandler) SelectProduct(ctx context.Context, id int32) (*ProductMaster, error) {
	model := new(ProductMaster)

	if err := h.db.NewSelect().Model(model).Where("inv_mast_uid = ?", id).Scan(ctx); err != nil {
		return nil, fmt.Errorf("could not select product: %v", err)
	}

	return model, nil
}

func (h *ProductHandler) InsertProductGroup(ctx context.Context, input InsertProductGroupInput) error {
	productGroup := &ProductGroup{
		CompanyId:        "MRS",
		ProductGroupId:   "",
		ProductGroupDesc: "",
		AssetAccountNo:   "",
		RevenueAccountNo: "",
		CosAccountNo:     "",
		DateCreated:      time.Now(),
		DateLastModified: time.Now(),
		LastMaintainedBy: "ADMIN",
	}

	res, err := h.db.NewInsert().Model(productGroup).
		Column(
			"company_id",
			"product_group_id",
			"product_group_desc",
			"asset_account_no",
			"revenue_account_no",
			"cos_account_no",
			"date_created",
			"date_last_modified",
			"last_maintained_by",
		).Exec(ctx)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	return nil
}
