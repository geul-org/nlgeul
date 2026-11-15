//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAloneRule — the rule is named as the trace prints it, adds bonusSentence, and speaks of (and waits for the word after) a candidate that reads the letters of a closed-set bound noun or determiner as another word (시[열쇠:체언], 이[열쇠:체언]+가), not of the bound noun or the determiner itself (시[의존명사:시], 이[관형사:이], 이[열쇠:관형사]), nor of a word of no closed set or a candidate with no stem
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAloneRule(t *testing.T) {
	r := aloneRule()
	if r.name != ruleAlone || r.bonus != bonusSentence || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "시", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "이", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}, true},
		{rule.Near{Stem: "수", Kind: ""}, true},
		{rule.Near{Stem: "시", Kind: "의존명사:시"}, false},
		{rule.Near{Stem: "이", Kind: "관형사:이"}, false},
		{rule.Near{Stem: "이", Kind: "열쇠:관형사"}, false},
		{rule.Near{Stem: "책", Kind: "열쇠:체언"}, false},
		{rule.Near{Last: "주격:이", Marks: "주격:이"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
		if got := r.waits(c.near); got != c.want {
			t.Errorf("waits(%+v) = %v", c.near, got)
		}
	}
}
