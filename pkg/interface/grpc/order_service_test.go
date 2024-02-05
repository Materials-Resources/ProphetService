package grpc

import (
	"context"
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/mocks/repomocks"
	pb "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context, repo *repomocks.OrderRepository) (pb.OrderServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	NewOrderService(baseServer, repo)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error starting server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		},
		), grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listner: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewOrderServiceClient(conn)

	return client, closer

}

func Test_orderService_GetOrder(t *testing.T) {
	ctx := context.Background()

	repo := &repomocks.OrderRepository{}

	repo.On("FindOrderById", "3").Return(&entities.ValidatedOrder{
		Order: entities.Order{
			ID:              "3",
			ShippingAddress: entities.ValidatedAddress{},
			CustomerContact: struct {
				Id          string
				Name        string
				PhoneNumber string
				Email       string
			}{},
			Items: []entities.ValidatedOrderItem{{OrderItem: entities.OrderItem{
				ID:              "1",
				SN:              "1",
				Name:            "Product 1",
				QuantityOrdered: 2,
			},
			},
			},
			PurchaseOrder: "TESTNO",
		},
	}, nil,
	)

	client, closer := server(ctx, repo)
	defer closer()

	type expectation struct {
		want    *pb.GetOrderResponse
		wantErr bool
	}
	tests := []struct {
		name     string
		request  *pb.GetOrderRequest
		expected expectation
	}{
		{
			name:    "Valid_Order_Id",
			request: &pb.GetOrderRequest{Id: "3"},
			expected: expectation{
				want: &pb.GetOrderResponse{
					Id: "3",
					ShippingAddress: &pb.Address{
						Name:       "",
						LineOne:    "",
						LineTwo:    "",
						City:       "",
						State:      "",
						PostalCode: "",
						Country:    "",
					}, CustomerContact: &pb.CustomerContact{
						Id:          "",
						Name:        "",
						PhoneNumber: "",
						Email:       "",
						Title:       "",
					},
					PurchaseOrder: "TESTNO",
				},
				wantErr: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetOrder(ctx, tt.request)
			if (err != nil) != tt.expected.wantErr {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.expected.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected.want) {
				t.Errorf("GetOrder() got = %v, want %v", got, tt.expected.want)
			}
		},
		)
	}
}

func Test_orderService_GetPickTicketById(t *testing.T) {
	ctx := context.Background()

	repo := &repomocks.OrderRepository{}

	repo.On("FindOrderById", "3").Return(&entities.ValidatedOrder{
		Order: entities.Order{
			ID:              "3",
			ShippingAddress: entities.ValidatedAddress{},
			CustomerContact: struct {
				Id          string
				Name        string
				PhoneNumber string
				Email       string
			}{},
			Items:         nil,
			PurchaseOrder: "TESTNO",
		},
	}, nil,
	)

	client, closer := server(ctx, repo)
	defer closer()

	type expectation struct {
		want *pb.GetPickTicketByIdResponse
		err  bool
	}
	tests := []struct {
		name     string
		request  *pb.GetPickTicketByIdRequest
		expected expectation
	}{
		{
			name:    "Valid_Pick_Ticket_Id",
			request: &pb.GetPickTicketByIdRequest{Id: "500001"},
			expected: expectation{
				want: &pb.GetPickTicketByIdResponse{
					Id:                   "",
					OrderId:              "",
					OrderPurchaseOrder:   "",
					ShippingAddress:      nil,
					OrderCustomerContact: nil,
				},
				err: false,
			},
		},
		{
			name: "Invalid_Pick_Ticket_Id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetPickTicketById(ctx, tt.request)
			if (err != nil) != tt.expected.err {
				t.Errorf("GetPickTicketById() error = %v, wantErr %v", err, tt.expected.err)
				return
			}
			if !reflect.DeepEqual(got, tt.expected.want) {
				t.Errorf("GetPickTicketById() got = %v, want %v", got, tt.expected.want)
			}
		},
		)
	}
}
