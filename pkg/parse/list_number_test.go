//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestListNumber — a round closing bracket and its full-width form right after a digit number an item of a list; not after a letter or with nothing before them, and no other closing sign after a digit
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestListNumber(t *testing.T) {
	digit := rule.Mark{Left: rule.SideAlnum | rule.SideDigit}
	cases := []struct {
		sign rune
		m    rule.Mark
		want bool
	}{
		{')', digit, true},
		{'）', digit, true},
		{')', rule.Mark{Left: rule.SideHangul}, false},
		{')', rule.Mark{}, false},
		{']', digit, false},
		{'>', digit, false},
		{'”', digit, false},
	}
	for _, c := range cases {
		if got := listNumber(c.sign, c.m); got != c.want {
			t.Errorf("%c %+v: %v", c.sign, c.m, got)
		}
	}
}
