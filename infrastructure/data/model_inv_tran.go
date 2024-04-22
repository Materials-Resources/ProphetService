package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type InvTran struct {
	bun.BaseModel          `bun:"table:inv_tran"`
	TransactionNumber      float64         `bun:"transaction_number,type:decimal(19,0),pk"`
	LocationId             float64         `bun:"location_id,type:decimal(19,0)"`
	Period                 sql.NullFloat64 `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod          sql.NullFloat64 `bun:"year_for_period,type:decimal(4,0),nullzero"`
	TransType              string          `bun:"trans_type,type:varchar(5)"`
	Quantity               float64         `bun:"quantity,type:decimal(19,9)"`
	UnitCostAmt            float64         `bun:"unit_cost_amt,type:decimal(19,9)"`
	Freight                float64         `bun:"freight,type:decimal(19,4)"`
	ReservedBeforeTrans    sql.NullFloat64 `bun:"reserved_before_trans,type:decimal(19,9),default:(0)"`
	DocumentNo             float64         `bun:"document_no,type:decimal(19,0)"`
	OnHandBeforeTrans      float64         `bun:"on_hand_before_trans,type:decimal(19,9)"`
	AllocatedBeforeTrans   float64         `bun:"allocated_before_trans,type:decimal(19,9)"`
	BackorderedBeforeTrans float64         `bun:"backordered_before_trans,type:decimal(19,9)"`
	OnPoBeforeTrans        float64         `bun:"on_po_before_trans,type:decimal(19,9)"`
	InTransitBeforeTrans   float64         `bun:"in_transit_before_trans,type:decimal(19,9)"`
	DateCreated            time.Time       `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified       time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CurrencyId             sql.NullFloat64 `bun:"currency_id,type:decimal(19,0),nullzero"`
	SubDocumentNo          sql.NullFloat64 `bun:"sub_document_no,type:decimal(19,0),nullzero"`
	QtyAllocated           float64         `bun:"qty_allocated,type:decimal(19,9)"`
	QtyOnBo                float64         `bun:"qty_on_bo,type:decimal(19,9)"`
	QtyOnPo                float64         `bun:"qty_on_po,type:decimal(19,9)"`
	QtyInTransit           float64         `bun:"qty_in_transit,type:decimal(19,9)"`
	LineNo                 sql.NullFloat64 `bun:"line_no,type:decimal(19,0),nullzero"`
	UnitOfMeasure          string          `bun:"unit_of_measure,type:varchar(8)"`
	UnitSize               float64         `bun:"unit_size,type:decimal(19,4)"`
	QtyReservedDueIn       sql.NullFloat64 `bun:"qty_reserved_due_in,type:decimal(19,9),default:(0)"`
	InvMastUid             int32           `bun:"inv_mast_uid,type:int"`
	InProcessBeforeTrans   float64         `bun:"in_process_before_trans,type:decimal(19,9),default:(0)"`
	QtyInProcess           float64         `bun:"qty_in_process,type:decimal(19,9),default:(0)"`
	UsedSpecificCostFlag   sql.NullString  `bun:"used_specific_cost_flag,type:char,nullzero"`
}

type InvTranModel struct {
	bun bun.IDB
}

func (m InvTranModel) GetByInvMastUidSubDocumentNo(
	ctx context.Context, invMastUid int32, subDocumentNo float64) ([]*InvTran, error) {
	var invTrans []*InvTran
	err := m.bun.NewSelect().Model(&invTrans).Where(
		"inv_mast_uid = ? AND sub_document_no = ?", invMastUid, subDocumentNo).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invTrans, nil
}

func (m InvTranModel) GetBySubDocumentNo(ctx context.Context, subDocumentNo float64) ([]*InvTran, error) {
	var invTrans []*InvTran
	err := m.bun.NewSelect().Model(&invTrans).Where("sub_document_no = ?", subDocumentNo).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invTrans, nil
}

// GetByInvMastUid returns a slice of InvTran by the given InvMastUid
func (m InvTranModel) GetByInvMastUid(ctx context.Context, invMastUid int32) ([]*InvTran, error) {
	var invTrans []*InvTran
	err := m.bun.NewSelect().Model(&invTrans).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invTrans, nil
}

// Delete deletes the record from the database.
func (m InvTranModel) Delete(ctx context.Context, invTran *InvTran) error {
	_, err := m.bun.NewDelete().Model(invTran).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
