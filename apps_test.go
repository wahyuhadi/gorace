package main

import (
	"testing"

	"github.com/wahyuhadi/gorace/models"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestCheckOptions(t *testing.T) {
	type args struct {
		opts *models.Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"User not input file locations", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckOptions(tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("CheckOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
