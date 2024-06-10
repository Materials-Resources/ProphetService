package models

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type QuoteHdr struct {
	bun.BaseModel    `bun:"table:quote_hdr"`
	QuoteHdrUid      int32        `bun:"quote_hdr_uid,type:int,pk"`
	OeHdrUid         int32        `bun:"oe_hdr_uid,type:int"`
	CompleteFlag     string       `bun:"complete_flag,type:char,default:('N')"`
	DateCreated      time.Time    `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time    `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string       `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	Date843Sent      sql.NullTime `bun:"date_843_sent,type:datetime,nullzero"`
	ExpirationDate   sql.NullTime `bun:"expiration_date,type:datetime,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*QuoteHdr)(nil)

func (m *QuoteHdr) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.CompleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type QuoteHdrModel struct {
	bun bun.IDB
}

type CreateQuoteHdrParams struct {
	OeHdrUid       int32
	ExpirationDate time.Time
}

func (m *QuoteHdrModel) Create(ctx context.Context, params CreateQuoteHdrParams) (*QuoteHdr, error) {
	quoteHdr := &QuoteHdr{
		OeHdrUid:         params.OeHdrUid,
		LastMaintainedBy: "admin",
	}

	err := m.generateQuoteHdrUid(ctx, &quoteHdr.QuoteHdrUid)
	if err != nil {
		return nil, err
	}
	err = m.Insert(ctx, quoteHdr)
	if err != nil {
		return nil, err
	}
	return quoteHdr, nil
}

func (m *QuoteHdrModel) Insert(ctx context.Context, quoteHdr *QuoteHdr) error {
	_, err := m.bun.NewInsert().Model(quoteHdr).Exec(ctx)
	return err
}

func (m *QuoteHdrModel) generateQuoteHdrUid(ctx context.Context, uid *int32) error {
	q := `DECLARE @uid int
			EXEC @uid = p21_get_counter 'quote_hdr', 1
			SELECT @uid`
	err := m.bun.QueryRowContext(ctx, q).Scan(uid)
	if err != nil {
		return err
	}
	return nil
}
