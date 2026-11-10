//ff:func feature=token type=reader control=sequence
//ff:what TestTailOf — a closing that holds the copula, one with its own kind of tail, a case particle, another particle, a closing only verbs take, and anything else of a predicate tail
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestTailOf(t *testing.T) {
	if tailOf(&entry{class: classClosing, holds: true}) != rule.TailCopula || tailOf(&entry{class: classClosing, tail: rule.TailAdj}) != rule.TailAdj || tailOf(&entry{class: classParticle, kase: rule.CaseObj}) != rule.TailCase || tailOf(&entry{class: classParticle}) != rule.TailAux || tailOf(&entry{class: classClosing, name: "맺음:지만"}) != rule.TailPred || tailOf(&entry{class: classSlot}) != rule.TailPred || tailOf(&entry{class: classClosing, name: "맺음:ㄴ다"}) != rule.TailVerb {
		t.Fatal("tailOf")
	}
}
