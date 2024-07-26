package app

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (a *App) newGrpcServer() *grpc.Server {
	handler := grpc.StatsHandler(otelgrpc.NewServerHandler(otelgrpc.WithTracerProvider(a.traceProvider)))

	s := grpc.NewServer(handler)
	reflection.Register(s)
	return s
}
