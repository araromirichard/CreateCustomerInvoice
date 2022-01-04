package paystack

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		baseUrl   string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
		{
			name: "Basic",
			args: args{
				baseUrl:   "https://api.paystack.co",
				secretKey: "sk_test_1374752eb291108be0fff0424ca81369b0f266ac",
			},
			want: &Client{
				client:    *http.DefaultClient,
				baseUrl:   "https://api.paystack.co",
				secretKey: "sk_test_1374752eb291108be0fff0424ca81369b0f266ac",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.baseUrl, tt.args.secretKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCustomer(t *testing.T) {
	type args struct {
		customerEmail string
	}
	tests := []struct {
		name    string
		cl      *Client
		args    args
		want    *GetCustomerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Basic",
			cl: &Client{
				client:    *http.DefaultClient,
				baseUrl:   "https://api.paystack.co",
				secretKey: "sk_test_1374752eb291108be0fff0424ca81369b0f266ac",
			},
			args: args{
				customerEmail: "sederyn@gmail.com",
			},
			want: &GetCustomerResponse{
				Customer: Customer{
					ID:         63863064,
					FirstName:  "Richard",
					LastName:   "Gbegede",
					Email:      "sederyn@gmail.com",
					Code:       "CUS_o13f0zmewiauk7r",
					Phone:      "07012205938",
					RiskAction: "default",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cl.GetCustomer(tt.args.customerEmail)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
