//ff:func feature=token type=rule control=sequence
//ff:what TestStemVerdict — a long, often seen noun before a particle is fully a known word; a rarely seen one and a one-syllable one less so; a kind that does not take the tail not at all; the same facts are judged once
package token

import (
	"math"
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestStemVerdict(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	full := ix.stemVerdict(kindNominal, rule.TailCase, 50000, "고양이")
	few := ix.stemVerdict(kindNominal, rule.TailCase, 100, "고양이")
	short := ix.stemVerdict(kindNominal, rule.TailCase, 50000, "물")
	if full != 1 || math.Abs(few-(2/(1+forceFew*0.5)-1)) > 1e-9 || math.Abs(short-(2/(1+forceShort)-1)) > 1e-9 {
		t.Fatal(full, few, short)
	}
	if ix.stemVerdict(kindNominal, rule.TailPred, 50000, "고양이") != -1 || ix.stemVerdict(kindAdj, rule.TailVerb, 50000, "작") != -1 || ix.stemVerdict(kindVerb, rule.TailVerb, 50000, "마시") != 1 {
		t.Fatal("tails")
	}
	n := len(ix.seen)
	if ix.stemVerdict(kindNominal, rule.TailCase, 50000, "마당이") != 1 || len(ix.seen) != n {
		t.Fatal("not remembered")
	}
}
