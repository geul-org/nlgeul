//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBoxWaysLines — two ways of a stretch: the boxes of both are as wide as the widest, the brackets are at the right of each, every box line ends at the same cell, the tags sit where they should; with fewer ways shown than there are, the count follows the last box, and with one way shown the outer bracket is a single "├"
package token

import (
	"slices"
	"testing"
)

func TestBoxWaysLines(t *testing.T) {
	region := Span{0, 6}
	segs := []Segment{
		boxSegment(1, 2, -11.98, region, boxNode(0, 3, "아버지", "아버지", "아버"), boxNode(3, 6, "가방에", "가방에")),
		boxSegment(2, 2, -12.37, region, boxNode(0, 4, "아버지가", "아버지가"), boxNode(4, 6, "방에", "방에")),
	}
	got := boxWaysLines(segs, 0)
	want := []string{
		"┌ [0-3] 아버지 ─ 후보경합 2 ─────┐ ┬┐ ① λ 합 -11.98",
		"│ 아버지  성립 1.00 선호 -3.25   │ ││",
		"│ 아버    성립 0.75 선호 -6.50   │ ││",
		"└────────────────────────────────┘ │├ 어절경합 [0-6]",
		"┌ [3-6] 가방에 ─ 후보경합 1 ─────┐ ┘│",
		"│ 가방에  성립 1.00 선호 -3.25   │  │",
		"└────────────────────────────────┘  │",
		"┌ [0-4] 아버지가 ─ 후보경합 1 ───┐ ┬┘ ② λ 합 -12.37",
		"│ 아버지가  성립 1.00 선호 -3.25 │ │",
		"└────────────────────────────────┘ │",
		"┌ [4-6] 방에 ─ 후보경합 1 ───────┐ ┘",
		"│ 방에  성립 1.00 선호 -3.25     │",
		"└────────────────────────────────┘",
	}
	if !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
	for _, l := range got {
		if r, ok := edgeAt(l, 34); !ok || (r != '┐' && r != '│' && r != '┘') {
			t.Fatalf("%q", l)
		}
	}

	one := boxWaysLines(segs[:1], 0)
	if len(one) != 8 || one[0] != "┌ [0-3] 아버지 ─ 후보경합 2 ───┐ ┬├ ① λ 합 -11.98  어절경합 [0-6]" || one[4] != "┌ [3-6] 가방에 ─ 후보경합 1 ───┐ ┘" {
		t.Fatalf("%q", one)
	}
	if one[7] != "                                 (1개 생략됨)" {
		t.Fatalf("%q", one[7])
	}
	if all := boxWaysLines(segs, 0); slices.Contains(all, one[7]) {
		t.Fatal("nothing left out")
	}
}
