package network

import (
	"reflect"
	"testing"
)

func TestGetIpsFromHosts(t *testing.T) {
	type args struct {
		hosts []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "GetIpsFromHosts",
			args: args{
				hosts: []string{"www.baidu.com"},
			},
			want: []string{"110.242.68.4", "110.242.68.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIpsFromHosts(tt.args.hosts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIpsFromHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
