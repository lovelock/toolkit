package handlers

import (
	"errors"
	"testing"
)

func TestPrintErrorAndGoon(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "print error and go on",
			args: args{
				e: errors.New("this is an subtle error and program can go on"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintErrorAndGoon(tt.args.e)
		})
	}
}

func TestLogErrorAndGoon(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "log error and go on",
			args: args{
				e: errors.New("this is an subtle error and program can go on"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogErrorAndGoon(tt.args.e)
		})
	}
}
