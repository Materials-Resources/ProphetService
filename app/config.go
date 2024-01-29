package app

import (
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Configuration Config
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"database"`
	App struct {
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

func NewConfig(path string) error {

	var config Config

	// Open config file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Decode config file into Config
	if err := d.Decode(&config); err != nil {
		return err
	}

	Configuration = config

	return nil

}
