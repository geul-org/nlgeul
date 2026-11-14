//ff:func feature=token type=helper control=sequence
//ff:what TestBoxWaysLinesThree — three ways of one node each: the inner bracket of each is a single "─", the outer bracket has "┤" at the first line of a way between, and the middle line, which is the first line of the second way, is "┼" and has both tags
package token

import (
	"slices"
	"testing"
)

func TestBoxWaysLinesThree(t *testing.T) {
	region := Span{0, 3}
	segs := []Segment{
		boxSegment(1, 3, -1, region, boxNode(0, 3, "가나다", "가나다")),
		boxSegment(2, 3, -2, region, boxNode(0, 2, "가나", "가나")),
		boxSegment(3, 3, -3, region, boxNode(0, 3, "가나다", "가나다", "가나", "가")),
	}
	want := []string{
		"┌ [0-3] 가나다 ─ 후보경합 1 ───┐ ─┐ ① λ 합 -1.00",
		"│ 가나다  성립 1.00 선호 -3.25 │  │",
		"└──────────────────────────────┘  │",
		"┌ [0-2] 가나 ─ 후보경합 1 ─────┐ ─┼ ② λ 합 -2.00  어절경합 [0-3]",
		"│ 가나  성립 1.00 선호 -3.25   │  │",
		"└──────────────────────────────┘  │",
		"┌ [0-3] 가나다 ─ 후보경합 3 ───┐ ─┘ ③ λ 합 -3.00",
		"│ 가나다  성립 1.00 선호 -3.25 │",
		"│ 가나    성립 0.75 선호 -6.50 │",
		"│ (1개 생략됨)                 │",
		"└──────────────────────────────┘",
	}
	if got := boxWaysLines(segs, 2); !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
}
