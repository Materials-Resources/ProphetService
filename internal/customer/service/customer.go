package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/customer/domain"
	"github.com/materials-resources/s-prophet/internal/customer/repository"
	"github.com/materials-resources/s-prophet/pkg/helpers"
)

type Customer struct {
	repository *repository.Repository
}

func NewCustomerService(repository *repository.Repository) *Customer {
	return &Customer{
		repository: repository,
	}
}

func (c Customer) GetContactById(ctx context.Context, id string) (domain.Contact, error) {

	// get contact record
	contactRec, err := c.repository.Contacts.Select(ctx, id)
	if err != nil {
		return domain.Contact{}, err
	}

	// map data
	dom := domain.Contact{
		Id:        contactRec.Id,
		FirstName: contactRec.FirstName,
		LastName:  contactRec.LastName,
		Phone:     helpers.GetValueOrDefault(contactRec.DirectPhone, ""),
		Email:     helpers.GetValueOrDefault(contactRec.EmailAddress, ""),
	}

	for _, r := range contactRec.ContactsXShipTo {
		dom.Branches = append(
			dom.Branches, &domain.CustomerBranch{
				Id:   r.ShipToId,
				Name: r.Address.Name,
			})
	}
	return dom, nil
}
