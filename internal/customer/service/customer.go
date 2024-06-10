package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/customer/domain"
	"github.com/materials-resources/s-prophet/pkg/models"
)

type Customer struct {
	m models.Models
}

func NewCustomerService(models *models.Models) *Customer {
	return &Customer{
		m: *models,
	}
}

func (c Customer) GetContactById(ctx context.Context, id string) (domain.Contact, error) {

	// get contact record
	contactRec, err := c.m.Contacts.Get(ctx, id)
	if err != nil {
		return domain.Contact{}, err
	}

	// get ship to records for contact
	contactsXShipToRec, err := c.m.ContactsXShipTo.GetByContactId(ctx, "MRS", id)
	if err != nil {
		return domain.Contact{}, err

	}

	// map data
	dom := domain.Contact{
		Id:        contactRec.Id,
		FirstName: contactRec.FirstName,
		LastName:  contactRec.LastName,
		Phone:     contactRec.DirectPhone.String,
		Email:     contactRec.EmailAddress.String,
	}

	for _, r := range contactsXShipToRec {
		// get address info for customer branch
		addressRec, err := c.m.Address.Get(ctx, r.ShipToId)
		if err != nil {
			return domain.Contact{}, err

		}
		dom.Branches = append(
			dom.Branches, &domain.CustomerBranch{
				Id:   r.ShipToId,
				Name: addressRec.Name,
			})
	}
	return dom, nil
}
