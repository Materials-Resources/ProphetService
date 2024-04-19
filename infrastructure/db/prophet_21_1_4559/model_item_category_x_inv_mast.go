package prophet_21_1_4559

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type ItemCategoryXInvMast struct {
	bun.BaseModel           `bun:"table:item_category_x_inv_mast"`
	ItemCategoryXInvMastUid int32          `bun:"item_category_x_inv_mast_uid,type:int,pk"`
	ItemCategoryUid         int32          `bun:"item_category_uid,type:int"`
	InvMastUid              int32          `bun:"inv_mast_uid,type:int"`
	SequenceNo              int32          `bun:"sequence_no,type:int"`
	DisplayOnWebFlag        string         `bun:"display_on_web_flag,type:char,default:('N')"`
	DeleteFlag              string         `bun:"delete_flag,type:char,default:('N')"`
	DateCreated             time.Time      `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string         `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time      `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string         `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DisplayDesc             string         `bun:"display_desc,type:varchar(255)"`
	Comments                sql.NullString `bun:"comments,type:varchar(255),nullzero"`
	AlternateCodeUid        sql.NullInt32  `bun:"alternate_code_uid,type:int,nullzero"`
}

type ItemCategoryXInvMastModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns a slice of ItemCategoryXInvMast by the given InvMastUid
func (m ItemCategoryXInvMastModel) GetByInvMastUid(ctx context.Context, invMastUid int32) (
	[]*ItemCategoryXInvMast, error) {
	var itemCategoryXInvMasts []*ItemCategoryXInvMast
	err := m.bun.NewSelect().Model(&itemCategoryXInvMasts).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return itemCategoryXInvMasts, nil
}

// Delete deletes the record from the database.
func (m ItemCategoryXInvMastModel) Delete(ctx context.Context, itemCategoryXInvMast *ItemCategoryXInvMast) error {
	_, err := m.bun.NewDelete().Model(itemCategoryXInvMast).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
