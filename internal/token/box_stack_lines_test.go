//ff:func feature=token type=helper control=sequence
//ff:what TestBoxStackLines — each node in a box of its own width, one over the other, with nothing at the right (the boxes of two nodes differ in width)
package token

import (
	"slices"
	"testing"
)

func TestBoxStackLines(t *testing.T) {
	seg := boxSegment(1, 1, 0, Span{0, 6}, boxNode(0, 3, "아버지", "아버지", "아버"), boxNode(3, 6, "가방에", "가"))
	want := []string{
		"┌ [0-3] 아버지 ─ 후보경합 2 ───┐",
		"│ 아버지  성립 1.00 선호 -3.25 │",
		"│ 아버    성립 0.75 선호 -6.50 │",
		"└──────────────────────────────┘",
		"┌ [3-6] 가방에 ─ 후보경합 1 ──┐",
		"│ 가  성립 1.00 선호 -3.25    │",
		"└─────────────────────────────┘",
	}
	if got := boxStackLines(seg, 0); !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
	if got := boxStackLines(Segment{}, 0); len(got) != 0 {
		t.Fatal(got)
	}
}
