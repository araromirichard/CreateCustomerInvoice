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
		//TODO: Add test cases.
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

			wantErr: false,
		},
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

func Test_createCustomerInvoice(t *testing.T) {
	type args struct {
		customerId string
		ciie       CreateInvoiceInputEvent
	}
	tests := []struct {
		name string
		args args
		//want    *paystack.Invoice
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Basic",
			args: args{
				customerId: "63863064",
				ciie: CreateInvoiceInputEvent{
					CustomerEmail:    "sederyn@gmail.com",
					CustomerId:       "63863064",
					Amount:           4000,
					Currency:         "NGN",
					DueDate:          "2022-12-29",
					SendNotification: false,
					Draft:            false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := createCustomerInvoice(tt.args.customerId, tt.args.ciie)
			if (err != nil) != tt.wantErr {
				t.Errorf("createCustomerInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_getOrCreateCustomer(t *testing.T) {
	type args struct {
		customerId    string
		customerEmail string
	}
	tests := []struct {
		name    string
		args    args
		//want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Basic",
			args: args{
				customerId:    "",
				customerEmail: "sedoruna@test.com",
			},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getOrCreateCustomer(tt.args.customerId, tt.args.customerEmail)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOrCreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("getOrCreateCustomer() = %v, want %v", got, tt.want)
			// }
		})
	}
}
