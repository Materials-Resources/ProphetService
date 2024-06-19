package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandForecastFormula struct {
	bun.BaseModel            `bun:"table:demand_forecast_formula"`
	DemandForecastFormulaUid int32     `bun:"demand_forecast_formula_uid,type:int,pk"`                      // Unique identifier for demand_forecast_formula
	Factor1                  float64   `bun:"factor1,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor2                  float64   `bun:"factor2,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor3                  float64   `bun:"factor3,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor4                  float64   `bun:"factor4,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor5                  float64   `bun:"factor5,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor6                  float64   `bun:"factor6,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor7                  float64   `bun:"factor7,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor8                  float64   `bun:"factor8,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor9                  float64   `bun:"factor9,type:decimal(19,9),nullzero"`                          // Period weighting for weighted average formula.
	Factor10                 float64   `bun:"factor10,type:decimal(19,9),nullzero"`                         // Period weighting for weighted average formula.
	Factor11                 float64   `bun:"factor11,type:decimal(19,9),nullzero"`                         // Period weighting for weighted average formula.
	Factor12                 float64   `bun:"factor12,type:decimal(19,9),nullzero"`                         // Period weighting for weighted average formula.
	Factor13                 float64   `bun:"factor13,type:decimal(19,9),nullzero"`                         // Period weighting for weighted average formula.
	SequenceNo               int32     `bun:"sequence_no,type:int"`                                         // Used in ordering of formulas.
	VelocityTypeCd           int32     `bun:"velocity_type_cd,type:int"`                                    // Categorizes formulas based on velocity level.
	RowStatusFlag            int32     `bun:"row_status_flag,type:int"`                                     // Row status flag.
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	FormulaName              string    `bun:"formula_name,type:varchar(255),nullzero"`                      // Custom column for the name of the formula
}
