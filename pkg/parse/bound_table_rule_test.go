//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestBoundTableRule — the rule is named as the trace prints it, adds bonusOverWhole, waits for nothing, and speaks of a candidate that reads its stem as a noun of the word table whose letters are one of the words given, with a marker after it (리+가, 바+를), not the bare noun (리), not a longer noun (리가), not the dependent noun of the closed set (수[의존명사:수]+도) and not another noun (측+은)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundTableRule(t *testing.T) {
	r := boundTableRule([]string{"리", "바"})
	if r.name != ruleBoundTable || r.bonus != bonusOverWhole || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "리", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}, true},
		{rule.Near{Stem: "바", Kind: "열쇠:체언", Last: "목적격:를", Marks: "목적격:를"}, true},
		{rule.Near{Stem: "리", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "리가", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "수", Kind: "의존명사:수", Last: "보조사:도", Marks: "보조사:도"}, false},
		{rule.Near{Stem: "측", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
