//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPossessiveRule — the rule is named as the trace prints it, adds bonusOverCount, and speaks of and waits for a candidate that reads a pronoun of the closed set (내, 네, 제) as the nominal of the table with no particle; not the interjection of those letters, not one with a particle, not another nominal
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPossessiveRule(t *testing.T) {
	r := possessiveRule()
	if r.name != rulePossessive || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "네", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "내", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "제", Kind: "열쇠:체언", Glued: true}, true},
		{rule.Near{Stem: "네", Kind: "열쇠:홀로"}, false},
		{rule.Near{Stem: "네", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}, false},
		{rule.Near{Stem: "집", Kind: "열쇠:체언"}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about/waits(%+v), want %v", c.near, c.want)
		}
	}
}
