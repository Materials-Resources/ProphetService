package prophet_21_1_4559

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvLocMsp struct {
	bun.BaseModel      `bun:"table:inv_loc_msp"`
	InvLocMspUid       int32         `bun:"inv_loc_msp_uid,type:int,pk"`
	InvMastUid         int32         `bun:"inv_mast_uid,type:int"`
	LocationId         float64       `bun:"location_id,type:decimal(19,0)"`
	ReceiptProcessFlag string        `bun:"receipt_process_flag,type:char,default:('N')"`
	ReceiptProcessUid  sql.NullInt32 `bun:"receipt_process_uid,type:int,nullzero"`
	RowStatusFlag      int32         `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated        time.Time     `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string        `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time     `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string        `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}

type InvLocMspModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns a slice of InvLocMsp by the given InvMastUid
func (m InvLocMspModel) GetByInvMastUid(ctx context.Context, invMastUid int32) ([]*InvLocMsp, error) {
	var invLocMsps []*InvLocMsp
	err := m.bun.NewSelect().Model(&invLocMsps).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invLocMsps, nil
}

// Delete deletes the InvLocMsp from the database.
func (m InvLocMspModel) Delete(ctx context.Context, invLocMsp *InvLocMsp) error {
	_, err := m.bun.NewDelete().Model(invLocMsp).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
