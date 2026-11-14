//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBoxLines — the title line with dashes up to the right line, a line for each row and the bottom line, all as wide as width + 4 cells; at a wider width the same box has the same width everywhere
package token

import (
	"slices"
	"testing"
)

func TestBoxLines(t *testing.T) {
	b := makeBox(boxNode(0, 1, "가", "가나"), 0)
	got := b.lines(b.width)
	want := []string{
		"┌ [0-1] 가 ─ 후보경합 1 ─────┐",
		"│ 가나  성립 1.00 선호 -3.25 │",
		"└────────────────────────────┘",
	}
	if !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
	for _, l := range append(b.lines(b.width), b.lines(b.width+7)...) {
		if n := cellWidth(l); n != b.width+4 && n != b.width+11 {
			t.Fatalf("%q is %d cells", l, n)
		}
	}
	if !slices.Equal(b.lines(b.width), got) {
		t.Fatal("twice")
	}
}
