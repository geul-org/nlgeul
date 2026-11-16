//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNominalBoundRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of what particledNoun takes with the words given (측+은), not of the copula's reading (내+(이)+지) nor of a longer noun (측은)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNominalBoundRule(t *testing.T) {
	r := nominalBoundRule([]string{"측", "내"})
	if r.name != ruleNominalBound || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "측", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}, true},
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Last: "맺음:지", Marks: "서술격:이(빠짐)+맺음:지"}, false},
		{rule.Near{Stem: "측은", Kind: "열쇠:체언"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
