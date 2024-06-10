package data

import (
	"context"
	"github.com/materials-resources/s-prophet/pkg/models"
	"github.com/uptrace/bun"
)

type invMast struct {
	models.InvMast `bun:",extend"`
}

type InvMastModel struct {
	bun bun.IDB
}

func NewInvMastModel(bun bun.IDB) InvMastModel {
	return InvMastModel{
		bun: bun,
	}
}

type InvMastCreateParams struct {
	Sn          string
	Name        string
	Description string
}

func (m *InvMastModel) Create(ctx context.Context, params *InvMastCreateParams) error {

	return nil
}

func (m *InvMastModel) deleteAssociatedRecords(ctx context.Context, invMastUid string) error {

	return nil
}
