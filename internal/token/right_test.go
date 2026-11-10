//ff:func feature=token type=reader control=sequence
//ff:what TestRight — the unit's end and a letter that is no Hangul syllable are an end; a particle's form is a particle; anything else is something else
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestRight(t *testing.T) {
	sc := wordsScan("가신에서 가신고 가신.", layerCost)
	if sc.right(2) != rule.RightParticle || sc.right(7) != rule.RightOther || sc.right(11) != rule.RightDone || sc.right(12) != rule.RightDone || sc.right(4) != rule.RightDone {
		t.Fatal(sc.right(2), sc.right(7), sc.right(11), sc.right(4))
	}
}
