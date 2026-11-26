//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestNamedGate — the graph the word before the suffix 이 of a called name must pass lets a nominal of the word table and a new word through, and turns down each dependent noun of the closed set and a word that stands as an adverb, a determiner, a predicate or an interjection
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNamedGate(t *testing.T) {
	e := namedGate()
	if e.name != nameGate || e.g == nil {
		t.Fatal(e)
	}
	for _, kinds := range []int{0, kindNominal} {
		if v := verdict(e.g, rule.Piece{Form: "영숙", HasStem: true, Next: -1, Stem: kinds}); v <= 0 {
			t.Errorf("kinds %d are turned down: %v", kinds, v)
		}
	}
	for _, kinds := range []int{kindAdverb, kindDeterminer, kindVerb, kindAdj, kindAlone, kindNominal | kindAdverb} {
		if v := verdict(e.g, rule.Piece{Form: "가장", HasStem: true, Next: -1, Stem: kinds}); v > 0 {
			t.Errorf("kinds %d are let through: %v", kinds, v)
		}
	}
	bound := ClosedWords()[boundSet]
	if len(bound) == 0 {
		t.Fatal("no dependent nouns in the closed set")
	}
	for _, stem := range bound {
		if v := verdict(e.g, rule.Piece{Form: stem, HasStem: true, Next: -1, Stem: kindNominal}); v > 0 {
			t.Errorf("the dependent noun %q is let through: %v", stem, v)
		}
	}
}
