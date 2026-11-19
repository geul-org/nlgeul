//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestLoneDepth — a closing bracket with none of its kind open closes the outermost (loneBracketDepth, 1) whichever bracket it is, a quotation mark with none of its kind open 0, and 1 when a mark that ends a sentence stands right before it
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLoneDepth(t *testing.T) {
	if loneBracketDepth != 1 {
		t.Fatal(loneBracketDepth)
	}
	for sign, want := range map[rune]int{')': 1, '}': 1, ']': 1, '>': 1, '」': 1, '』': 1, '》': 1, '〉': 1, '｣': 1, '"': 0, '\'': 0, '”': 0, '’': 0} {
		if got := loneDepth(sign, 0); got != want {
			t.Errorf("%c: %d, want %d", sign, got, want)
		}
	}
	for _, sign := range "\"'”’)』" {
		if got := loneDepth(sign, rule.MarkAfterStop); got != loneBracketDepth {
			t.Errorf("%c after a full stop: %d", sign, got)
		}
	}
}
