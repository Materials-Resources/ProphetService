package config

type Kafka struct {
	Brokers  []string `koanf:"brokers"`
	Registry []string `koanf:"registry"`
}
