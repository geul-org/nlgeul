//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDegreeUnitRule — the rule of a sentence is named as the trace prints it, adds bonusSameLetters and waits for nothing, and speaks only of a candidate whose stem is a number and whose last piece is the unit 도 (not 도 as a particle, not the unit with a particle after it, not another unit, not a stem that is no number)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDegreeUnitRule(t *testing.T) {
	r := degreeUnitRule()
	if r.name != ruleDegreeUnit || r.graph == nil || r.bonus != bonusSameLetters || r.waits != nil {
		t.Fatal(r.name, r.bonus, r.waits != nil)
	}
	cases := []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "10", Kind: "수", Last: "단위:도", Marks: "단위:도"}, true},
		{rule.Near{Stem: "10", Kind: "수", Last: "보조사:도", Marks: "보조사:도"}, false},
		{rule.Near{Stem: "10", Kind: "수", Last: "부사격:로", Marks: "단위:도+부사격:로"}, false},
		{rule.Near{Stem: "10", Kind: "수", Last: "단위:개", Marks: "단위:개"}, false},
		{rule.Near{Stem: "구조", Kind: "열쇠:체언", Last: "단위:도", Marks: "단위:도"}, false},
		{rule.Near{}, false},
	}
	for _, c := range cases {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v; want %v", c.near, got, c.want)
		}
	}
}
