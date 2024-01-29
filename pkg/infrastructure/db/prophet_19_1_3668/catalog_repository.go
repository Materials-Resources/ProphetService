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

	err := b.db.NewSelect().Model(&dbInvLoc).Relation("InvMast").Where("inv_loc.product_group_id = ?",
		id,
	).Scan(ctx)
	fmt.Println(len(dbInvLoc))

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("could not find products for product group")
	}
	return FromDBListProduct(dbInvLoc)
}

func (b BunCatalogRepository) UpdateProduct() {
	//TODO implement me
	panic("implement me")
}

func (b BunCatalogRepository) DeleteProduct() {
	//TODO implement me
	panic("implement me")
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

func (b BunCatalogRepository) FindGroupByID(id string) (*entities.ValidatedProductGroup, error) {
	ctx := context.Background()
	dbProductGroup := new(model.ProductGroup)

	dbID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID provided was not a integer")
	}

	err = b.db.NewSelect().Model(dbProductGroup).Where("product_group.product_group_uid = ?",
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
