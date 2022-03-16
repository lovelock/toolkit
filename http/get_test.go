package http

import (
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestGet(t *testing.T) {
	type args struct {
		url    string
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *resty.Response
		wantErr bool
	}{
		{
			name: "simple get should get status 200",
			args: args{
				url: "https://httpbin.org/get",
				params: map[string]interface{}{
					"hello": "world",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.url, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Status(), "200 OK") {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
