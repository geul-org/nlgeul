//ff:func feature=token type=rule control=sequence
//ff:what TestGateOf — the graph passes a word that has one of the kinds and not one that has none
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestGateOf(t *testing.T) {
	e := gateOf("앞말:x", kindVerb|kindAdj)
	yes, _ := verdictOf(e.g, rule.Piece{Stem: kindVerb | kindNominal}, "")
	no, _ := verdictOf(e.g, rule.Piece{Stem: kindNominal}, "")
	if e.name != "앞말:x" || yes <= 0 || no > 0 {
		t.Fatal(yes, no)
	}
}
