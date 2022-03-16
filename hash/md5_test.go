package hash

import "testing"

func TestMd5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "generic md5 function",
			args: args{
				s: "aksjdfjasljdfasjdfkjasdjasdfj",
			},
			want: "760f0283b0d42b84727c2664c73263f3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.s); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}
