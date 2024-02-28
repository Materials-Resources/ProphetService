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

func (b BunCatalogRepository) AddProductSupplier(ctx context.Context, supplier *domain.ValidatedProductSupplier) error {
	inventorySupplier := &prophet_19_1_3668.InventorySupplier{
		LeadTimeDays:     0,
		DeleteFlag:       "N",
		LastMaintainedBy: "admin",
		BackhaulType:     "R",
		CreatedBy:        sql.NullString{String: "admin", Valid: true},
	}

	err := b.db.NewRaw(
		`DECLARE @id int
			EXEC @id = p21_get_counter 'inventory_supplier', 1
			SELECT @id`,
	).Scan(ctx, &inventorySupplier.InventorySupplierUid)

	err = ProductSupplierFromDomain(&supplier.ProductSupplier, inventorySupplier)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = b.db.NewInsert().Model(inventorySupplier).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) CreateProductSupplier(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) GetProductSupplier(
	ctx context.Context,
	productId, supplierId string) (*domain.ProductSupplier, error) {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) UpdateProductSupplier(ctx context.Context, p *domain.ValidatedProductSupplier) error {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) SetPrimaryProductSupplier(ctx context.Context, productId, supplierId string) error {
	sIdF, err := strconv.ParseFloat(supplierId, 64)
	if err != nil {
		return err
	}

	// Get existing InventorySuppliers from an inv_mast_uid
	var existingInventorySuppliers []prophet_19_1_3668.InventorySupplier
	err = b.db.NewSelect().Model(&existingInventorySuppliers).Where("inv_mast_uid = ?", productId).Relation("InventorySupplierXLoc").Scan(
		ctx,
	)
	if err != nil {
		return err
	}

	// Check if the given supplierId exists in existingInventorySupplier
	var supplierFound = false
	for _, supplier := range existingInventorySuppliers {
		if supplier.SupplierId == sIdF {
			supplierFound = true
		}
	}
	if supplierFound != true {
		return errors.New("could not find supplier")
	}

	return b.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		var uInventorySupplierXLoc []prophet_19_1_3668.InventorySupplierXLoc
		for _, supplier := range existingInventorySuppliers {

			if supplier.SupplierId == sIdF {
				if supplier.InventorySupplierXLoc != nil {
					supplier.InventorySupplierXLoc.PrimarySupplier = "Y"
					uInventorySupplierXLoc = append(uInventorySupplierXLoc, *supplier.InventorySupplierXLoc)
				} else {
					uid, err := b.getInventorySupplierXLocUid(ctx)
					if err != nil {
						return err
					}
					nIXL := &prophet_19_1_3668.InventorySupplierXLoc{
						InventorySupplierXLocUid: int32(uid),
						InventorySupplierUid:     supplier.InventorySupplierUid,
						LocationId:               1001,
						PrimarySupplier:          "Y",
						RowStatusFlag:            704,
						LastMaintainedBy:         "admin",
						OverrideVmiStatusFlag:    "N",
					}
					if _, err = tx.NewInsert().Model(nIXL).Exec(ctx); err != nil {
						return err
					}
				}
			} else {
				if supplier.InventorySupplierXLoc != nil {
					supplier.InventorySupplierXLoc.PrimarySupplier = "N"
					uInventorySupplierXLoc = append(uInventorySupplierXLoc, *supplier.InventorySupplierXLoc)
				}
			}
		}

		if _, err = tx.NewUpdate().Model(&uInventorySupplierXLoc).Bulk().Exec(ctx); err != nil {
			return err
		}
		return nil
	},
	)
}

func (b BunCatalogRepository) ListProduct(ctx context.Context, cursor int32, count int) (res []*domain.Product,
	nextCursor int,
	err error,
) {
	var m []prophet_19_1_3668.InvLoc

	b.db.NewSelect().Model(&m).Relation("InvMast").Relation("ProductGroup").Where("inv_loc.inv_mast_uid > ?",
		cursor,
	).Order("inv_mast_uid ASC").Limit(count).Scan(ctx)

	return FromDBListProduct(m), int(m[len(m)-1].InvMastUid), nil
}

func (b BunCatalogRepository) CreateProduct() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) FindProductByID(id string) (*domain.ValidatedProduct, error) {
	ctx := context.Background()
	invMast := NewInvMast()

	dbID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID provided was not a integer")
	}

	err = b.db.NewSelect().Model(invMast.GetModel()).Where("inv_mast.inv_mast_uid = ?", dbID).Scan(ctx)
	if err != nil {
		return nil, errors.New("could not find requested product")
	}
	return FromDBProduct(invMast.GetModel())
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

func (b BunCatalogRepository) ReadProductByGroup(id string) ([]*domain.Product, error) {
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
	return FromDBListProduct(dbInvLoc), nil
}

func (b BunCatalogRepository) UpdateProduct(ctx context.Context, p *domain.ValidatedProduct) {
	invMast := new(prophet_19_1_3668.InvMast)

	idInt, err := strconv.Atoi(p.ID)
	if err != nil {
		return
	}

	b.db.NewSelect().Model(&invMast).Relation("InvLocItems").Where("inv_mast_uid = ?", idInt).Scan(ctx)

	fmt.Println(invMast)
	fmt.Println(invMast.InvLocItems)

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

		// Retrieve records for ItemCategoryXInvMast and delete them if they exist
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

		// Retrieve records for InvLocStockStatus and delete them if they exist
		var invLocStockStatus []prophet_19_1_3668.InvLocStockStatus
		if err := tx.NewSelect().Model(&invLocStockStatus).Column("inv_loc_stock_status_uid").Where("inv_mast_uid = ?",
			dbId,
		).Scan(ctx); err != nil {
			return err
		}
		if len(invLocStockStatus) > 0 {
			if _, err := tx.NewDelete().Model(&invLocStockStatus).WherePK().Exec(ctx); err != nil {
				return err
			}
		}

		// Retrieve records for ItemIdChangeHistory and delete them if they exist
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

func (b BunCatalogRepository) CreateGroup(ctx context.Context, d *domain.ValidatedProductGroup) error {
	m := NewProductGroup().WithDefaults().FromDomain(&d.ProductGroup)

	_, err := b.db.NewInsert().Model(m.model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) UpdateGroup(ctx context.Context, group *domain.ProductGroup) error {
	m := NewProductGroup()

	err := b.db.NewSelect().Model(m.GetModel()).Where("product_group_id = ?", group.SN).Scan(ctx)
	if err != nil {
		return err
	}

	m = m.FromDomain(group)
	if _, err := b.db.NewUpdate().Model(m.GetModel()).WherePK().Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) FindGroupByID(ctx context.Context, id string) (*domain.ValidatedProductGroup, error) {
	m := NewProductGroup()

	err := b.db.NewSelect().Model(m.GetModel()).Column("product_group_id", "product_group_desc",
		"product_group_uid",
	).Where("product_group.product_group_id = ?",
		id,
	).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not find requested product group")
	}
	return FromDBProductGroup(m.GetModel())
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

func (b BunCatalogRepository) getInventorySupplierXLocUid(ctx context.Context) (id int, err error) {
	err = b.db.NewRaw(
		`DECLARE @id int
			EXEC @id = p21_get_counter 'inventory_supplier_x_loc', 1
			SELECT @id`,
	).Scan(ctx, &id)
	return id, err
}
