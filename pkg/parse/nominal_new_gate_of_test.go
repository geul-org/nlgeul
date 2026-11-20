//ff:func feature=parse type=rule control=sequence
//ff:what TestNominalNewGateOf — the gate of a nominal of the table or a new word: a stem the table gives as a nominal passes, and so does a stem the tables know no kind of (a new word, with the reason of the exception's exception), a stem the table gives only as a predicate does not, and one that is a nominal and a predicate too does
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNominalNewGateOf(t *testing.T) {
	e := nominalNewGateOf(nominalNewGate)
	if e.name != nominalNewGate || e.g == nil {
		t.Fatal(e)
	}
	ok := func(kinds int) bool {
		return verdict(e.g, rule.Piece{Form: "뚜루", HasStem: true, Next: -1, Stem: kinds}) > 0
	}
	if !ok(kindNominal) || !ok(0) || ok(kindVerb) || ok(kindVerb|kindAdj) || !ok(kindNominal|kindVerb) {
		t.Fatal(ok(kindNominal), ok(0), ok(kindVerb), ok(kindVerb|kindAdj), ok(kindNominal|kindVerb))
	}
}
