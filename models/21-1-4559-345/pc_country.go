package model

import "github.com/uptrace/bun"

type PcCountry struct {
	bun.BaseModel    `bun:"table:pc_country"`
	CountrySkey      int32  `bun:"country_skey,type:int,pk"`
	CountryName      string `bun:"country_name,type:varchar(50)"`
	CountryCode      string `bun:"country_code,type:char(3),nullzero"`
	PhoneCode        int32  `bun:"phone_code,type:int"`
	PhoneFormat      string `bun:"phone_format,type:varchar(20),nullzero"`
	StateReqFlag     string `bun:"state_req_flag,type:char,default:('Y')"`
	CurrencyName     string `bun:"currency_name,type:varchar(40),nullzero"`
	CurrencyFormat   string `bun:"currency_format,type:varchar(10),nullzero"`
	PostalCodeFormat string `bun:"postal_code_format,type:varchar(15),nullzero"`
}
