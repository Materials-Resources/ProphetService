package event

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"github.com/twmb/franz-go/plugin/kotel"
	"go.opentelemetry.io/otel/propagation"
	"sync"
)

const (
	UpdateGroupTopic   = "catalog_group_update"
	CreateGroupTopic   = "catalog_group_create"
	UpdateProductTopic = "product_update"
	DeleteProductTopic = "product_delete"
)

type Manager struct {
	app   *app.App
	Serde *sr.Serde

	tracer     *kotel.Tracer
	tracerOnce sync.Once
}

func NewManager(app *app.App) (*Manager, error) {
	manager := &Manager{
		Serde: &sr.Serde{},
		app:   app,
	}
	registry := NewRegistry()

	err := registry.RegisterSchemas(manager.Serde)
	if err != nil {
		return nil, err
	}
	return manager, nil
}

func (m *Manager) GetDefaultKgoOptions() []kgo.Opt {
	return []kgo.Opt{
		kgo.SeedBrokers("localhost:19092"),
		kgo.WithHooks(m.createKotelService().Hooks()),
	}
}

func (m *Manager) createKotelService() *kotel.Kotel {
	kotelOpts := []kotel.Opt{
		kotel.WithTracer(m.GetKotelTracer()),
	}
	return kotel.NewKotel(kotelOpts...)
}

func (m *Manager) GetKotelTracer() *kotel.Tracer {

	m.tracerOnce.Do(func() {
		tracerOpts := []kotel.TracerOpt{
			kotel.TracerProvider(m.app.GetTP()),
			kotel.TracerPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{})),
		}

		m.tracer = kotel.NewTracer(tracerOpts...)
	})

	return m.tracer
}
