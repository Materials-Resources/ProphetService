package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/materials-resources/s_prophet/config"
	pb "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"github.com/materials-resources/s_prophet/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.OrderServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	// Create a new grpc server
	baseServer := grpc.NewServer()

	config.NewConfig("../config.yml")

	// Register a new db instance
	db, _ := repository.NewDB(&config.Configuration)

	orderRepo := repository.NewOrderRepository(db)

	// Register the order server to the grpc server instance
	NewOrderServer(
		baseServer,
		orderRepo,
	)

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf(
				"error starting server %v",
				err,
			)
		}
	}()

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return lis.Dial()
			},
		),
	)
	if err != nil {
		log.Printf(
			"error connecting to server: %v",
			err,
		)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: ")
		}
		baseServer.Stop()
	}

	client := pb.NewOrderServiceClient(conn)

	return client, closer

}

func TestOrderService_GetOrder(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *pb.GetOrderResponse
		err error
	}

	tests := map[string]struct {
		in       *pb.GetOrderRequest
		expected expectation
	}{
		"Must_Success": {
			in: &pb.GetOrderRequest{Id: "1334044"},
			expected: expectation{
				out: &pb.GetOrderResponse{
					Id: "1334044",
					ShippingDetails: &pb.GetOrderResponse_ShippingDetails{
						Name:                 "",
						LineOne:              "",
						LineTwo:              "",
						City:                 "",
						State:                "",
						PostalCode:           "",
						Country:              "",
						DeliveryInstructions: "",
					},
					PurchaseOrder: "",
					ContactId:     "",
					Taker:         "",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(
			scenario,
			func(t *testing.T) {
				out, err := client.GetOrder(
					ctx,
					tt.in,
				)
				fmt.Println(out)
				if err != nil {
					if tt.expected.err.Error() != err.Error() {
						t.Errorf(
							"Err -> \nWant: %q\nGot: %q",
							tt.expected.out,
							out,
						)
					}
				} else {
					if tt.expected.out.Id != out.Id ||
						tt.expected.out.ShippingDetails != out.ShippingDetails {
						t.Errorf(
							"Out -> \nWant: %q\nGot: %q",
							tt.expected.out,
							out,
						)
					}
				}
			},
		)
	}
}
