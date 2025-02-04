package service

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/internal/customer/repository"
	customerv1 "github.com/materials-resources/s-prophet/internal/grpc/customer"
	"github.com/materials-resources/s-prophet/pkg/helpers"
)

type CustomerService interface {
	GetBranchesForContact(ctx context.Context, req *customerv1.GetBranchesForContactRequest) (*customerv1.GetBranchesForContactResponse, error)
	GetBranch(ctx context.Context, req *customerv1.GetBranchRequest) (*customerv1.GetBranchResponse, error)
	GetRecentPurchasesByBranch(ctx context.Context, req *customerv1.GetRecentPurchasesByBranchRequest) (*customerv1.GetRecentPurchasesByBranchResponse, error)
}
type Customer struct {
	repository *repository.Repository
}

func (c Customer) GetRecentPurchasesByBranch(ctx context.Context, req *customerv1.GetRecentPurchasesByBranchRequest) (*customerv1.GetRecentPurchasesByBranchResponse, error) {
	purchases, err := c.repository.Invoice.GetRecentPurchasesByBranch(ctx, req.GetId(), req.GetLimit())
	if err != nil {
		return &customerv1.GetRecentPurchasesByBranchResponse{}, err
	}
	response := &customerv1.GetRecentPurchasesByBranchResponse{}

	for _, purchase := range purchases {
		response.Items = append(response.Items, &customerv1.GetRecentPurchasesByBranchResponse_Item{
			ProductId:          purchase.ProductId,
			ProductSn:          helpers.GetOptionalValue(purchase.ProductSn, ""),
			ProductName:        helpers.GetOptionalValue(purchase.ProductName, ""),
			ProductDescription: helpers.GetOptionalValue(purchase.ProductDescription, ""),
			OrderedQuantity:    helpers.GetOptionalValue(purchase.QuantityPurchased, 0),
			UnitType:           helpers.GetOptionalValue(purchase.UnitType, ""),
		})
	}
	return response, nil
}

func (c Customer) GetBranch(ctx context.Context, req *customerv1.GetBranchRequest) (*customerv1.GetBranchResponse, error) {
	branch, err := c.repository.Branch.GetBranchById(ctx, req.Id)

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			return nil, connect.NewError(connect.CodeNotFound, err)
		default:
			return nil, err
		}
	}

	response := &customerv1.GetBranchResponse{
		Branch: &customerv1.Branch{
			Id:   branch.Id,
			Name: branch.Name,
		},
	}

	return response, nil
}

func (c Customer) GetBranchesForContact(ctx context.Context, req *customerv1.GetBranchesForContactRequest) (*customerv1.GetBranchesForContactResponse, error) {
	branches, err := c.repository.Branch.GetBranchesByContact(ctx, req.ContactId)

	if err != nil {
		return &customerv1.GetBranchesForContactResponse{}, err
	}

	response := &customerv1.GetBranchesForContactResponse{}

	for _, branch := range branches {
		response.Branches = append(response.Branches, &customerv1.BranchSummary{
			Id:   branch.Id,
			Name: branch.Name,
		})
	}

	return response, nil
}

func NewCustomerService(repository *repository.Repository) *Customer {
	return &Customer{
		repository: repository,
	}
}

var _ CustomerService = (*Customer)(nil)
