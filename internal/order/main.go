package order

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/grpc/order/orderconnect"
	"github.com/materials-resources/s-prophet/internal/order/api"
	"github.com/materials-resources/s-prophet/internal/order/repository"
	"github.com/materials-resources/s-prophet/internal/order/service"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	app.OnStart(
		"order.start", func(ctx context.Context, a *app.App) error {
			otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
			if err != nil {
				return fmt.Errorf("could not create otel interceptor %v", err)
			}

			path, handler := orderconnect.NewOrderServiceHandler(api.NewOrderServiceHandler(service.NewOrderService(repository.NewRepository(a.GetDB()))), connect.WithInterceptors(otelInterceptor))
			a.RegisterService(path, handler, orderconnect.OrderServiceName)
			return nil
		},
	)
}

func newInterceptor(tp trace.TracerProvider, mp metric.MeterProvider, p propagation.TextMapPropagator) (connect.Interceptor, error) {
	return otelconnect.NewInterceptor(
		otelconnect.WithTrustRemote(),
		otelconnect.WithTracerProvider(tp),
		otelconnect.WithMeterProvider(mp),
		otelconnect.WithPropagator(p),
	)
}
