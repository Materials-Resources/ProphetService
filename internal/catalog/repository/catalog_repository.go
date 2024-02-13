package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/catalog"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	"github.com/uptrace/bun"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func NewBunCatalogRepository(db *bun.DB, tp *tracesdk.TracerProvider) catalog.Repository {
	return &BunCatalogRepository{db: db, tracer: tp.Tracer("CatalogRepository")}
}

type BunCatalogRepository struct {
	db     *bun.DB
	tracer trace.Tracer
}

func (b BunCatalogRepository) CreateProduct() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) FindProductByID(id string) (*domain.ValidatedProduct, error) {
	ctx := context.Background()
	dbProduct := new(prophet_19_1_3668.InvMast)

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
func (b BunCatalogRepository) ListGroup() ([]*domain.ValidatedProductGroup, error) {
	var dbProductGroup []prophet_19_1_3668.ProductGroup
	ctx := context.Background()
	err := b.db.NewSelect().Model(&dbProductGroup).Scan(ctx)
	if err != nil {
		return nil, errors.New("could not list product groups")
	}
	return FromDBListGroup(dbProductGroup)
}
func (b BunCatalogRepository) ReadProductByGroup(id string) ([]*domain.ValidatedProduct, error) {
	var dbInvLoc []prophet_19_1_3668.InvLoc
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
	span := trace.SpanFromContext(ctx)
	defer span.End()

	span.SetName("repository.catalog.deleteProduct")

	dbId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("could not parse id")
	}

	var inventorySupplierXLoc []prophet_19_1_3668.InventorySupplierXLoc
	var pricePageXBook []prophet_19_1_3668.PricePageXBook

	if !b.canDeleteProduct(ctx, dbId) {
		return errors.New("product cannot be deleted")
	}

	err = b.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		invMast := new(prophet_19_1_3668.InvMast)

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
			if _, err := tx.NewDelete().Model((*prophet_19_1_3668.AlternateCode)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
				return err
			}
		}

		assemblyHdr := new(prophet_19_1_3668.AssemblyHdr)
		if err := tx.NewSelect().Model(assemblyHdr).Where("inv_mast_uid = ?",
			dbId,
		).Column("inv_mast_uid").Relation("AssemblyLineItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("assembly_line_uid", "inv_mast_uid")
		},
		).Scan(ctx); err != nil {
			if !errors.As(err, &sql.ErrNoRows) {
				return err
			}

		} else {
			if _, err := tx.NewDelete().Model(&assemblyHdr.AssemblyLineItems).WherePK("assembly_line_uid").Exec(ctx); err != nil {
				return err
			}
			if _, err := tx.NewDelete().Model(assemblyHdr).WherePK().Exec(ctx); err != nil {
				return err
			}
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
			if _, err := tx.NewDelete().Model((*prophet_19_1_3668.PricePage)(nil)).Where("inv_mast_uid = ?",
				dbId,
			).Exec(ctx); err != nil {
				return err
			}
		}

		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.ItemConversion)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvAdjLine)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvTran)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.ItemUom)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}

		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.AverageInventoryValue)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}

		var itemCategoryXInvMast []prophet_19_1_3668.ItemCategoryXInvMast
		if err := tx.NewSelect().Model(&itemCategoryXInvMast).Column("item_category_x_inv_mast_uid").Where("inv_mast_uid = ?",
			dbId,
		).Scan(ctx); err != nil {
			return err
		}

		if len(itemCategoryXInvMast) > 0 {
			if _, err := tx.NewDelete().Model(&itemCategoryXInvMast).WherePK().Exec(ctx); err != nil {
				return err
			}
		}

		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvLocMsp)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InventorySupplier)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvLocStockStatus)(nil)).Where("inv_mast_uid = ?",
			dbId,
		).Exec(ctx); err != nil {
			return err
		}

		var itemIdChangeHistory []prophet_19_1_3668.ItemIdChangeHistory

		if err := tx.NewSelect().Model(&itemIdChangeHistory).Column("item_id_change_history_uid").Where("inv_mast_uid = ?",
			dbId,
		).Scan(ctx); err != nil {
			return err
		}

		if len(itemIdChangeHistory) > 0 {
			if _, err := tx.NewDelete().Model(&itemIdChangeHistory).WherePK().Exec(ctx); err != nil {
				return err
			}
		}

		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvBin)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvLoc)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewDelete().Model((*prophet_19_1_3668.InvMast)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
			return err
		}
		return nil
	},
	)

	return err
}

func (b BunCatalogRepository) CreateGroup() {
	//productGroup := &prophet_19_1_3668.ProductGroup{
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

func (b BunCatalogRepository) FindGroupByID(ctx context.Context, id string) (*domain.ValidatedProductGroup, error) {
	dbProductGroup := new(prophet_19_1_3668.ProductGroup)

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
