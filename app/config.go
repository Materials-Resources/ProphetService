package app

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const EnvironmentDevelopment = "development"

type Config struct {
	Tracing struct {
		Service     string `yaml:"service"`
		Environment string `yaml:"environment"`
		Id          string `yaml:"id"`
	} `yaml:"tracing"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"database"`

	Broker struct {
		Host string `yaml:"host"`
	} `yaml:"broker"`

	App struct {
		Environment string `yaml:"environment"`
		Defaults    struct {
			InvMast struct {
				DefaultSalesDiscountGroup string `yaml:"default_sales_discount_group"`
			} `yaml:"inv_mast"`
			ProductGroup struct {
				CompanyId        string `yaml:"company_id"`
				AssetAccountNo   string `yaml:"asset_account_no"`
				RevenueAccountNo string `yaml:"revenue_account_no"`
				CosAccountNo     string `yaml:"cos_account_no"`
			} `yaml:"product_group"`
		} `yaml:"defaults"`
	} `yaml:"app"`
}

func NewConfig(configPath string) *Config {

	var config Config

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Decode config file into Config
	if err := d.Decode(&config); err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	return &config

}
