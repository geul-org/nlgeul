//ff:func feature=token type=reader control=sequence
//ff:what TestDigitLefts — when letters run on, only the runs after digits are kept; else all of them
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDigitLefts(t *testing.T) {
	sc := scanOf("x")
	lefts := []tail{{start: 1}, {start: 1, as: numName}}
	if got := sc.digitLefts(lefts, rule.RightOther); len(got) != 1 || got[0].as != "" {
		t.Fatal(got)
	}
	if got := sc.digitLefts(lefts, rule.RightDone); len(got) != 2 {
		t.Fatal(got)
	}
}
