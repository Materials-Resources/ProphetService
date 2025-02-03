package customer

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/customer/api"
	"github.com/materials-resources/s-prophet/internal/customer/repository"
	"github.com/materials-resources/s-prophet/internal/customer/service"
	"github.com/materials-resources/s-prophet/internal/grpc/customer/customerconnect"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	app.OnStart(
		"customer.start", func(ctx context.Context, a *app.App) error {
			svc := service.NewCustomerService(repository.NewRepository(a.GetDB()))

			otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
			if err != nil {
				return fmt.Errorf("could not create otel interceptor %v", err)
			}

			path, handler := customerconnect.NewCustomerServiceHandler(api.NewCustomerServiceHandler(svc),
				connect.WithInterceptors(otelInterceptor))

			a.RegisterService(path, handler, customerconnect.CustomerServiceName)
			return nil
		},
	)
}

func newInterceptor(tp trace.TracerProvider, mp metric.MeterProvider, p propagation.TextMapPropagator) (connect.Interceptor, error) {
	return otelconnect.NewInterceptor(
		otelconnect.WithTracerProvider(tp),
		otelconnect.WithMeterProvider(mp),
		otelconnect.WithPropagator(p),
	)
}
