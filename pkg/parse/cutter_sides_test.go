//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCutterSides — where a quotation mark is written: a space or the sentence's start before it and a letter or digit right after it is an opening mark's place (MarkOpenSide), unless a marker that takes a quotation stands there (MarkQuoted with MarkShutSide: "라고 at a sentence's start); a letter right before it and a space or the end after it is a closing mark's place (MarkShutSide); between two letters, between spaces, after a comma or bracket and before a space, none
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCutterSides(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text string
		at   int
		want int
	}{
		{"\"곰", 0, rule.MarkOpenSide},
		{"그는 ”두 곰", 3, rule.MarkOpenSide},
		{"그는 '3번", 3, rule.MarkOpenSide},
		{"\"라고 했다", 0, rule.MarkQuoted | rule.MarkShutSide},
		{"그는 '이라는 말", 3, rule.MarkQuoted | rule.MarkShutSide},
		{"좋다\"고 했다", 2, rule.MarkQuoted | rule.MarkShutSide},
		{"했다' 로", 2, rule.MarkShutSide},
		{"했다‘", 2, rule.MarkShutSide},
		{"Mall\" 곰", 4, rule.MarkShutSide},
		{"곰\"이", 1, 0},
		{"곰 \" 곰", 2, 0},
		{"곰,\" 곰", 2, 0},
		{"곰\",", 1, 0},
		{"그는 \"(곰", 3, 0},
	} {
		if got := newCutter(ix, c.text, false).sides(c.at); got != c.want {
			t.Errorf("%q at %d: %b, want %b", c.text, c.at, got, c.want)
		}
	}
}
