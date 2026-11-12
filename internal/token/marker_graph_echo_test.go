//ff:func feature=token type=rule control=sequence
//ff:what TestMarkerGraphEcho — the graph of the particle 에 holds a rebuttal that it stands right on a stem with no final consonant before 요, of force forceClash: with the fact the verdict is 2/(1+forceClash)-1 and the reading is not rejected; without it 1; a particle that is not listed (echoes) has none
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphEcho(t *testing.T) {
	e := entry{name: "부사격:에", form: "에", class: classParticle, slot: rule.LinkAdv}
	markerGraph(&e, nil, true)
	open := rule.Piece{Form: "에", HasStem: true, Prev: '이', Slot: rule.LinkAdv, Next: rule.LinkYo, YoOpen: true}
	v, _ := verdictOf(e.g, open, "고양이에", false)
	if math.Abs(v-(2/(1+forceClash)-1)) > 1e-9 {
		t.Fatal("rebutted, not rejected", v)
	}
	shut := open
	shut.YoOpen = false
	if v, _ = verdictOf(e.g, shut, "집에", false); v != 1 {
		t.Fatal("without the fact", v)
	}
	other := entry{name: "부사격:에서", form: "에서", class: classParticle, slot: rule.LinkAdv}
	markerGraph(&other, nil, true)
	open.Form = "에서"
	if v, _ = verdictOf(other.g, open, "고양이에서", false); v != 1 {
		t.Fatal("a particle not listed", v)
	}
}
