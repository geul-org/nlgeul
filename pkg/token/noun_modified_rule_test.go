//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNounModifiedRule — the rule is named as the trace prints it, adds bonusOverCount, waits for nothing, and speaks of a candidate that reads its stem as a noun of the word table alone, not one whose letters are a dependent noun's or a determiner's of the closed sets (수, 때, 이) and not one read under another name
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounModifiedRule(t *testing.T) {
	r := nounModifiedRule(modifiedSkip)
	if r.name != ruleNounModified || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "보고", Kind: "열쇠:체언"}, true},
		{rule.Near{Stem: "떼", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}, true},
		{rule.Near{Stem: "수", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "때", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "이", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "보", Kind: "열쇠:동사", Last: "맺음:고", Pred: true, Marks: "맺음:고"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
