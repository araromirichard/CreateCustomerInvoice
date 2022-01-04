package main

import (
	"reflect"
	"testing"
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
