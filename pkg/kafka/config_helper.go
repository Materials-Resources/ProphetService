package kafka

import (
	"github.com/materials-resources/s-prophet/config"
	"github.com/twmb/franz-go/plugin/kotel"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

// NewKgoConfig creates a new Config for the Kafka Client as exposed by the franz-go library.
// If TLS certificates can't be read an error will be returned.
func NewKgoConfig(cfg *config.Kafka, kotel *kotel.Kotel) ([]kgo.Opt, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.RecordDeliveryTimeout(10 * time.Second),
		kgo.WithHooks(kotel.Hooks()...),
	}

	//// Configure SASL
	//if cfg.SASL.Enabled {
	//	// SASL Plain
	//	if cfg.SASL.Mechanism == "PLAIN" {
	//		mechanism := plain.Auth{
	//			User: cfg.SASL.Username,
	//			Pass: cfg.SASL.Password,
	//		}.AsMechanism()
	//		opts = append(opts, kgo.SASL(mechanism))
	//	}
	//
	//	// SASL SCRAM
	//	if cfg.SASL.Mechanism == "SCRAM-SHA-256" || cfg.SASL.Mechanism == "SCRAM-SHA-512" {
	//		var mechanism sasl.Mechanism
	//		scramAuth := scram.Auth{
	//			User: cfg.SASL.Username,
	//			Pass: cfg.SASL.Password,
	//		}
	//		if cfg.SASL.Mechanism == "SCRAM-SHA-256" {
	//			mechanism = scramAuth.AsSha256Mechanism()
	//		}
	//		if cfg.SASL.Mechanism == "SCRAM-SHA-512" {
	//			mechanism = scramAuth.AsSha512Mechanism()
	//		}
	//		opts = append(opts, kgo.SASL(mechanism))
	//	}
	//}
	//
	//// Configure TLS
	//if cfg.TLS.Enabled {
	//	tlsCfg, err := cfg.TLS.TLSConfig()
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to create tls config: %w", err)
	//	}
	//
	//	opts = append(opts, kgo.DialTLSConfig(tlsCfg))
	//}

	return opts, nil
}
