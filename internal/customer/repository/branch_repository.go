package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/materials-resources/s-prophet/internal/customer/domain"
	"github.com/uptrace/bun"
	"strconv"
)

func NewBranchRepository(db bun.IDB) *BranchRepository {
	return &BranchRepository{db: db}
}

type BranchRepository struct {
	db bun.IDB
}

func (r *BranchRepository) GetBranchesByContact(ctx context.Context, contactID string) ([]*domain.Branch, error) {
	contactRec := Contacts{}

	err := r.db.NewSelect().Model(&contactRec).Relation("ContactsXShipTo", func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Relation("ShipTo").Relation("ShipTo.Address")
	}).Where("id = ?", contactID).Scan(ctx)

	if err != nil {
		return nil, err
	}

	var branches []*domain.Branch

	for _, shipToRec := range contactRec.ContactsXShipTo {
		branches = append(branches, &domain.Branch{
			Id:   strconv.FormatFloat(shipToRec.ShipTo.ShipToId, 'f', 0, 64),
			Name: shipToRec.ShipTo.Address.Name,
		})
	}

	return branches, nil

}

func (r *BranchRepository) GetBranchById(ctx context.Context, id string) (*domain.Branch, error) {
	shipToRec := ShipTo{}

	// TODO: Dynamic company_id
	err := r.db.NewSelect().Model(&shipToRec).Relation("Address").Where("ship_to_id = ? and company_id = ?", id, "MRS").Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	branch := &domain.Branch{
		Id:   strconv.FormatFloat(shipToRec.ShipToId, 'f', 0, 64),
		Name: shipToRec.Address.Name,
	}

	return branch, nil

}
