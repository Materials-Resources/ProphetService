package app

import (
	"os"

	"gopkg.in/yaml.v3"
)

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
	App struct {
		Events struct {
			Brokers []string `yaml:"brokers"`
		} `yaml:"events"`
		Defaults struct {
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

func NewConfigFromPath(path string) (*Config, error) {

	var config Config

	// Open config file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Decode config file into Config
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil

}
