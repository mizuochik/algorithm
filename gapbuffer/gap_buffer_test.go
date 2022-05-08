package gapbuffer_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mizuochikeita/algorithm/gapbuffer"
)

func TestGapBuffer(t *testing.T) {
	t.Run("Insert()", func(t *testing.T) {
		tests := []struct {
			desc string
			size int
			s    string
			want string
		}{
			{"shorter than size", 10, "ab", "ab"},
		}
		for _, tt := range tests {
			gb := gapbuffer.New(tt.size)
			gb.Insert(tt.s)
			if diff := cmp.Diff(tt.want, gb.String()); diff != "" {
				t.Errorf("%s: %s", tt.desc, diff)
			}
		}
	})
}
