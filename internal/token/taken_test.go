//ff:func feature=token type=reader control=sequence
//ff:what TestTaken — a key of two kinds has the count of the kind that takes what follows; the key's greatest count when no kind of it takes that, and when kinds are not looked at
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestTaken(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if ix.taken("작", rule.TailCase) != 400 || ix.taken("작", rule.TailPred) != 7000 || ix.taken("고양이", rule.TailPred) != 1000 || ix.taken("없는말", rule.TailBare) != 0 {
		t.Fatal(ix.taken("작", rule.TailCase), ix.taken("작", rule.TailPred))
	}
	if wordsScan("x", layerKnown).ix.taken("작", rule.TailCase) != 7000 {
		t.Fatal("kinds are not looked at in the first layer")
	}
}
