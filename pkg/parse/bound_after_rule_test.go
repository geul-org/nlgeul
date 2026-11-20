//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestBoundAfterRule — the rule is named as the trace prints it, adds bonusOverWhole, waits for nothing, and speaks of a candidate that reads its stem as a dependent noun alone (not as a noun of the same letters, not a verb)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundAfterRule(t *testing.T) {
	r := boundAfterRule(boundNominal)
	if r.name != ruleBoundAfter || r.bonus != bonusOverWhole || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "수", Kind: "의존명사:수", Last: "보조사:도", Marks: "보조사:도"}, true},
		{rule.Near{Stem: "거", Kind: "의존명사:거", Last: "보조사:는", Marks: "보조사:는"}, true},
		{rule.Near{Stem: "수도", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "걸", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true, Marks: "꾸밈:ㄴ"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
