package prophet_19_1_3668

import (
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
)

func FromDBPickTicket(dbPickTicket *model.OePickTicket) (*entities.ValidatedPickTicket, error) {
	var shippingAddress = &entities.Address{
		Name:       dbPickTicket.OeHdr.Ship2Name.String,
		LineOne:    dbPickTicket.OeHdr.Ship2Add1.String,
		LineTwo:    dbPickTicket.OeHdr.Ship2Add2.String,
		City:       dbPickTicket.OeHdr.Ship2City.String,
		State:      dbPickTicket.OeHdr.Ship2State.String,
		PostalCode: dbPickTicket.OeHdr.Ship2Zip.String,
	}
	validatedShippingAddress, err := entities.NewValidatedAddress(shippingAddress)
	if err != nil {
		return nil, err
	}
	var e = &entities.PickTicket{
		ID:                 fmt.Sprintf("%.0f", dbPickTicket.PickTicketNo),
		ShippingAddress:    *validatedShippingAddress,
		OrderId:            dbPickTicket.OrderNo,
		OrderPurchaseOrder: dbPickTicket.OeHdr.PoNo.String,
		OrderContactName:   "",
	}
	return entities.NewValidatedPickTicket(e)
}

func FromDBOrder(dbOrder *model.OeHdr) (*entities.ValidatedOrder, error) {

	var shippingAddress = &entities.Address{
		Name:       dbOrder.Ship2Name.String,
		LineOne:    dbOrder.Ship2Add1.String,
		LineTwo:    dbOrder.Ship2Add2.String,
		City:       dbOrder.Ship2City.String,
		State:      dbOrder.Ship2State.String,
		PostalCode: dbOrder.Ship2Zip.String,
	}

	validatedShippingAddress, err := entities.NewValidatedAddress(shippingAddress)
	if err != nil {
		return nil, err
	}

	var o = &entities.Order{
		ShippingAddress: *validatedShippingAddress,
		PurchaseOrder:   dbOrder.PoNo.String,
	}
	o.CustomerContact.Id = dbOrder.ContactId.String
	o.CustomerContact.Name = fmt.Sprintf("%s %s", dbOrder.Contact.FirstName, dbOrder.Contact.LastName)
	o.CustomerContact.Email = dbOrder.Contact.EmailAddress.String
	o.CustomerContact.PhoneNumber = dbOrder.Contact.DirectPhone.String

	for _, item := range dbOrder.OeLineItems {
		var orderItem = &entities.OrderItem{
			ID:              strconv.Itoa(int(item.InvMastUid)),
			SN:              item.InvMast.ItemId,
			Name:            item.InvMast.ItemDesc,
			QuantityOrdered: item.UnitQuantity,
		}
		validatedOrderItem, err := entities.NewValidatedOrderItem(orderItem)
		if err != nil {
			//TODO handle error

			fmt.Println(err)
		}
		o.Items = append(o.Items, *validatedOrderItem)

	}

	o.ID = dbOrder.OrderNo
	return entities.NewValidatedOrder(o)
}
