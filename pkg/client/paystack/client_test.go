package paystack

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_getCustomer(t *testing.T) {
	type args struct {
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *GetCustomerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCustomer(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
