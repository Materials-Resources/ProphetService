package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/core/domain"
	"github.com/materials-resources/s_prophet/core/port/repository"
	"github.com/materials-resources/s_prophet/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type productRepository struct {
	bun *bun.DB
	db  repository.Database
}

func (r productRepository) CreateProduct(ctx context.Context) {
	var productID int
	err := r.bun.NewRaw(
		"Declare @product_id int\n"+
			"exec @product_id = p21_get_counter @strCounterID = 'inv_mast', @iIncrementValue = 1\n"+
			"SELECT @product_id\n",
	).Scan(
		ctx,
		&productID,
	)
	fmt.Println(productID)
	if err != nil {
		fmt.Println(err)
	}
}
func (r productRepository) ReadProduct(ctx context.Context, id string) (*domain.Product, error) {
	invMast := new(model.InvMast)

	pk, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	err = r.bun.NewSelect().Model(invMast).Relation("InvLoc").Where(
		"inv_mast.inv_mast_uid = ?",
		pk,
	).Scan(ctx)
	if err != nil {
		return nil, err
	}
	product := invMastToDomain(invMast)

	return &product, nil

}

func (r productRepository) Update(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (r productRepository) Delete(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (r productRepository) CreateProductGroup(
	ctx context.Context, domain *domain.ProductGroup,
) error {
	productGroup := model.ProductGroup{
		ProductGroupId:   domain.SN,
		ProductGroupDesc: domain.Name,
	}
	_, err := r.bun.NewInsert().Model(&productGroup).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r productRepository) ReadProductBySupplier(
	ctx context.Context, supplierId string,
	supplierPartNo string,
) (*string, error) {

	parsedSupplierId, err := strconv.ParseFloat(
		supplierId,
		64,
	)
	if err != nil {
		return nil, err
	}

	inventorySupplier := new(model.InventorySupplier)
	err = r.bun.NewSelect().Model(inventorySupplier).Relation("InvMast").Where(
		"supplier_id = ?",
		parsedSupplierId,
	).Where(
		"supplier_part_no = ?",
		supplierPartNo,
	).Scan(ctx)

	return &inventorySupplier.InvMast.ItemId, nil
}

func (r productRepository) ReadProductGroup(ctx context.Context, id string) (*domain.ProductGroup, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db repository.Database) repository.ProductRepository {
	bundb := bun.NewDB(
		db.GetDB(),
		mssqldialect.New(),
	)
	bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return &productRepository{
		db:  db,
		bun: bundb,
	}
}

func invMastToDomain(invMast *model.InvMast) domain.Product {
	product := domain.Product{
		Name:           invMast.ItemDesc,
		Description:    invMast.ExtendedDesc.String,
		SN:             invMast.ItemId,
		ID:             strconv.Itoa(int(invMast.InvMastUid)),
		Price:          invMast.Price1.Float64,
		UnitsAvailable: invMast.InvLoc.QtyOnHand.Float64,
	}
	return product
}
