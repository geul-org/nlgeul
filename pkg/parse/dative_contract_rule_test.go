//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDativeContractRule — the rule is named as the trace prints it, adds bonusSameLetters, and speaks of and waits for a candidate that reads 내, 네 or 제 with the particle 게 named 부사격:에게 (not 나 with 에게, not the verb 내 with the ending 게, not 내 with another particle)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDativeContractRule(t *testing.T) {
	r := dativeContractRule()
	if r.name != ruleDativeContract || r.bonus != bonusSameLetters || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Last: "부사격:에게", Marks: "부사격:에게"}, true},
		{rule.Near{Stem: "네", Kind: "열쇠:체언", Last: "보조사:는", Marks: "부사격:에게+보조사:는"}, true},
		{rule.Near{Stem: "제", Kind: "열쇠:체언", Last: "부사격:에게", Marks: "부사격:에게"}, true},
		{rule.Near{Stem: "나", Kind: "열쇠:체언", Last: "부사격:에게", Marks: "부사격:에게"}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:동사", Last: "맺음:게", Marks: "맺음:게", Pred: true}, false},
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
