//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestLoneDepth — a closing bracket with none of its kind open closes the outermost (loneBracketDepth, 1) whichever bracket it is, a quotation mark with none of its kind open 0 between a letter and a Hangul syllable, and 1 when a mark that ends a sentence stands right before it, when no letter stands right before it (a sentence's start), and when no Hangul syllable stands right after it; a round bracket right after a digit numbers an item of a list and is 0, any other bracket there 1
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLoneDepth(t *testing.T) {
	if loneBracketDepth != 1 {
		t.Fatal(loneBracketDepth)
	}
	onto := rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul}
	for sign, want := range map[rune]int{')': 1, '}': 1, ']': 1, '>': 1, '」': 1, '』': 1, '》': 1, '〉': 1, '｣': 1, '"': 0, '\'': 0, '”': 0, '’': 0} {
		if got := loneDepth(sign, onto); got != want {
			t.Errorf("%c: %d, want %d", sign, got, want)
		}
	}
	stop := rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Facts: rule.MarkAfterStop}
	start := rule.Mark{Right: rule.SideHangul, Facts: rule.MarkQuoted}
	spaced := rule.Mark{Left: rule.SideAlnum}
	for _, m := range []rule.Mark{stop, start, spaced} {
		if q, b := loneDepth('"', m), loneDepth('』', m); q != loneBracketDepth || b != loneBracketDepth {
			t.Errorf("with %+v: %d %d", m, q, b)
		}
		if q, r := loneDepth('\'', m), loneDepth('’', m); q != loneBracketDepth || r != loneBracketDepth {
			t.Errorf("with %+v: %d %d", m, q, r)
		}
	}
	digit := rule.Mark{Left: rule.SideAlnum | rule.SideDigit, Right: rule.SideHangul}
	for sign, want := range map[rune]int{')': 0, '）': 0, ']': 1, '>': 1, '』': 1, '"': 0, '’': 0} {
		if got := loneDepth(sign, digit); got != want {
			t.Errorf("%c after a digit: %d, want %d", sign, got, want)
		}
	}
}
