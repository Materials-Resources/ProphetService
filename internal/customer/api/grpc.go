package api

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/customer/service"
	rpc "github.com/materials-resources/s_prophet/proto/customer/v1"
)

type CustomerServer struct {
	sv service.Service
}

func NewCustomerServer(sv *service.Customer) CustomerServer {
	return CustomerServer{
		sv: *sv,
	}
}

func (c CustomerServer) GetOrders(ctx context.Context, request *rpc.GetOrdersRequest) (*rpc.GetOrdersResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c CustomerServer) GetCustomer(ctx context.Context, request *rpc.GetCustomerRequest) (
	*rpc.GetCustomerResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c CustomerServer) GetContactById(
	ctx context.Context, request *rpc.GetContactByIdRequest) (*rpc.GetContactByIdResponse, error) {
	contact, err := c.sv.GetContactById(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	res := &rpc.GetContactByIdResponse{}

	res.Contact = &rpc.GetContactByIdResponse_Contact{
		Id:        contact.Id,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}

	res.Branches = make([]*rpc.CustomerBranch, len(contact.Branches))
	for i, b := range contact.Branches {
		res.Branches[i] = &rpc.CustomerBranch{
			Id:   b.Id,
			Name: b.Name,
		}
	}

	return res, nil

}
