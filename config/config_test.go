package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNewConfig_Read(t *testing.T) {
	testCases := []struct {
		name       string
		input      []byte
		wantConfig *Config
		wantErr    bool
	}{
		{
			name: "Valid file",
			input: []byte(`---
database:
  host: "sql.corp.com:1433"
  user: "admin"
  password: "password"
  name: "prophet"
observability:
  service: "prophet"
  environment: "development"
  id: "prophet"
kafka:
  brokers:
    - "kafka.corp.com:9092"
  registry:
    - "schema-registry.corp.com:8081"
environment: "development"
defaults:
  - company_id: "CORP"
  - asset_account_no: "1000"
  - revenue_account_no: "2000"
  - cos_account_no: "3000"
`),
			wantErr: false,
			wantConfig: &Config{
				Environment: "development",
				Defaults: map[string]string{
					"company_id":         "CORP",
					"asset_account_no":   "1000",
					"revenue_account_no": "2000",
					"cos_account_no":     "3000",
				},
				Database: DatabaseConfig{
					Host:     "sql.corp.com:1433",
					User:     "admin",
					Password: "password",
					Name:     "prophet",
				},
				Kafka: KafkaConfig{
					Brokers:  []string{"kafka.corp.com:9092"},
					Registry: []string{"schema-registry.corp.com:8081"},
				},
				Observability: ObservabilityConfig{
					Service:     "prophet",
					Environment: "development",
					Id:          "prophet",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// create config
			file, err := os.CreateTemp("", "config.*.yaml")
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println(file.Name())
			defer os.Remove(file.Name())

			_, err = file.Write(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			got := NewConfig(file.Name())
			if !reflect.DeepEqual(got, tc.wantConfig) {
				t.Errorf("NewConfig() got = %v, want %v", got, tc.wantConfig)
			}
		})
	}
}
