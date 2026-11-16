//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestApartParticleRule — the rule is named as the trace prints it, adds bonusNewTaken (what the rebuttal against a new word takes from lambda, about 7.16), reads nothing after the candidate, and speaks of a word read as markers alone whose first is a particle (는, 이라고, 에+도), not of the demonstrative's letters 이, nor of a word with a stem, nor of a closing alone
package token

import (
	"math"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestApartParticleRule(t *testing.T) {
	r := apartParticleRule()
	if r.name != ruleApartParticle || r.bonus != bonusNewTaken || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	if math.Abs(bonusNewTaken-7.16) > 0.01 {
		t.Errorf("bonusNewTaken = %v", bonusNewTaken)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Last: "보조사:는", Marks: "보조사:는", Text: "는"}, true},
		{rule.Near{Last: "부사격:이라고", Marks: "부사격:이라고", Text: "이라고"}, true},
		{rule.Near{Last: "보조사:도", Marks: "부사격:에+보조사:도", Text: "에도"}, true},
		{rule.Near{Last: "주격:이", Marks: "주격:이", Text: "이"}, false},
		{rule.Near{Stem: "나무", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는", Text: "나무는"}, false},
		{rule.Near{Last: "맺음:며", Marks: "맺음:며", Text: "며"}, false},
		{rule.Near{Text: "는"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
