//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestGluedBoundRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of a candidate that reads its stem as a noun of the word table whose letters are one of the words given with a particle as its first marker (자+가, 성+과), not with the copula first (자+(이)+라면), not the bare noun (자), not a longer noun (자가), not another name (자[의존명사:자]+가) and not another noun (측+은)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGluedBoundRule(t *testing.T) {
	r := gluedBoundRule([]string{"자", "성"})
	if r.name != ruleGluedBound || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "자", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}, true},
		{rule.Near{Stem: "성", Kind: "열쇠:체언", Last: "부사격:과", Marks: "부사격:과"}, true},
		{rule.Near{Stem: "자", Kind: "열쇠:체언", Last: "맺음:라면", Marks: "서술격:이(빠짐)+맺음:라면"}, false},
		{rule.Near{Stem: "자", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "자가", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "자", Kind: "의존명사:자", Last: "주격:가", Marks: "주격:가"}, false},
		{rule.Near{Stem: "측", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
