package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/data"
	"github.com/materials-resources/s-prophet/internal/billing/domain"
	"reflect"
	"testing"
)

func TestBilling_GetInvoicesByOrder(t *testing.T) {
	type fields struct {
		model *data.Models
	}
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Invoice
		wantErr bool
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Billing{
				model: tt.fields.model,
			}
			got, err := s.GetInvoicesByOrder(tt.args.ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInvoicesByOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInvoicesByOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
