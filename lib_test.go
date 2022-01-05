package main

import (
	"reflect"
	"testing"

	"github.com/karosaxy/paystack-client/pkg/client/paystack"
)

func Test_createInvoice(t *testing.T) {
	type args struct {
		ciie CreateInvoiceInputEvent
	}
	tests := []struct {
		name    string
		args    args
		want    *CreateInvoiceOutputEvent
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Basic",
			args: args{
				ciie: CreateInvoiceInputEvent{
					CustomerEmail:    "sederyn@gmail.com",
					CustomerId:       "63863064",
					Amount:           4000,
					Currency:         "NGN",
					DueDate:          "2022-02-15",
					SendNotification: true,
					Draft:            true,
				},
			},
			want: &CreateInvoiceOutputEvent{
				Message: "Payment request created",
				Invoice: &paystack.Invoice{
					ID:               9070195,
					Domain:           "test",
					Amount:           400,
					Currency:         "NGN",
					DueDate:          "2022-02-18T00:00:00.000Z",
					HasInvoice:       false,
					InvoiceNumber:    0,
					Description:      "",
					PdfUrl:           "",
					LineItems:        []paystack.LineItem{},
					Tax:              []paystack.Tax{},
					Customer:         paystack.Customer{},
					Status:           "pending",
					Paid:             false,
					SendNotification: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createInvoice(tt.args.ciie)
			if (err != nil) != tt.wantErr {
				t.Errorf("createInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
