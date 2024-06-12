package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandForecastFormula struct {
	bun.BaseModel            `bun:"table:demand_forecast_formula"`
	DemandForecastFormulaUid int32     `bun:"demand_forecast_formula_uid,type:int,pk"`
	Factor1                  float64   `bun:"factor1,type:decimal(19,9),nullzero"`
	Factor2                  float64   `bun:"factor2,type:decimal(19,9),nullzero"`
	Factor3                  float64   `bun:"factor3,type:decimal(19,9),nullzero"`
	Factor4                  float64   `bun:"factor4,type:decimal(19,9),nullzero"`
	Factor5                  float64   `bun:"factor5,type:decimal(19,9),nullzero"`
	Factor6                  float64   `bun:"factor6,type:decimal(19,9),nullzero"`
	Factor7                  float64   `bun:"factor7,type:decimal(19,9),nullzero"`
	Factor8                  float64   `bun:"factor8,type:decimal(19,9),nullzero"`
	Factor9                  float64   `bun:"factor9,type:decimal(19,9),nullzero"`
	Factor10                 float64   `bun:"factor10,type:decimal(19,9),nullzero"`
	Factor11                 float64   `bun:"factor11,type:decimal(19,9),nullzero"`
	Factor12                 float64   `bun:"factor12,type:decimal(19,9),nullzero"`
	Factor13                 float64   `bun:"factor13,type:decimal(19,9),nullzero"`
	SequenceNo               int32     `bun:"sequence_no,type:int"`
	VelocityTypeCd           int32     `bun:"velocity_type_cd,type:int"`
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	FormulaName              string    `bun:"formula_name,type:varchar(255),nullzero"`
}
