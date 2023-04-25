package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Defaults struct {
			CompanyId        string `yaml:"company_id"`
			AssetAccountNo   string `yaml:"asset_account_no"`
			RevenueAccountNo string `yaml:"revenue_account_no"`
			CosAccountNo     string `yaml:"cos_account_no"`
		} `yaml:"defaults"`
	} `yaml:"server"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, err
}
