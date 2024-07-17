package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
	"strings"
)

const EnvironmentDevelopment = "development"

type Config struct {
	Environment   string              `koanf:"environment"`
	Database      DatabaseConfig      `koanf:"database"`
	Observability ObservabilityConfig `koanf:"observability"`

	Kafka KafkaConfig `koanf:"kafka"`

	Defaults map[string]string `koanf:"defaults"`
}

type ObservabilityConfig struct {
	Service     string `koanf:"service"`
	Environment string `koanf:"environment"`
	Id          string `koanf:"id"`
}

type DatabaseConfig struct {
	Host     string `koanf:"host"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
}

type KafkaConfig struct {
	Brokers  []string `koanf:"brokers"`
	Registry []string `koanf:"registry"`
}

// NewConfig reads from a config file or environment variables starting with PROPHET_ and returns a Config struct
func NewConfig(configPath string) *Config {

	var k = koanf.New(".")

	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		log.Fatalf("error loading config file: %v", err)
	}

	k.Load(env.Provider("PROPHET_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "PROPHET_")), "_", ".", -1)
	}), nil)

	var config Config

	k.Unmarshal("", &config)

	return &config

}
