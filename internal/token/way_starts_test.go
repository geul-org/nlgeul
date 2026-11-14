//ff:func feature=token type=helper control=sequence
//ff:what TestWayStarts — the line each way starts at: 0, then after the lines of the ways before it
package token

import (
	"slices"
	"testing"
)

func TestWayStarts(t *testing.T) {
	ways := []boxWay{{lines: make([]string, 4)}, {lines: make([]string, 7)}, {lines: make([]string, 3)}}
	if got := wayStarts(ways); !slices.Equal(got, []int{0, 4, 11}) {
		t.Fatal(got)
	}
	if got := wayStarts(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
