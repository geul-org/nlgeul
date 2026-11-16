//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCopulaEndRule — the rule is named as the trace prints it, adds bonusShortTaken, waits for the words after a candidate it speaks of, and speaks only of a candidate whose last piece is the ending 야 or 라고 (not of 야 read as the vocative or the auxiliary particle, nor of 라고 with a particle after it, nor of a word with no marker)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCopulaEndRule(t *testing.T) {
	r := copulaEndRule()
	if r.name != ruleCopulaEnd || r.bonus != bonusShortTaken || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "차", Last: "맺음:야", Marks: "서술격:이(빠짐)+맺음:야"}, true},
		{rule.Near{Stem: "책", Last: "맺음:야", Marks: "서술격:이+맺음:야"}, true},
		{rule.Near{Stem: "책", Last: "맺음:라고", Marks: "서술격:이+맺음:라고"}, true},
		{rule.Near{Stem: "책", Last: "보조사:는", Marks: "서술격:이+맺음:라고+보조사:는"}, false},
		{rule.Near{Stem: "차", Last: "호격:야", Marks: "호격:야"}, false},
		{rule.Near{Stem: "차", Last: "보조사:야", Marks: "보조사:야"}, false},
		{rule.Near{Stem: "차", Kind: "열쇠:체언"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
		if got := r.waits(c.near); got != c.want {
			t.Errorf("waits(%+v) = %v", c.near, got)
		}
	}
}
