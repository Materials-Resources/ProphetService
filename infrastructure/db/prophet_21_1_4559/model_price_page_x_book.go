package prophet_21_1_4559

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type PricePageXBook struct {
	bun.BaseModel     `bun:"table:price_page_x_book"`
	PricePageXBookUid int32     `bun:"price_page_x_book_uid,type:int,pk"`
	PriceBookUid      int32     `bun:"price_book_uid,type:int"`
	PricePageUid      int32     `bun:"price_page_uid,type:int"`
	RowStatusFlag     int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30)"`

	PricePage *PricePage `bun:"rel:belongs-to,join:price_page_uid=price_page_uid"`
}

type PricePageXBookModel struct {
	bun bun.IDB
}

// GetByPricePageUid returns a slice of PricePageXBook by the given PricePageUid
func (m PricePageXBookModel) GetByPricePageUid(ctx context.Context, pricePageUid int32) (*PricePageXBook, error) {
	pricePageXBook := new(PricePageXBook)
	err := m.bun.NewSelect().Model(pricePageXBook).Where("price_page_uid = ?", pricePageUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return pricePageXBook, nil
}

// Delete deletes the PricePageXBook from the database.
func (m PricePageXBookModel) Delete(ctx context.Context, pricePageXBook *PricePageXBook) error {
	_, err := m.bun.NewDelete().Model(pricePageXBook).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
