package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceValue struct {
	bun.BaseModel       `bun:"table:job_price_value"`
	JobPriceValueUid    int32     `bun:"job_price_value_uid,type:int,pk"`                   // Unique ID for each record in the table.
	JobPriceLineUid     int32     `bun:"job_price_line_uid,type:int,unique"`                // Unique ID for associated Job/Contract Line Items.
	CustomerId          float64   `bun:"customer_id,type:decimal(9,0),unique"`              // The customer identifier.
	ShipToId            float64   `bun:"ship_to_id,type:decimal(9,0),unique"`               // The ship to identifier.
	CalculationMethodCd int32     `bun:"calculation_method_cd,type:int"`                    // Multiplier, Difference, Mark up & Percentage
	CalculationValue1   float64   `bun:"calculation_value1,type:decimal(19,9),default:(0)"` // Calculation value to be used to calculate the price for each breaks.
	CalculationValue2   float64   `bun:"calculation_value2,type:decimal(19,9),default:(0)"`
	CalculationValue3   float64   `bun:"calculation_value3,type:decimal(19,9),default:(0)"`
	CalculationValue4   float64   `bun:"calculation_value4,type:decimal(19,9),default:(0)"`
	CalculationValue5   float64   `bun:"calculation_value5,type:decimal(19,9),default:(0)"`
	CalculationValue6   float64   `bun:"calculation_value6,type:decimal(19,9),default:(0)"`
	CalculationValue7   float64   `bun:"calculation_value7,type:decimal(19,9),default:(0)"`
	CalculationValue8   float64   `bun:"calculation_value8,type:decimal(19,9),default:(0)"`
	CalculationValue9   float64   `bun:"calculation_value9,type:decimal(19,9),default:(0)"`
	CalculationValue10  float64   `bun:"calculation_value10,type:decimal(19,9),default:(0)"`
	CalculationValue11  float64   `bun:"calculation_value11,type:decimal(19,9),default:(0)"`
	CalculationValue12  float64   `bun:"calculation_value12,type:decimal(19,9),default:(0)"`
	CalculationValue13  float64   `bun:"calculation_value13,type:decimal(19,9),default:(0)"`
	CalculationValue14  float64   `bun:"calculation_value14,type:decimal(19,9),default:(0)"`
	CalculationValue15  float64   `bun:"calculation_value15,type:decimal(19,9),default:(0)"`
	Break1              float64   `bun:"break1,type:decimal(19,9),default:(0)"` // Indicates the Quantity Breaks.
	Break2              float64   `bun:"break2,type:decimal(19,9),default:(0)"`
	Break3              float64   `bun:"break3,type:decimal(19,9),default:(0)"`
	Break4              float64   `bun:"break4,type:decimal(19,9),default:(0)"`
	Break5              float64   `bun:"break5,type:decimal(19,9),default:(0)"`
	Break6              float64   `bun:"break6,type:decimal(19,9),default:(0)"`
	Break7              float64   `bun:"break7,type:decimal(19,9),default:(0)"`
	Break8              float64   `bun:"break8,type:decimal(19,9),default:(0)"`
	Break9              float64   `bun:"break9,type:decimal(19,9),default:(0)"`
	Break10             float64   `bun:"break10,type:decimal(19,9),default:(0)"`
	Break11             float64   `bun:"break11,type:decimal(19,9),default:(0)"`
	Break12             float64   `bun:"break12,type:decimal(19,9),default:(0)"`
	Break13             float64   `bun:"break13,type:decimal(19,9),default:(0)"`
	Break14             float64   `bun:"break14,type:decimal(19,9),default:(0)"`
	OtherCost1          *float64  `bun:"other_cost1,type:decimal(19,9)"` // Other Cost corresponding to the quantity break that is used to price the item.
	OtherCost2          *float64  `bun:"other_cost2,type:decimal(19,9)"`
	OtherCost3          *float64  `bun:"other_cost3,type:decimal(19,9)"`
	OtherCost4          *float64  `bun:"other_cost4,type:decimal(19,9)"`
	OtherCost5          *float64  `bun:"other_cost5,type:decimal(19,9)"`
	OtherCost6          *float64  `bun:"other_cost6,type:decimal(19,9)"`
	OtherCost7          *float64  `bun:"other_cost7,type:decimal(19,9)"`
	OtherCost8          *float64  `bun:"other_cost8,type:decimal(19,9)"`
	OtherCost9          *float64  `bun:"other_cost9,type:decimal(19,9)"`
	OtherCost10         *float64  `bun:"other_cost10,type:decimal(19,9)"`
	OtherCost11         *float64  `bun:"other_cost11,type:decimal(19,9)"`
	OtherCost12         *float64  `bun:"other_cost12,type:decimal(19,9)"`
	OtherCost13         *float64  `bun:"other_cost13,type:decimal(19,9)"`
	OtherCost14         *float64  `bun:"other_cost14,type:decimal(19,9)"`
	OtherCost15         *float64  `bun:"other_cost15,type:decimal(19,9)"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`                                     // Status of a row (e.g., active, inactive, deleted)
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
