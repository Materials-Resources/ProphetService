package api

import (
	billingpb "github.com/materials-resources/s-prophet/proto/billing/v1"
	"google.golang.org/grpc"
)

func RegisterRoutes(s *grpc.Server, routes billingpb.BillingServiceServer) {
	billingpb.RegisterBillingServiceServer(s, routes)
}
