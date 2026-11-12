//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCutterFacts — what stands right on each side of the sign at a place (nothing at the sentence's edge), a space or an edge on both sides sets MarkApart, a hyphen or an equals sign right next to it sets MarkArrow; the facts that hang on what stands open are not set
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCutterFacts(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		text        string
		at          int
		left, right int
		facts       int
	}{
		{"(가", 0, 0, rule.SideHangul, 0},
		{"가)", 1, rule.SideHangul, 0, 0},
		{"(", 0, 0, 0, rule.MarkApart},
		{"x < y", 2, 0, 0, rule.MarkApart},
		{"x<y", 1, rule.SideAlnum, rule.SideAlnum, 0},
		{"A -> B", 3, 0, 0, rule.MarkArrow},
		{"A->B", 2, 0, rule.SideAlnum, rule.MarkArrow},
		{"a<=b", 1, rule.SideAlnum, 0, rule.MarkArrow},
		{"5'10", 1, rule.SideAlnum | rule.SideDigit, rule.SideAlnum | rule.SideDigit, 0},
		{"가 (나", 2, 0, rule.SideHangul, 0},
	}
	for _, c := range cases {
		m := newCutter(ix, c.text, false).facts(c.at)
		if m.Left != c.left || m.Right != c.right || m.Facts != c.facts || m.Run != 1 {
			t.Errorf("%q at %d: left %b right %b facts %b run %d", c.text, c.at, m.Left, m.Right, m.Facts, m.Run)
		}
	}
}
