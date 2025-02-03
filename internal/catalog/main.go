package catalog

import (
	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/api"
	"github.com/materials-resources/s-prophet/internal/catalog/service"
	"github.com/materials-resources/s-prophet/internal/grpc/catalog/catalogconnect"
	"github.com/materials-resources/s-prophet/pkg/kafka"
	"github.com/materials-resources/s-prophet/pkg/sr"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	app.OnStart(
		"catalog.start", func(ctx context.Context, a *app.App) error {

			//manager, err := event.NewManager(a)
			//if err != nil {
			//	return err
			//}
			//
			//err = manager.RegisterSchemasWithSerde()
			//if err != nil {
			//	return err
			//}

			//producer := produce.NewProducer(a, manager)

			factory := kafka.NewFactory(a.Config().Kafka, a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider())

			schemaFactory := sr.NewFactory(a.Config().SchemaRegistry)

			srClient, err := schemaFactory.NewSchemaRegistryClient()
			if err != nil {
				return err
			}

			svc := service.NewCatalogService(a, factory, srClient)

			err = svc.Initialize(context.Background())
			if err != nil {
				return fmt.Errorf("could not initialize catalog service %v", err)
			}

			otelInterceptor, err := newInterceptor(a.Otel.GetTracerProvider(), a.Otel.GetMeterProvider(), a.Otel.GetTextMapPropagator())
			if err != nil {
				return fmt.Errorf("could not create otel interceptor %v", err)
			}

			path, handler := catalogconnect.NewCatalogServiceHandler(api.NewCatalogServiceHandler(&svc), connect.WithInterceptors(otelInterceptor))

			a.RegisterService(path, handler, catalogconnect.CatalogServiceName)

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
