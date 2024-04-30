package data

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type QuoteLine struct {
	bun.BaseModel    `bun:"table:quote_line"`
	QuoteLineUid     int32     `bun:"quote_line_uid,type:int,pk"`
	OeLineUid        int32     `bun:"oe_line_uid,type:int"`
	LineCompleteFlag string    `bun:"line_complete_flag,type:char,default:('N')"`
	QtyConverted     float64   `bun:"qty_converted,type:decimal(19,9)"`
	UomConverted     string    `bun:"uom_converted,type:varchar(8)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}

var _ bun.BeforeAppendModelHook = (*QuoteLine)(nil)

func (m *QuoteLine) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.LineCompleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type QuoteLineModel struct {
	bun bun.IDB
}

type CreateQuoteLineParams struct {
	OeLineUid    int32
	QtyConverted float64
	UomConverted string
}

func (m *QuoteLineModel) Create(ctx context.Context, params CreateQuoteLineParams) (*QuoteLine, error) {
	quoteLine := &QuoteLine{
		OeLineUid:        params.OeLineUid,
		QtyConverted:     params.QtyConverted,
		UomConverted:     params.UomConverted,
		LastMaintainedBy: "admin",
	}

	err := m.generateQuoteLineUid(ctx, &quoteLine.QuoteLineUid)
	if err != nil {
		return nil, err
	}
	err = m.Insert(ctx, quoteLine)
	if err != nil {
		return nil, err
	}
	return quoteLine, nil
}

func (m *QuoteLineModel) Insert(ctx context.Context, quoteLine *QuoteLine) error {
	_, err := m.bun.NewInsert().Model(quoteLine).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *QuoteLineModel) generateQuoteLineUid(ctx context.Context, quoteLineUid *int32) error {
	q := `DECLARE @uid int
			EXEC @uid = p21_get_counter 'quote_line', 1
			SELECT @uid`
	err := m.bun.QueryRowContext(ctx, q).Scan(quoteLineUid)
	if err != nil {
		return err
	}
	return nil
}
