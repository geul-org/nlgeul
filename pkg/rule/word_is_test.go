//ff:func feature=rule type=rule control=sequence
//ff:what TestWordIs — applies when the word has one of the facts asked for: last, bound
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWordIs(t *testing.T) {
	last := toulmin.Specs{&Bits{Of: "word", Mask: WordLast}}
	bound := toulmin.Specs{&Bits{Of: "word", Mask: WordBound}}
	a, _ := WordIs(ctxFor(&Word{Last: true}, Tail{}), last)
	b, _ := WordIs(ctxFor(&Word{Last: true}, Tail{}), bound)
	c, _ := WordIs(ctxFor(&Word{Bound: true}, Tail{}), bound)
	d, _ := WordIs(ctxFor(&Word{}, Tail{}), last)
	if a != 1 || b != -1 || c != 1 || d != -1 {
		t.Fatal(a, b, c, d)
	}
}
