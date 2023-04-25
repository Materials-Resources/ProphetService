package grpc_service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	pb "materials-resources.com/msv/prophet/proto/order/v1alpha0"
	"net"
	"testing"
)

func dialer() func(ctx context.Context, string2 string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterOrderServiceServer(server, &OrderServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
func TestOrderServer_GetOrder(t *testing.T) {

	//tests := []struct {
	//	name string
	//	req  *pb.GetOrderRequest
	//	res  *pb.GetOrderResponse
	//}{
	//	{
	//		"valid request with details of order",
	//		&pb.GetOrderRequest{Id: ""},
	//		&pb.GetOrderResponse{},
	//	},
	//}
}
