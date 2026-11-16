//ff:func feature=token type=rule control=sequence
//ff:what TestStemVerdict — a long, often seen noun before a particle is fully a known word; a rarely seen one and a one-syllable one less so; a kind that does not take the tail not at all; the same facts are judged once, and another stem's facts are judged apart, not answered from the first's verdict
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStemVerdict(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	cat := rule.Piece{Form: "고양이", Tail: rule.TailCase, Count: 50000}
	full := ix.stemVerdict(kindNominal, cat)
	few := ix.stemVerdict(kindNominal, rule.Piece{Form: "고양이", Tail: rule.TailCase, Count: 100})
	short := ix.stemVerdict(kindNominal, rule.Piece{Form: "물", Tail: rule.TailCase, Count: 50000})
	part := 1 - math.Log10(100)/math.Log10(float64(ix.seenEnough()))
	if full != 1 || math.Abs(few-(2/(1+forceFew*part)-1)) > 1e-9 || math.Abs(short-(2/(1+forceShort)-1)) > 1e-9 {
		t.Fatal(full, few, short)
	}
	if ix.stemVerdict(kindNominal, rule.Piece{Form: "고양이", Tail: rule.TailPred, Count: 50000}) != -1 || ix.stemVerdict(kindAdj, rule.Piece{Form: "작", Tail: rule.TailVerb, Count: 50000}) != -1 || ix.stemVerdict(kindVerb, rule.Piece{Form: "마시", Tail: rule.TailVerb, Count: 50000}) != 1 {
		t.Fatal("tails")
	}
	n := len(ix.memo)
	if ix.stemVerdict(kindNominal, cat) != full || len(ix.memo) != n {
		t.Fatal("not remembered")
	}
	if ix.stemVerdict(kindNominal, rule.Piece{Form: "마당이", Tail: rule.TailCase, Count: 50000}) != 1 || len(ix.memo) != n+1 {
		t.Fatal("another stem is judged apart", len(ix.memo), n)
	}
}
