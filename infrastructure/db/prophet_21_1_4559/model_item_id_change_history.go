package prophet_21_1_4559

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type ItemIdChangeHistory struct {
	bun.BaseModel          `bun:"table:item_id_change_history"`
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`
	OldItemId              string    `bun:"old_item_id,type:varchar(40)"`
	NewItemId              string    `bun:"new_item_id,type:varchar(40)"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ItemIdChangeHistoryUid int32     `bun:"item_id_change_history_uid,type:int,pk,identity"`
}

type ItemIdChangeHistoryModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns a slice of ItemIdChangeHistory by the given InvMastUid
func (m ItemIdChangeHistoryModel) GetByInvMastUid(ctx context.Context, invMastUid int32) (
	[]*ItemIdChangeHistory, error) {
	var itemIdChangeHistories []*ItemIdChangeHistory
	err := m.bun.NewSelect().Model(&itemIdChangeHistories).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return itemIdChangeHistories, nil

}

// Delete deletes the record from the database.
func (m ItemIdChangeHistoryModel) Delete(ctx context.Context, itemIdChangeHistory *ItemIdChangeHistory) error {
	_, err := m.bun.NewDelete().Model(itemIdChangeHistory).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
