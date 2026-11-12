//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCalledGate — the graph the word before the vocative 야 must pass lets a word through and turns down each dependent noun of the closed set
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCalledGate(t *testing.T) {
	e := calledGate()
	if e.name != callGate || e.g == nil {
		t.Fatal(e)
	}
	for _, stem := range []string{"철수", "문제", "abc", ""} {
		if v := verdict(e.g, rule.Piece{Form: stem, HasStem: true, Next: -1}); v <= 0 {
			t.Errorf("%q is turned down: %v", stem, v)
		}
	}
	bound := parse.ClosedWords()[boundSet]
	if len(bound) == 0 {
		t.Fatal("no dependent nouns in the closed set")
	}
	for _, stem := range bound {
		if v := verdict(e.g, rule.Piece{Form: stem, HasStem: true, Next: -1}); v > 0 {
			t.Errorf("the dependent noun %q is let through: %v", stem, v)
		}
	}
}
