//ff:func feature=token type=rule control=sequence
//ff:what TestPartVerdict — a known word stands as it does with what follows; a run of Latin letters and digits as such a run; anything else not at all
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPartVerdict(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	if ix.partVerdict("입력", rule.TailCase) <= 0 || ix.partVerdict("입력", rule.TailPred) != 0 || ix.partVerdict("ab1", rule.TailBare) != ix.run || ix.partVerdict("없는말", rule.TailBare) != 0 {
		t.Fatal(ix.partVerdict("입력", rule.TailCase), ix.partVerdict("ab1", rule.TailBare))
	}
}
