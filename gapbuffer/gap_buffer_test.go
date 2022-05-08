package gapbuffer_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mizuochikeita/algorithm/gapbuffer"
)

func TestGapBuffer(t *testing.T) {
	t.Run("Insert()", func(t *testing.T) {
		tests := []struct {
			desc        string
			size        int
			initString  string
			cursor      int
			inputString string
			wantString  string
			wantCursor  int
		}{
			{
				desc:        "shorter than size",
				size:        10,
				initString:  "ab",
				cursor:      0,
				inputString: "",
				wantString:  "ab",
				wantCursor:  0,
			},
			{
				desc:        "same as size",
				size:        2,
				initString:  "ab",
				cursor:      0,
				inputString: "",
				wantString:  "ab",
				wantCursor:  0,
			},
			{
				desc:        "longer than size",
				size:        2,
				initString:  "abc",
				cursor:      0,
				inputString: "",
				wantString:  "abc",
				wantCursor:  0,
			},
			{
				desc:        "much longer than size",
				size:        2,
				initString:  "abcdefghijklmn",
				cursor:      0,
				inputString: "",
				wantString:  "abcdefghijklmn",
				wantCursor:  0,
			},
			{
				desc:        "between words",
				size:        2,
				initString:  "hello world",
				cursor:      5,
				inputString: " bye",
				wantString:  "hello bye world",
				wantCursor:  9,
			},
		}
		for _, tt := range tests {
			gb := gapbuffer.New(tt.size)
			gb.Insert(tt.initString)
			gb.SetCursor(tt.cursor)
			gb.Insert(tt.inputString)
			if diff := cmp.Diff(tt.wantString, gb.String()); diff != "" {
				t.Errorf("%s: %s", tt.desc, diff)
			}
			if diff := cmp.Diff(tt.wantCursor, gb.Cursor()); diff != "" {
				t.Errorf("%s: %s", tt.desc, diff)
			}
		}
	})

	t.Run("SetCursor()", func(t *testing.T) {
		tests := []struct {
			desc   string
			size   int
			s      string
			cursor int
		}{
			{
				desc:   "go to head",
				size:   10,
				s:      "foobar",
				cursor: 0,
			},
		}
		for _, tt := range tests {
			gb := gapbuffer.New(tt.size)
			gb.Insert(tt.s)
			gb.SetCursor(tt.cursor)
			if diff := cmp.Diff(tt.cursor, gb.Cursor()); diff != "" {
				t.Errorf("%s/cursor: %s", tt.desc, diff)
			}
			if diff := cmp.Diff(tt.s, gb.String()); diff != "" {
				t.Errorf("%s/string: %s", tt.desc, diff)
			}
		}
	})

	t.Run("Grow()", func(t *testing.T) {
		tests := []struct {
			desc     string
			s        string
			initSize int
			growSize int
		}{}
		for _, tt := range tests {
			gb := gapbuffer.New(tt.initSize)
			gb.Insert(tt.s)
			gb.Grow(tt.growSize)
			if diff := cmp.Diff(tt.s, gb.String()); diff != "" {
				t.Errorf("string changed/%s: %s", tt.desc, diff)
			}
		}
	})
}
