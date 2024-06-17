package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/core/data"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/materials-resources/s-prophet/internal/validator"
	"github.com/materials-resources/s-prophet/pkg/consumer"
	"github.com/materials-resources/s-prophet/pkg/models"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"strconv"
)

type EventService struct {
	Producer *EventProducer
	Consumer *EventConsumer
}

var (
	ErrorResourceNotFound = errors.New("resource not found")
)

func NewCatalogService(
	models models.Models, tracer trace.Tracer, a *app.App) Catalog {

	return Catalog{
		localModel: data.NewModels(a.GetDB()),
		m:          models, tracer: tracer, app: a,
	}
}

type Catalog struct {
	app        *app.App
	m          models.Models
	tracer     trace.Tracer
	event      EventService
	localModel *data.Models
}

func (c Catalog) UpdateProduct(ctx context.Context, product *domain.Product, locations []float64) error {
	//TODO implement me
	panic("implement me")
}

func (c Catalog) ClerkUpdateProduct(ctx context.Context, product *domain.Product) error {
	v := validator.New()

	domain.ValidateProductUpdate(v, *product)

	if !v.Valid() {
		return errors.New(fmt.Sprintf("%s", v.Errors))

	}

	return c.localModel.InvLoc.Update(ctx, product)
}

func (c Catalog) ClerkDeleteProduct(ctx context.Context, uid string, deleteMode domain.DeleteMode) error {
	//TODO implement me
	panic("implement me")
}

func (c Catalog) ClerkListProducts(ctx context.Context, pagination domain.Pagination) ([]*domain.Product, *domain.PaginationMetadata, error) {
	opts := &data.InvLocListOptions{}
	return c.localModel.InvLoc.List(ctx, opts, pagination)
}

func (c Catalog) ClerkCreateGroup(ctx context.Context, productGroup *domain.ProductGroup) error {
	v := validator.New()

	domain.ValidateProductGroupCreate(v, *productGroup)

	if !v.Valid() {
		return errors.New(fmt.Sprintf("%s", v.Errors))
	}

	return c.localModel.ProductGroup.Create(ctx, productGroup)
}

func (c Catalog) ClerkListGroups(ctx context.Context) ([]*domain.ProductGroup, error) {
	opts := &data.ProductGroupListOptions{}
	return c.localModel.ProductGroup.List(ctx, opts)
}

func (c Catalog) ClerkReadGroup(ctx context.Context, sn string) (*domain.ProductGroup, error) {
	return c.localModel.ProductGroup.Get(ctx, sn)

}

func (c Catalog) ClerkUpdateGroup(ctx context.Context, productGroup *domain.ProductGroup) error {

	v := validator.New()
	domain.ValidateProductGroupUpdate(v, *productGroup)

	valid := v.Valid()
	if !valid {
		return errors.New(fmt.Sprintf("%s", v.Errors))

	}

	return c.localModel.ProductGroup.Update(ctx, productGroup)

}

func (c Catalog) ClerkGetProduct(ctx context.Context, uid string) (*domain.Product, error) {
	return c.getProduct(ctx, uid)
}

func (c Catalog) CustomerGetProduct(ctx context.Context, uid string) (*domain.Product, error) {
	return c.getProduct(ctx, uid)
}

func (c Catalog) getProduct(ctx context.Context, uid string) (*domain.Product, error) {

	product, err := c.localModel.InvLoc.Get(ctx, uid)
	if err != nil {
		return nil, err

	}
	return product, nil
}

func (c Catalog) SetPrimaryProductSupplier(
	ctx context.Context, productUid int32, locationId, supplierUid, divisionId float64) error {

	err := c.m.ExecTx(
		ctx, func(m *models.Models) error {

			inventorySupplier, err := m.InventorySupplier.GetBySupplierIdDivisionIdInvMastUid(
				ctx, supplierUid,
				divisionId, productUid)
			if err != nil {
				return err

			}
			inventorySupplierXLocs, err := m.InventorySupplierXLoc.GetByInvMastUidAndLocationId(
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
				if err := m.InventorySupplierXLoc.Update(ctx, inventorySupplierXLoc); err != nil {
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

func (c Catalog) RequestDeleteProduct(ctx context.Context, uid string) error {

	err := c.event.Producer.PublishDeleteProduct(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (c Catalog) DeleteProduct(ctx context.Context, uid string) error {

	ctx, span := c.tracer.Start(ctx, "service.catalog.DeleteProduct")
	defer span.End()
	//
	//err = c.m.ExecTx(
	//
	//	ctx, func(m *models.Models) error {
	//		// get invMast record to see if it exists
	//		invMast, err := m.InvMast.Get(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//
	//		// check if product has receipts
	//		receiptCount, err := m.InventoryReceiptsLine.CountByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		if receiptCount > 0 {
	//			return errors.New("cannot delete product with receipts")
	//		}
	//
	//		// check if product has orders
	//		orderCount, err := m.OeLine.CountByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		if orderCount > 0 {
	//			return errors.New("cannot delete product with orders")
	//		}
	//		invBins, err := m.InvBin.GetByInvMastUid(ctx, int32(invMastUid))
	//		for _, invBin := range invBins {
	//			err = m.InvBin.Delete(ctx, invBin)
	//			if err != nil {
	//				return err
	//			}
	//
	//		}
	//
	//		// query and delete alternateCode records belonging to the product
	//		alternateCodes, err := m.AlternateCode.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, alternateCode := range alternateCodes {
	//			err = m.AlternateCode.Delete(ctx, alternateCode)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete assemblyHdr and assemblyLine records belonging to the product
	//		assemblyHdr, err := m.AssemblyHdr.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		if assemblyHdr != nil {
	//			assemblyLine, err := m.AssemblyLine.GetByInvMastUid(ctx, assemblyHdr.InvMastUid)
	//			if err != nil {
	//				return err
	//			}
	//			for _, line := range assemblyLine {
	//				err = m.AssemblyLine.Delete(ctx, line)
	//				if err != nil {
	//					return err
	//				}
	//			}
	//			err = m.AssemblyHdr.Delete(ctx, assemblyHdr)
	//		}
	//
	//		// query and delete pricePage and pricePageXBook records belonging to the product
	//		pricePages, err := m.PricePage.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, pricePage := range pricePages {
	//			pricePageXBook, err := m.PricePageXBook.GetByPricePageUid(ctx, pricePage.PricePageUid)
	//			if err != nil {
	//				return err
	//			}
	//			err = m.PricePageXBook.Delete(ctx, pricePageXBook)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete itemConversion records belonging to the product
	//		itemConversions, err := m.ItemConversion.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, conversion := range itemConversions {
	//			err = m.ItemConversion.Delete(ctx, conversion)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete invAdjLine records belonging to the product
	//		invAdjLines, err := m.InvAdjLine.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, invAdjLine := range invAdjLines {
	//			err = m.InvAdjLine.Delete(ctx, invAdjLine)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete invTran records belonging to the product
	//		invTrans, err := m.InvTran.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, invTran := range invTrans {
	//			err = m.InvTran.Delete(ctx, invTran)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete itemCategoryXInvMast records belonging to the product
	//		itemCategoryXInvMasts, err := m.ItemCategoryXInvMast.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, itemCategoryXInvMast := range itemCategoryXInvMasts {
	//			err = m.ItemCategoryXInvMast.Delete(ctx, itemCategoryXInvMast)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete invLocMsp records belonging to the product
	//		invLocMsps, err := m.InvLocMsp.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, invLocMsp := range invLocMsps {
	//			err = m.InvLocMsp.Delete(ctx, invLocMsp)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete invLocStockStatus records belonging to the product
	//		invLocStockStatus, err := m.InvLocStockStatus.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, stockStatus := range invLocStockStatus {
	//			if err := m.InvLocStockStatus.Delete(ctx, stockStatus); err != nil {
	//				return err
	//			}
	//		}
	//
	//		itemIdChangeHistories, err := m.ItemIdChangeHistory.GetByInvMastUid(ctx, int32(invMastUid))
	//		if err != nil {
	//			return err
	//		}
	//		for _, itemIdChangeHistory := range itemIdChangeHistories {
	//			if err := m.ItemIdChangeHistory.Delete(ctx, itemIdChangeHistory); err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete itemUom records belonging to the product
	//		itemUoms, err := m.ItemUom.GetByInvMastUid(ctx, int32(invMastUid))
	//		for _, itemUom := range itemUoms {
	//			inventorySupplierXUoms, err := m.InventorySupplierXUom.GetByItemUomUid(
	//				ctx,
	//				itemUom.ItemUomUid)
	//			if err != nil {
	//				return err
	//			}
	//			for _, inventorySupplierXUom := range inventorySupplierXUoms {
	//				if err := m.InventorySupplierXUom.Delete(ctx, inventorySupplierXUom); err != nil {
	//					return err
	//				}
	//			}
	//			if err := m.ItemUom.Delete(ctx, itemUom); err != nil {
	//				return err
	//			}
	//		}
	//
	//		// query and delete inventorySupplier records belonging to the product
	//		inventorySuppliers, err := m.InventorySupplier.GetByInvMastUid(ctx, int32(invMastUid))
	//		for _, inventorySupplier := range inventorySuppliers {
	//
	//			inventorySupplierXLocs, err := m.InventorySupplierXLoc.GetByInventorySupplierUid(
	//				ctx, inventorySupplier.InventorySupplierUid)
	//			if err != nil {
	//				return err
	//
	//			}
	//			for _, inventorySupplierXLoc := range inventorySupplierXLocs {
	//				if err := m.InventorySupplierXLoc.Delete(ctx, inventorySupplierXLoc); err != nil {
	//					return err
	//				}
	//
	//			}
	//			if err := m.InventorySupplier.Delete(ctx, inventorySupplier); err != nil {
	//				return err
	//			}
	//
	//		}
	//
	//		//// query and delete invLoc records belonging to the product
	//		//invLocs, err := m.InvLoc.GetByInvMastUid(ctx, []float64{1001}, int32(invMastUid))
	//		//if err != nil {
	//		//	return err
	//		//}
	//		//for _, invLoc := range invLocs {
	//		//	if err := m.InvLoc.Delete(ctx, invLoc); err != nil {
	//		//		return err
	//		//	}
	//		//}
	//
	//		// query and delete averageInventoryValue records belonging to the product
	//		averageInventoryValues, err := m.AverageInventoryValue.GetByInvMastUid(ctx, int32(invMastUid))
	//		for _, averageInventoryValue := range averageInventoryValues {
	//			if err := m.AverageInventoryValue.Delete(ctx, averageInventoryValue); err != nil {
	//				return err
	//			}
	//
	//		}
	//
	//		// delete invMast records belonging to the product
	//		if err := m.InvMast.Delete(ctx, invMast); err != nil {
	//			return err
	//		}

	err := c.localModel.ModelsTx(ctx, func(m *data.Models) error {
		invMastUid, err := strconv.Atoi(uid)
		if err != nil {
			return err
		}

		err = m.InvMast.DeleteByInvMastUid(ctx, int32(invMastUid))
		if err != nil {
			fmt.Println(err)
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

func (c Catalog) GetProductBySupplierPartNumber(
	ctx context.Context, partNumber string,
	supplier float64) (domain.Product, error) {
	invMast, err := c.m.InvMast.GetBySupplierPartNumber(ctx, supplier, partNumber)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		Uid: strconv.Itoa(int(invMast.InvMastUid)), Sn: &invMast.ItemId, Name: &invMast.ItemDesc,
	}, nil
}
