//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPriorNounRule — the rule is named as the trace prints it, adds bonusOverCount, does not wait, and speaks of a candidate that reads 전 as the noun of the word table (not 전 the determiner, not another noun)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPriorNounRule(t *testing.T) {
	r := priorNounRule(nil)
	if r.name != rulePriorNoun || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "전", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "전", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}, true},
		{rule.Near{Stem: "전", Kind: "열쇠:관형사"}, false},
		{rule.Near{Stem: "후", Kind: "열쇠:체언"}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
