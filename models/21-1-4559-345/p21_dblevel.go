package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type P21Dblevel struct {
	bun.BaseModel `bun:"table:p21_dblevel"`
	Version       int16     `bun:"version,type:tinyint,nullzero"`
	DateCreated   time.Time `bun:"date_created,type:datetime,default:(getdate())"`
}
