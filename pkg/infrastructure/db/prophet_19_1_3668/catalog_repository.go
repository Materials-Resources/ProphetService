package prophet_19_1_3668

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/domain/repositories"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
	"github.com/uptrace/bun"
)

func NewBunCatalogRepository(db *bun.DB) repositories.CatalogRepository {
	return &BunCatalogRepository{db: db}
}

type BunCatalogRepository struct {
	db *bun.DB
}

func (b BunCatalogRepository) CreateProduct() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) FindProductByID(id string) (*entities.ValidatedProduct, error) {
	ctx := context.Background()
	dbProduct := new(model.InvMast)

	dbID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID provided was not a integer")
	}

	err = b.db.NewSelect().Model(dbProduct).Where("inv_mast.inv_mast_uid = ?", dbID).Scan(ctx)
	if err != nil {
		return nil, errors.New("could not find requested product")
	}
	return FromDBProduct(dbProduct)
}
func (b BunCatalogRepository) ListGroup() ([]*entities.ValidatedProductGroup, error) {
	var dbProductGroup []model.ProductGroup
	ctx := context.Background()
	err := b.db.NewSelect().Model(&dbProductGroup).Scan(ctx)
	if err != nil {
		return nil, errors.New("could not list product groups")
	}
	return FromDBListGroup(dbProductGroup)
}
func (b BunCatalogRepository) ReadProductByGroup(id string) ([]*entities.ValidatedProduct, error) {
	var dbInvLoc []model.InvLoc
	ctx := context.Background()

	err := b.db.NewSelect().Model(&dbInvLoc).Column("inv_mast_uid").Relation("InvMast",
		func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("item_id", "item_desc")
		},
	).Where("inv_loc.product_group_id = ?",
		id,
	).Scan(ctx)

	if err != nil {
		return nil, errors.New("could not find products for product group")
	}
	return FromDBListProduct(dbInvLoc)
}

func (b BunCatalogRepository) UpdateProduct() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) DeleteProduct(ctx context.Context, id string) error {

	dbId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("could not parse id")
	}

	var inventorySupplierXLoc []model.InventorySupplierXLoc
	var pricePageXBook []model.PricePageXBook

	if !b.canDeleteProduct(ctx, dbId) {
		return errors.New("product cannot be deleted")
	}

	err = b.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		invMast := new(model.InvMast)

		if err := tx.NewSelect().Model(invMast).Where("inv_mast_uid = ?", dbId).Column("inv_mast_uid").Relation("InvLocItems",
			func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.Column("inv_mast_uid", "company_id", "location_id")
			},
		).Relation("AlternateCodes", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("inv_mast_uid", "alternate_code")
		},
		).Scan(ctx); err != nil {
			return err
		}

		if invMast.AlternateCodes != nil {
			if _, err := tx.NewDelete().Model((*model.AlternateCode)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
				return err
			}
		}

		assemblyHdr := new(model.AssemblyHdr)
		if err := tx.NewSelect().Model(assemblyHdr).Where("inv_mast_uid = ?",
			dbId,
		).Column("inv_mast_uid").Relation("AssemblyLineItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("assembly_line_uid", "inv_mast_uid")
		},
		).Scan(ctx); err != nil {
			return err
		}

		if assemblyHdr.AssemblyLineItems != nil {
			if _, err := tx.NewDelete().Model(&assemblyHdr.AssemblyLineItems).WherePK("assembly_line_uid").Exec(ctx); err != nil {
				return err
			}
		}

		if _, err := tx.NewDelete().Model(assemblyHdr).WherePK().Exec(ctx); err != nil {
			return err
		}

		if err := tx.NewSelect().Model(&inventorySupplierXLoc).Column("inventory_supplier_x_loc_uid").Relation(
			"InventorySupplier",
			func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.ExcludeColumn("*").Where("inv_mast_uid = ?", dbId)
			},
		).Scan(ctx); err != nil {
			return err
		}

		if inventorySupplierXLoc != nil {
			if _, err := tx.NewDelete().Model(&inventorySupplierXLoc).WherePK().Exec(ctx); err != nil {
				return err
			}
		}

		if err := tx.NewSelect().Model(&pricePageXBook).Column("price_page_x_book_uid").Relation("PricePage",
			func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.ExcludeColumn("*").Where("inv_mast_uid = ?", dbId)
			},
		).Scan(
			ctx,
		); err != nil {
			return err
		}

		if pricePageXBook != nil {
			if _, err := tx.NewDelete().Model(&pricePageXBook).WherePK().Exec(ctx); err != nil {
				return err
			}
			if _, err := tx.NewDelete().Model((*model.PricePage)(nil)).Where("inv_mast_uid = ?",
				dbId,
			).Exec(ctx); err != nil {
				return err
			}
		}

		if _, err := tx.NewDelete().Model((*model.ItemConversion)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvAdjLine)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvTran)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.ItemUom)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}

		if _, err := tx.NewDelete().Model((*model.AverageInventoryValue)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.ItemCategoryXInvMast)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvLocMsp)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InventorySupplier)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvLocStockStatus)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}

		if _, err := tx.NewDelete().Model((*model.ItemIdChangeHistory)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvBin)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvLoc)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*model.InvMast)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		return nil
	},
	)

	return err
}

func (b BunCatalogRepository) CreateGroup() {
	//productGroup := &model.ProductGroup{
	//	CompanyId:        "MRS",
	//	ProductGroupId:   domain.SN,
	//	ProductGroupDesc: domain.Name,
	//	AssetAccountNo:   "121001",
	//	DeleteFlag:       "N",
	//	RevenueAccountNo: sql.NullString{String: "401001", Valid: true},
	//	CosAccountNo:     sql.NullString{String: "501001", Valid: true},
	//	DateCreated:      time.Now(),
	//	DateLastModified: time.Now(),
	//	LastMaintainedBy: "admin",
	//}
	panic("implement me")
}

func (b BunCatalogRepository) FindGroupByID(ctx context.Context, id string) (*entities.ValidatedProductGroup, error) {
	dbProductGroup := new(model.ProductGroup)

	dbID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID provided was not a integer")
	}

	err = b.db.NewSelect().Model(dbProductGroup).Column("product_group_id", "product_group_desc", "product_group_uid").Where("product_group.product_group_uid = ?",
		dbID,
	).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not find requested product group")
	}
	return FromDBProductGroup(dbProductGroup)
}

func (b BunCatalogRepository) UpdateGroup() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) DeleteGroup() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) canDeleteProduct(ctx context.Context, id int) bool {

	q :=
		`select im.inv_mast_uid, purchase_history.purchase_count, sales_history.sales_count
FROM inv_mast im
LEFT JOIN (SELECT inventory_receipts_line.inv_mast_uid, count(inventory_receipts_line.inv_mast_uid) as purchase_count FROM inventory_receipts_line
           group by inventory_receipts_line.inv_mast_uid) purchase_history on purchase_history.inv_mast_uid = im.inv_mast_uid
LEFT JOIN (SELECT oe_line.inv_mast_uid, count(oe_line.inv_mast_uid) as sales_count FROM oe_line
           group by oe_line.inv_mast_uid) sales_history on sales_history.inv_mast_uid = im.inv_mast_uid
WHERE im.inv_mast_uid = ?`

	type productWithHistory struct {
		InvMastUid    int32
		PurchaseCount int32
		SalesCount    int32
	}

	p := new(productWithHistory)
	b.db.NewRaw(q, id).Scan(ctx, &p.InvMastUid, &p.PurchaseCount, &p.SalesCount)

	if p.SalesCount == 0 && p.PurchaseCount == 0 {
		return true
	}

	return false
}
