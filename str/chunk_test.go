package str

import (
	"reflect"
	"testing"
)

func TestChunkByLength(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "chunk by length no remains",
			args: args{
				s: `abcdefghijkl`,
				n: 2,
			},
			want: []string{"ab", "cd", "ef", "gh", "ij", "kl"},
		},
		{
			name: "chunk by length one remain",
			args: args{
				s: `abcdefghijklm`,
				n: 3,
			},
			want: []string{"abc", "def", "ghi", "jkl", "m"},
		},
		{
			name: "chunk by length two remains",
			args: args{
				s: `abcdefghijklmn`,
				n: 3,
			},
			want: []string{"abc", "def", "ghi", "jkl", "mn"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChunkByLength(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkByLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
