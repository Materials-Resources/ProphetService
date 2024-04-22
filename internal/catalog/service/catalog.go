package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
	"github.com/materials-resources/s_prophet/pkg/consumer"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type EventService struct {
	Producer *EventProducer
	Consumer *EventConsumer
}

func NewCatalogService(
	models data.Models, tracer trace.Tracer, cli *kgo.Client,
	serde *sr.Serde) Catalog {

	return Catalog{
		models: models, tracer: tracer, event: EventService{
			Producer: &EventProducer{client: cli, serde: serde}, Consumer: &EventConsumer{client: cli, serde: serde},
		}}
}

type Catalog struct {
	models data.Models
	tracer trace.Tracer
	event  EventService
}

func (c Catalog) SetPrimaryProductSupplier(
	ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error {

	err := c.models.ExecTx(
		ctx, func(models *data.Models) error {

			inventorySupplier, err := models.InventorySupplier.GetBySupplierIdDivisionIdInvMastUid(
				ctx, supplierUid,
				divisionId, productUid)
			if err != nil {
				return err

			}
			inventorySupplierXLocs, err := models.InventorySupplierXLoc.GetByInvMastUidAndLocationId(
				ctx, productUid,
				locationId)

			for _, inventorySupplierXLoc := range inventorySupplierXLocs {
				if inventorySupplierXLoc.InventorySupplierUid == inventorySupplier.InventorySupplierUid {
					inventorySupplierXLoc.PrimarySupplier = "Y"
				} else {
					inventorySupplierXLoc.PrimarySupplier = "N"
				}
				fmt.Println(
					inventorySupplier.InventorySupplierUid, inventorySupplierXLoc.InventorySupplierUid,
					inventorySupplierXLoc.PrimarySupplier)
				if err := models.InventorySupplierXLoc.Update(ctx, inventorySupplierXLoc); err != nil {
					return err
				}

			}
			return nil
		})

	return err
}

func (c Catalog) RegisterWorkers() {

	cg := consumer.NewConsumerGroup()

	cg.RegisterWorkers(
		consumer.NewWorker(
			DeleteProductTopic, func(rec *kgo.Record) error {
				return c.event.Consumer.ConsumeDeleteProduct(context.Background(), rec, c.DeleteProduct)

			}))

	cg.Start([]string{"localhost:19092"}, "18")
}

func (c Catalog) RequestDeleteProduct(ctx context.Context, uid int32) error {

	err := c.event.Producer.PublishDeleteProduct(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (c Catalog) DeleteProduct(ctx context.Context, uid int32) error {
	ctx, span := c.tracer.Start(ctx, "service.catalog.DeleteProduct")
	defer span.End()

	err := c.models.ExecTx(

		ctx, func(models *data.Models) error {
			// get invMast record to see if it exists
			invMast, err := models.InvMast.Get(ctx, uid)
			if err != nil {
				return err
			}

			// check if product has receipts
			receiptCount, err := models.InventoryReceiptsLine.CountByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			if receiptCount > 0 {
				return errors.New("cannot delete product with receipts")
			}

			// check if product has orders
			orderCount, err := models.OeLine.CountByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			if orderCount > 0 {
				return errors.New("cannot delete product with orders")
			}
			invBins, err := models.InvBin.GetByInvMastUid(ctx, uid)
			for _, invBin := range invBins {
				err = models.InvBin.Delete(ctx, invBin)
				if err != nil {
					return err
				}

			}

			// query and delete alternateCode records belonging to the product
			alternateCodes, err := models.AlternateCode.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, alternateCode := range alternateCodes {
				err = models.AlternateCode.Delete(ctx, alternateCode)
				if err != nil {
					return err
				}
			}

			// query and delete assemblyHdr and assemblyLine records belonging to the product
			assemblyHdr, err := models.AssemblyHdr.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			if assemblyHdr != nil {
				assemblyLine, err := models.AssemblyLine.GetByInvMastUid(ctx, assemblyHdr.InvMastUid)
				if err != nil {
					return err
				}
				for _, line := range assemblyLine {
					err = models.AssemblyLine.Delete(ctx, line)
					if err != nil {
						return err
					}
				}
				err = models.AssemblyHdr.Delete(ctx, assemblyHdr)
			}

			// query and delete pricePage and pricePageXBook records belonging to the product
			pricePages, err := models.PricePage.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, pricePage := range pricePages {
				pricePageXBook, err := models.PricePageXBook.GetByPricePageUid(ctx, pricePage.PricePageUid)
				if err != nil {
					return err
				}
				err = models.PricePageXBook.Delete(ctx, pricePageXBook)
				if err != nil {
					return err
				}
			}

			// query and delete itemConversion records belonging to the product
			itemConversions, err := models.ItemConversion.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, conversion := range itemConversions {
				err = models.ItemConversion.Delete(ctx, conversion)
				if err != nil {
					return err
				}
			}

			// query and delete invAdjLine records belonging to the product
			invAdjLines, err := models.InvAdjLine.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, invAdjLine := range invAdjLines {
				err = models.InvAdjLine.Delete(ctx, invAdjLine)
				if err != nil {
					return err
				}
			}

			// query and delete invTran records belonging to the product
			invTrans, err := models.InvTran.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, invTran := range invTrans {
				err = models.InvTran.Delete(ctx, invTran)
				if err != nil {
					return err
				}
			}

			// query and delete itemCategoryXInvMast records belonging to the product
			itemCategoryXInvMasts, err := models.ItemCategoryXInvMast.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, itemCategoryXInvMast := range itemCategoryXInvMasts {
				err = models.ItemCategoryXInvMast.Delete(ctx, itemCategoryXInvMast)
				if err != nil {
					return err
				}
			}

			// query and delete invLocMsp records belonging to the product
			invLocMsps, err := models.InvLocMsp.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, invLocMsp := range invLocMsps {
				err = models.InvLocMsp.Delete(ctx, invLocMsp)
				if err != nil {
					return err
				}
			}

			// query and delete invLocStockStatus records belonging to the product
			invLocStockStatus, err := models.InvLocStockStatus.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, stockStatus := range invLocStockStatus {
				if err := models.InvLocStockStatus.Delete(ctx, stockStatus); err != nil {
					return err
				}
			}

			itemIdChangeHistories, err := models.ItemIdChangeHistory.GetByInvMastUid(ctx, uid)
			if err != nil {
				return err
			}
			for _, itemIdChangeHistory := range itemIdChangeHistories {
				if err := models.ItemIdChangeHistory.Delete(ctx, itemIdChangeHistory); err != nil {
					return err
				}
			}

			// query and delete itemUom records belonging to the product
			itemUoms, err := models.ItemUom.GetByInvMastUid(ctx, uid)
			for _, itemUom := range itemUoms {
				inventorySupplierXUoms, err := models.InventorySupplierXUom.GetByItemUomUid(
					ctx,
					itemUom.ItemUomUid)
				if err != nil {
					return err
				}
				for _, inventorySupplierXUom := range inventorySupplierXUoms {
					if err := models.InventorySupplierXUom.Delete(ctx, inventorySupplierXUom); err != nil {
						return err
					}
				}
				if err := models.ItemUom.Delete(ctx, itemUom); err != nil {
					return err
				}
			}

			// query and delete inventorySupplier records belonging to the product
			inventorySuppliers, err := models.InventorySupplier.GetByInvMastUid(ctx, uid)
			for _, inventorySupplier := range inventorySuppliers {

				inventorySupplierXLocs, err := models.InventorySupplierXLoc.GetByInventorySupplierUid(
					ctx, inventorySupplier.InventorySupplierUid)
				if err != nil {
					return err

				}
				for _, inventorySupplierXLoc := range inventorySupplierXLocs {
					if err := models.InventorySupplierXLoc.Delete(ctx, inventorySupplierXLoc); err != nil {
						return err
					}

				}
				if err := models.InventorySupplier.Delete(ctx, inventorySupplier); err != nil {
					return err
				}

			}

			// query and delete invLoc records belonging to the product
			invLocs, err := models.InvLoc.GetByInvMastUid(ctx, []float64{1001}, uid)
			if err != nil {
				return err
			}
			for _, invLoc := range invLocs {
				if err := models.InvLoc.Delete(ctx, &invLoc); err != nil {
					return err
				}
			}

			// query and delete averageInventoryValue records belonging to the product
			averageInventoryValues, err := models.AverageInventoryValue.GetByInvMastUid(ctx, uid)
			for _, averageInventoryValue := range averageInventoryValues {
				if err := models.AverageInventoryValue.Delete(ctx, averageInventoryValue); err != nil {
					return err
				}

			}

			// delete invMast records belonging to the product
			if err := models.InvMast.Delete(ctx, invMast); err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	return nil
}
func (c Catalog) UpdateProductSupplier(ctx context.Context, productSupplier *domain.ProductSupplier) error {
	// TODO implement me
	panic("implement me")
}

func (c Catalog) UpdateProduct(con context.Context, product *domain.Product, locations []float64) error {
	ctx, span := c.tracer.Start(con, "UpdateProduct")
	defer span.End()

	invMast, err := c.models.InvMast.Get(ctx, product.UID)
	if err != nil {
		return err
	}

	if product.Name != "" {
		invMast.ItemDesc = product.Name
	}
	if product.Description != "" {
		invMast.ExtendedDesc = sql.NullString{String: product.Description, Valid: true}
	}

	if err = c.models.InvMast.Update(ctx, invMast); err != nil {
		return err
	}
	invLocs, err := c.models.InvLoc.GetByInvMastUid(ctx, locations, product.UID)
	if err != nil {
		return err
	}
	for _, invLoc := range invLocs {
		if product.ProductGroupSn != "" {
			invLoc.LastMaintainedBy = "admin"
			invLoc.ProductGroupId = sql.NullString{String: product.ProductGroupSn, Valid: true}
		}
		if err = c.models.InvLoc.Update(ctx, &invLoc); err != nil {
			return err
		}
	}

	return nil
}

func (c Catalog) GetProductBySupplierPartNumber(
	ctx context.Context, partNumber string,
	supplier float64) (domain.Product, error) {
	invMast, err := c.models.InvMast.GetBySupplierPartNumber(ctx, supplier, partNumber)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		UID: invMast.InvMastUid, SN: invMast.ItemId, Name: invMast.ItemDesc,
	}, nil
}

func (c Catalog) ListProduct(ctx context.Context, filter domain.Filter) ([]*domain.Product, error) {
	dbFilter := &data.Filters{Cursor: filter.Cursor, Limit: filter.Limit}

	invLocs, _, err := c.models.InvLoc.GetAll(ctx, 1001, data.DeleteFlagNo, *dbFilter)
	if err != nil {
		return nil, err
	}

	products := make([]*domain.Product, len(invLocs))
	for i, invLoc := range invLocs {
		products[i] = mapInvLocToProduct(invLoc)
	}

	return products, nil
}

func mapInvLocToProduct(invLoc data.InvLoc) *domain.Product {
	return &domain.Product{
		UID: invLoc.InvMastUid, SN: invLoc.InvMast.ItemId, Name: invLoc.InvMast.ItemDesc,
		Description: invLoc.InvMast.ExtendedDesc.String,
	}
}

func (c Catalog) GetProduct(ctx context.Context, uid int32) (domain.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (c Catalog) CreateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error {

	dbProductGroup := &data.ProductGroup{
		ProductGroupDesc: productGroup.Name, ProductGroupId: productGroup.SN,
	}

	err := c.models.ProductGroup.Insert(ctx, dbProductGroup)
	if err != nil {
		return err
	}

	productGroup.UID = dbProductGroup.ProductGroupUid

	return nil
}

func (c Catalog) GetProductGroup(ctx context.Context, sn string) (domain.ProductGroup, []*domain.Product, error) {
	dbProductGroup, err := c.models.ProductGroup.GetById(ctx, sn)
	if err != nil {
		return domain.ProductGroup{}, nil, err
	}

	dbInvLocs, err := c.models.InvLoc.GetByProductGroupId(ctx, sn)

	products := make([]*domain.Product, len(dbInvLocs))

	for i, dbInvLoc := range dbInvLocs {
		products[i] = &domain.Product{
			UID: dbInvLoc.InvMastUid, SN: dbInvLoc.InvMast.ItemId, Name: dbInvLoc.InvMast.ItemDesc,
			Description: dbInvLoc.InvMast.ExtendedDesc.String,
		}

	}

	return domain.ProductGroup{
		UID: dbProductGroup.ProductGroupUid, SN: dbProductGroup.ProductGroupId, Name: dbProductGroup.ProductGroupDesc,
	}, products, nil
}

func (c Catalog) UpdateProductGroup(ctx context.Context, productGroup *domain.ProductGroup) error {
	// select existing product group
	dbProductGroup, err := c.models.ProductGroup.GetById(ctx, productGroup.SN)
	if err != nil {
		return err
	}
	// update product group
	dbProductGroup.ProductGroupDesc = productGroup.Name

	// save updated product group
	err = c.models.ProductGroup.Update(ctx, dbProductGroup)
	if err != nil {
		return err
	}
	return nil
}

func (c Catalog) ListProductGroup(ctx context.Context) ([]*domain.ProductGroup, error) {
	dbProductGroups, err := c.models.ProductGroup.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	productGroups := make([]*domain.ProductGroup, len(dbProductGroups))
	for i, dbProductGroup := range dbProductGroups {
		productGroups[i] = &domain.ProductGroup{
			SN:   dbProductGroup.ProductGroupId,
			Name: dbProductGroup.ProductGroupDesc,
		}
	}
	return productGroups, nil
}
