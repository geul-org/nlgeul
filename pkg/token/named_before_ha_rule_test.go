//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNamedBeforeHaRule — the rule is named as the trace prints it, adds bonusOverCount, waits for the word after it, and speaks of a predicate's candidate with the naming ending -기 followed by a particle (서+기+도, 속+기+로), not of the -기 form with nothing after it, a noun 서기 with a particle, another closing, nor a new word with 기
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNamedBeforeHaRule(t *testing.T) {
	r := namedBeforeHaRule()
	if r.name != ruleNamedBeforeHa || r.bonus != bonusOverCount || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "서", Kind: "열쇠:동사", Last: "보조사:도", Marks: "이름:기+보조사:도", Pred: true}, true},
		{rule.Near{Stem: "속", Kind: "열쇠:동사", Last: "부사격:로", Marks: "이름:기+부사격:로", Pred: true}, true},
		{rule.Near{Stem: "좋", Kind: "열쇠:형용사", Last: "보조사:는", Marks: "이름:기+보조사:는", Pred: true}, true},
		{rule.Near{Stem: "서", Kind: "열쇠:동사", Last: "이름:기", Marks: "이름:기", Pred: true}, false},
		{rule.Near{Stem: "서기", Kind: "열쇠:체언", Last: "보조사:도", Marks: "보조사:도"}, false},
		{rule.Near{Stem: "서", Kind: "열쇠:동사", Last: "맺음:고", Marks: "맺음:고", Pred: true}, false},
		{rule.Near{Stem: "압푸", Kind: "", Last: "보조사:도", Marks: "이름:기+보조사:도", Pred: true}, false},
	} {
		if r.about(c.near) != c.want || r.waits(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
