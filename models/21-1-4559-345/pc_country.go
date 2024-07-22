package prophet

import "github.com/uptrace/bun"

type PcCountry struct {
	bun.BaseModel    `bun:"table:pc_country"`
	CountrySkey      int32   `bun:"country_skey,type:int,pk"`                  // Padlock table
	CountryName      string  `bun:"country_name,type:varchar(50),unique"`      // Padlock table
	CountryCode      *string `bun:"country_code,type:char(3)"`                 // Padlock table
	PhoneCode        int32   `bun:"phone_code,type:int"`                       // Padlock table
	PhoneFormat      *string `bun:"phone_format,type:varchar(20)"`             // Padlock table
	StateReqFlag     string  `bun:"state_req_flag,type:char(1),default:('Y')"` // Padlock table
	CurrencyName     *string `bun:"currency_name,type:varchar(40)"`            // Padlock table
	CurrencyFormat   *string `bun:"currency_format,type:varchar(10)"`          // Padlock table
	PostalCodeFormat *string `bun:"postal_code_format,type:varchar(15)"`       // Padlock table
}
