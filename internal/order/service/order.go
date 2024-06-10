package service

//
//func (s *OrderService) CreateQuote(ctx context.Context, order *domain.Order) error {
//	contactRec, err := s.models.Contacts.Get(ctx, order.ContactId)
//	if err != nil {
//		return err
//	}
//
//	_, err = s.models.ContactsXShipTo.GetByContactId(ctx, "MRS", order.ContactId)
//	if err != nil {
//		return err
//	}
//
//	contactExistsOnShipTo := false
//	// check to see if order.ShippingAddress.Id is in the slice of contactsXShipToRecs
//	//for _, rec := range contactsXShipToRecs {
//	//	if rec.ShipToId == order.CustomerBranchId {
//	//		contactExistsOnShipTo = true
//	//		break
//	//	}
//	//}
//
//	if !contactExistsOnShipTo {
//		return errors.New("provided shipping address is not associated with contact")
//
//	}
//
//	shipToRec, err := s.models.ShipTo.Get(ctx, "MRS", 0)
//	if err != nil {
//		return err
//	}
//
//	customerRec, err := s.models.Customer.Get(ctx, shipToRec.CustomerId, "MRS")
//	if err != nil {
//		return err
//	}
//
//	addressRec, err := s.models.Address.Get(ctx, shipToRec.ShipToId)
//	if err != nil {
//		return err
//	}
//
//	shipToSalesrepRecs, err := s.models.ShipToSalesrep.GetByShipToId(ctx, "MRS", shipToRec.ShipToId, true)
//	if err != nil {
//		return err
//	}
//
//	// create the order
//	oeHdrParams := data.CreateOeHdrParams{
//		CustomerId:           customerRec.CustomerId,
//		AddressId:            addressRec.Id,
//		ContactId:            contactRec.Id,
//		Ship2Name:            addressRec.Name,
//		Ship2Add1:            addressRec.MailAddress1.String,
//		Ship2Add2:            addressRec.MailAddress2.String,
//		Ship2City:            addressRec.MailCity.String,
//		Ship2State:           addressRec.MailState.String,
//		Ship2Zip:             addressRec.MailPostalCode.String,
//		Ship2Country:         addressRec.MailCountry.String,
//		ShipToPhone:          contactRec.DirectPhone.String,
//		PoNo:                 order.PurchaseOrder,
//		ProjectedOrder:       "Y",
//		DeliveryInstructions: addressRec.DeliveryInstructions1.String,
//		PackingBasis:         addressRec.ShipToPackingBasis.String,
//		PickTicketType:       customerRec.PickTicketType.String,
//		Terms:                customerRec.TermsId,
//	}
//
//	oeHdrParams.CarrierId, err = strconv.ParseFloat(addressRec.CarrierId.String, 64)
//	if err != nil {
//		return err
//	}
//	oeHdr, err := s.models.OeHdr.Create(ctx, oeHdrParams)
//	if err != nil {
//		return err
//	}
//
//	// add the salesrep to the order
//	_, err = s.models.OeHdrSalesrep.Create(
//		ctx, data.OeHdrSalesrepParams{
//			OrderNumber:     oeHdr.OrderNo,
//			SalesrepId:      shipToSalesrepRecs[0].SalesrepId,
//			CommissionSplit: shipToSalesrepRecs[0].CommissionPercentage,
//			PrimarySalesrep: shipToSalesrepRecs[0].PrimarySalesrep.String,
//		})
//
//	// create the quote record
//	_, err = s.models.QuoteHdr.Create(
//		ctx, data.CreateQuoteHdrParams{
//			OeHdrUid: oeHdr.OeHdrUid,
//		})
//
//	for _, item := range order.Items {
//		invMastUid, err := strconv.Atoi(item.ProductUid)
//		if err != nil {
//			return err
//		}
//		invLoc, err := s.models.InvLoc.Get(ctx, 1001, "MRS", int32(invMastUid))
//		if err != nil {
//			switch {
//			case errors.Is(err, data.ErrRecordNotFound):
//				return fmt.Errorf("product with uid %d not found", item.ProductUid)
//			default:
//				return err
//			}
//
//		}
//		inventorySupplier, err := s.models.InventorySupplier.GetPrimarySupplierByLocation(
//			ctx, invLoc.LocationId,
//			invLoc.InvMastUid)
//		if err != nil {
//			return err
//		}
//
//		oeLineRec, err := s.models.OeLine.Create(
//			ctx, data.CreateOeLineParams{
//				OrderNo:              oeHdr.OrderNo,
//				UnitPrice:            0,
//				QtyOrdered:           item.OrderQuantity,
//				BaseUtPrice:          0,
//				CalcValue:            0,
//				SourceLocId:          1001,
//				ShipLocId:            1001,
//				SupplierId:           inventorySupplier.SupplierId,
//				ProductGroupId:       invLoc.ProductGroupId.String,
//				ExtendedDesc:         invLoc.InvMast.ExtendedDesc.String,
//				CustomerPartNumber:   invLoc.InvMast.ItemId,
//				CommissionCost:       0,
//				OeHdrUid:             oeHdr.OeHdrUid,
//				InvMastUid:           invLoc.InvMastUid,
//				SalesDiscountGroupId: invLoc.SalesDiscountGroup.String,
//				UnitQuantity:         item.OrderQuantity,
//				UnitSize:             1,
//			})
//		if err != nil {
//			return err
//		}
//		_, err = s.models.QuoteLine.Create(
//			ctx, data.CreateQuoteLineParams{
//				OeLineUid:    oeLineRec.OeLineUid,
//				QtyConverted: 0,
//				UomConverted: oeLineRec.UnitOfMeasure.String,
//			})
//		if err != nil {
//			return err
//		}
//
//	}
//	order.Id = oeHdr.OrderNo
//	return err
//}
