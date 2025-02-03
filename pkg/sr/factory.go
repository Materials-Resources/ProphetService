package sr

import (
	"fmt"
	"github.com/materials-resources/s-prophet/config"
	"github.com/twmb/franz-go/pkg/sr"
)

// Factory allows requesters to create new preconfigured Schema Registry clients
// based on the given configuration.
type Factory struct {
	Config config.SchemaRegistry
}

// NewFactory creates a new SchemaRegistry Factory.
func NewFactory(cfg config.SchemaRegistry) *Factory {
	return &Factory{
		Config: cfg,
	}
}

// NewSchemaRegistryClient creates a new SchemaRegistry client with the same stored
// SchemaRegistry configuration. If SchemaRegistry is not configured, this will return
// a nil client without error.
func (s *Factory) NewSchemaRegistryClient(additionalOpts ...sr.ClientOpt) (*sr.Client, error) {
	if s.Config.Address == "" {
		return nil, nil
	}

	opts := []sr.ClientOpt{
		sr.URLs(s.Config.Address),
	}
	if s.Config.BasicAuth.Username != "" {
		opts = append(opts, sr.BasicAuth(s.Config.BasicAuth.Username, s.Config.BasicAuth.Password))
	}

	if s.Config.TLS.Enabled {
		tlsCfg, err := s.Config.TLS.TLSConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load tls config: %w", err)
		}
		opts = append(opts, sr.DialTLSConfig(tlsCfg))
	}

	opts = append(opts, additionalOpts...)

	return sr.NewClient(opts...)
}
