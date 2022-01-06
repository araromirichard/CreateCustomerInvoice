package main

import (
	"testing"
)

func Test_createInvoice(t *testing.T) {
	type args struct {
		ciie CreateInvoiceInputEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Basic",
		// 	args: args{
		// 		ciie: CreateInvoiceInputEvent{
		// 			CustomerEmail:    "sederyn@gmail.com",
		// 			CustomerId:       "63863064",
		// 			Amount:           4000,
		// 			Currency:         "NGN",
		// 			DueDate:          "2022-02-15",
		// 			SendNotification: true,
		// 			Draft:            true,
		// 		},
		// 	},

		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := createInvoice(tt.args.ciie)
			if (err != nil) != tt.wantErr {
				t.Errorf("createInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
