//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestOutsideNounRule — the rule is named as the trace prints it, adds bonusOverCount, and speaks of and waits for a candidate that reads 밖 as a noun of the word table with the adverbial 에 as its first marker (밖+에, 밖+에+는), not the particle 밖에, nor 밖 with another particle, nor another noun with 에
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestOutsideNounRule(t *testing.T) {
	r := outsideNounRule()
	if r.name != ruleOutsideNoun || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "밖", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}, true},
		{rule.Near{Stem: "밖", Kind: "열쇠:체언", Last: "보조사:는", Marks: "부사격:에+보조사:는"}, true},
		{rule.Near{Stem: "", Kind: "", Last: "보조사:밖에", Marks: "보조사:밖에"}, false},
		{rule.Near{Stem: "밖", Kind: "열쇠:체언", Last: "부사격:에서", Marks: "부사격:에서"}, false},
		{rule.Near{Stem: "밖", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "집", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
