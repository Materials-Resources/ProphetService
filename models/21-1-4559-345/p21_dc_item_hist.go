package model

import "github.com/uptrace/bun"

type P21DcItemHist struct {
	bun.BaseModel               `bun:"table:p21_dc_item_hist"`
	ACompanyId                  string  `bun:"A_Company_ID,type:char(16),nullzero"`
	BLocationId                 string  `bun:"B_Location_ID,type:char(16),nullzero"`
	CItemId                     string  `bun:"C_Item_ID,type:char(40),nullzero"`
	DYearForPeriod              string  `bun:"D_Year_for_Period,type:char(16),nullzero"`
	EPeriod                     string  `bun:"E_Period,type:char(16),nullzero"`
	FPeriodUsage                float64 `bun:"F_Period_Usage,type:float,nullzero"`
	GNumberOfOrders             string  `bun:"G_Number_of_Orders,type:numeric,nullzero"`
	HScheduledUsage             string  `bun:"H_Scheduled_Usage,type:char(16),nullzero"`
	IForecastUsage              string  `bun:"I_Forecast_Usage,type:char(16),nullzero"`
	JForecastDeviationPercentag string  `bun:"J_Forecast_Deviation_Percentag,type:char(16),nullzero"`
	KMadPercentage              string  `bun:"K_Mad_Percentage,type:char(16),nullzero"`
	LFilteredUsage              string  `bun:"L_Filtered_Usage,type:char(16),nullzero"`
	MEdited                     string  `bun:"M_Edited,type:char(16),nullzero"`
}
