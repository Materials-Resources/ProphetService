package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
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

func (b BunCatalogRepository) ListProducts(ctx context.Context, filter *domain.ProductFilter) ([]*domain.Product,
	error,
) {
	var dProducts []*domain.Product
	var mInvLoc []invLoc

	bq := b.db.NewSelect().Model(&mInvLoc).Column(new(invLoc).LimitColumns()...).Column().Relation("InvMast",
		func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Column(new(invMast).LimitColumns()...)
		},
	).Relation(
		"ProductGroup", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.ExcludeColumn("*")
		},
	)

	b.queryProductsWithFilter(context.Background(), bq, filter)

	err := bq.Scan(ctx)
	if err != nil {
		return nil, err
	}

	for _, m := range mInvLoc {
		d := domain.Product{}
		m.WriteToDomain(&d)
		dProducts = append(dProducts, &d)
	}

	return dProducts, nil

}

func (b BunCatalogRepository) FilterProductByGroup(filter *domain.ProductFilter) ([]*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) SelectProduct(ctx context.Context, id string) (*domain.Product, error) {
	d := domain.Product{}
	im := new(invMast)

	// Convert id to int
	dbID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID provided was not a integer")
	}

	// Retrieve product by id
	err = b.db.NewSelect().Model(im).Where("inv_mast.inv_mast_uid = ?", dbID).Scan(ctx)
	if err != nil {
		return nil, errors.New("could not find requested product")
	}

	im.WriteToDomain(&d)

	return &d, nil
}

func (b BunCatalogRepository) SelectProductSupplier(
	ctx context.Context,
	productId, supplierId string) (*domain.ProductSupplier, error) {
	pId, err := strconv.Atoi(productId)
	if err != nil {
		return nil, errors.New("invalid productId")
	}

	sId, err := strconv.ParseFloat(supplierId, 64)
	if err != nil {
		return nil, err
	}

	is := new(inventorySupplier)

	err = b.db.NewSelect().Model(is).Where("inv_mast_uid = ? AND supplier_id = ?", pId, sId).Scan(ctx)
	if err != nil {
		return nil, err
	}

	dps := domain.ProductSupplier{}

	is.WriteToDomain(&dps)

	return &dps, nil
}

func (b BunCatalogRepository) AddProductSupplier(ctx context.Context, d *domain.ValidatedProductSupplier) error {
	is := &inventorySupplier{
		LeadTimeDays:     0,
		DeleteFlag:       "N",
		LastMaintainedBy: "admin",
		BackhaulType:     "R",
		CreatedBy:        sql.NullString{String: "admin", Valid: true},
	}

	// Get the next inventory_supplier_uid
	err := b.db.NewRaw(
		`DECLARE @id int
			EXEC @id = p21_get_counter 'inventory_supplier', 1
			SELECT @id`,
	).Scan(ctx, &is.InventorySupplierUid)

	err = is.FromDomain(&d.ProductSupplier)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = b.db.NewInsert().Model(is).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) CreateProductSupplier(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) UpdateProductSupplier(
	ctx context.Context, pId, sId string,
	p *domain.ProductSupplierPatch) error {

	is := new(inventorySupplier)

	productId, err := strconv.Atoi(pId)
	if err != nil {
		return errors.New("invalid productId")
	}

	supplierId, err := strconv.ParseFloat(sId, 64)
	if err != nil {
		return err
	}

	err = b.db.NewSelect().Model(is).Where("inv_mast_uid = ? AND supplier_id = ?", productId,
		supplierId,
	).Scan(ctx)
	if err != nil {
		return err
	}

	is.FromDomainPatch(p)

	//	err = b.db.NewSelect().Model(is).Where("inv_mast_uid = ? AND supplier_id = ?", productId,
	//		supplierId,
	//	).Scan(ctx)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if p.Delete {
	//		var orderCount int32
	//		b.db.NewRaw(`
	//SELECT COUNT(*)
	//FROM inventory_supplier
	//         INNER JOIN oe_line on inventory_supplier.inv_mast_uid = oe_line.inv_mast_uid AND
	//                               oe_line.supplier_id = inventory_supplier.supplier_id
	//         LEFT JOIN quote_line on oe_line.oe_line_uid = quote_line.oe_line_uid
	//         INNER JOIN inv_mast on inventory_supplier.inv_mast_uid = inv_mast.inv_mast_uid
	//
	//WHERE inventory_supplier.supplier_id = ? AND oe_line.inv_mast_uid = ?
	//  AND (oe_line.complete != 'Y' OR quote_line.line_complete_flag != 'Y')`, supplierId, productId,
	//		).Exec(ctx, &orderCount)
	//		var poCount int32
	//		b.db.NewRaw(`
	//SELECT COUNT(*)
	//FROM inventory_supplier
	//         INNER JOIN inv_mast on inventory_supplier.inv_mast_uid = inv_mast.inv_mast_uid
	//         INNER JOIN po_hdr on inventory_supplier.supplier_id = po_hdr.supplier_id
	//         INNER JOIN po_line on inventory_supplier.inv_mast_uid = po_line.inv_mast_uid AND po_line.po_no = po_hdr.po_no
	//WHERE inventory_supplier.supplier_id = ? and inventory_supplier.inv_mast_uid = ? and po_line.complete != 'Y'`,
	//			supplierId, productId,
	//		).Exec(ctx, &orderCount)
	//		if orderCount == 0 && poCount == 0 {
	//			existing.DeleteFlag = "Y"
	//		}
	//	} else {
	//		existing.DeleteFlag = "N"
	//	}
	//	existing.LastMaintainedBy = "admin"
	//
	_, err = b.db.NewUpdate().Model((*prophet_21_1_4559.InventorySupplier)(is)).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) SetPrimaryProductSupplier(ctx context.Context, productId, supplierId string) error {
	sIdF, err := strconv.ParseFloat(supplierId, 64)
	if err != nil {
		return err
	}

	// Get existing InventorySuppliers from an inv_mast_uid
	var existingInventorySuppliers []prophet_21_1_4559.InventorySupplier
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

		var uInventorySupplierXLoc []prophet_21_1_4559.InventorySupplierXLoc
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

func (b BunCatalogRepository) ListGroup() ([]*domain.ProductGroup, error) {
	var ds []*domain.ProductGroup
	var msProductGroup []productGroup

	ctx := context.Background()
	// Retrieve all product groups
	err := b.db.NewSelect().Model(&msProductGroup).Order("product_group_id ASC").Scan(ctx)
	if err != nil {
		return nil, errors.New("could not list product groups")
	}

	// Iterate through msProductGroup and convert to domain.ProductGroup
	for _, m := range msProductGroup {
		d := domain.ProductGroup{}
		m.WriteToDomain(&d)
		ds = append(ds, &d)
	}

	return ds, nil
}

func (b BunCatalogRepository) ReadProductByGroup(id string) ([]*domain.Product, error) {
	var dProducts []*domain.Product
	var mInvLoc []invLoc
	ctx := context.Background()

	err := b.db.NewSelect().Model(&mInvLoc).Column("inv_mast_uid").Relation("InvMast",
		func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("item_id", "item_desc")
		},
	).Where("inv_loc.product_group_id = ?",
		id,
	).Scan(ctx)

	if err != nil {
		return nil, errors.New("could not find products for product group")
	}

	for _, m := range mInvLoc {
		dProducts = append(dProducts, &domain.Product{
			Name: m.InvMast.ItemDesc,
			ID:   m.InvMastUid,
			SN:   m.InvMast.ItemId,
		},
		)

	}
	return dProducts, nil
}

func (b BunCatalogRepository) UpdateProduct(ctx context.Context, p *domain.ValidatedProduct) {
	im := new(invMast)

	b.db.NewSelect().Model(im).Relation("InvLocItems").Where("inv_mast_uid = ?", p.ID).Scan(ctx)

	fmt.Println(im)
	fmt.Println(im.InvLocItems)

}

func (b BunCatalogRepository) DeleteProduct(ctx context.Context, id string) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()

	span.SetName("repository.catalog.deleteProduct")

	dbId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("could not parse id")
	}

	var inventorySupplierXLoc []prophet_21_1_4559.InventorySupplierXLoc
	var pricePageXBook []prophet_21_1_4559.PricePageXBook

	if !b.canDeleteProduct(ctx, dbId) {
		return errors.New("product cannot be deleted")
	}

	err = b.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		im := new(invMast)

		if err := tx.NewSelect().Model(im).Where("inv_mast_uid = ?", dbId).Column("inv_mast_uid").Relation(
			"InvLocItems",
			func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.Column("inv_mast_uid", "company_id", "location_id")
			},
		).Relation("AlternateCodes", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("inv_mast_uid", "alternate_code")
		},
		).Scan(ctx); err != nil {
			return err
		}

		if im.AlternateCodes != nil {
			if _, err := tx.NewDelete().Model((*prophet_19_1_3668.AlternateCode)(nil)).Where("inv_mast_uid = ?", dbId).Exec(ctx); err != nil {
				return err
			}
		}

		ah := new(assemblyHdr)
		if err := tx.NewSelect().Model(ah).Where("inv_mast_uid = ?",
			dbId,
		).Column("inv_mast_uid").Relation("AssemblyLineItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("assembly_line_uid", "inv_mast_uid")
		},
		).Scan(ctx); err != nil {
			if !errors.As(err, &sql.ErrNoRows) {
				return err
			}

		} else {
			if _, err := tx.NewDelete().Model(ah.AssemblyLineItems).WherePK("assembly_line_uid").Exec(ctx); err != nil {
				return err
			}
			if _, err := tx.NewDelete().Model(ah).WherePK().Exec(ctx); err != nil {
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
	pg := new(productGroup)

	pg.WithDefaults()
	pg.FromDomain(&d.ProductGroup)

	_, err := b.db.NewInsert().Model(pg).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) UpdateGroup(ctx context.Context, d *domain.ProductGroup) error {
	pg := new(productGroup)

	err := b.db.NewSelect().Model(pg).Where("product_group_id = ?", d.SN).Scan(ctx)
	if err != nil {
		return err
	}

	pg.FromDomain(d)
	if _, err := b.db.NewUpdate().Model(pg).WherePK().Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (b BunCatalogRepository) FindGroupByID(ctx context.Context, id string) (*domain.ProductGroup, error) {
	pg := new(productGroup)

	err := b.db.NewSelect().Model(pg).Column("product_group_id",
		"product_group_desc",
		"product_group_uid",
	).Where("product_group.product_group_uid = ?",
		id,
	).Scan(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not find requested product group")
	}

	dpg := domain.ProductGroup{}

	pg.WriteToDomain(&dpg)

	return &dpg, nil
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

func (b BunCatalogRepository) queryProductsWithFilter(
	ctx context.Context,
	query *bun.SelectQuery,
	filter *domain.ProductFilter) {
	if filter.GroupID != "" {
		query = query.Where("product_group."+
			"product_group_uid = ?",
			filter.GroupID,
		)
	}

	query = query.Limit(filter.Limit).Order("inv_mast_uid ASC")
}
