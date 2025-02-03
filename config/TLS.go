package config

import "crypto/tls"
import "github.com/twmb/tlscfg"

// TLS configuration for either the HTTP protocol or Kafka API.
type TLS struct {
	Enabled               bool   `koanf:"enabled"`
	CaFilepath            string `koanf:"caFilepath"`
	CertFilepath          string `koanf:"certFilepath"`
	KeyFilepath           string `koanf:"keyFilepath"`
	InsecureSkipTLSVerify bool   `koanf:"insecureSkipTlsVerify"`
}

// TLSConfig builds a tls.Config based on the configuration.
func (c *TLS) TLSConfig() (*tls.Config, error) {
	return tlscfg.New(
		tlscfg.MaybeWithDiskCA(c.CaFilepath, tlscfg.ForClient),
		tlscfg.MaybeWithDiskKeyPair(c.CertFilepath, c.KeyFilepath),
		tlscfg.WithOverride(func(config *tls.Config) error {
			if c.InsecureSkipTLSVerify {
				config.InsecureSkipVerify = true
			}
			return nil
		}),
	)
}
