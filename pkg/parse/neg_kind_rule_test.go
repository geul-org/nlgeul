//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestNegKindRule — the rule is named as the trace prints it, is a rebuttal of the force forceUngrammatical that adds nothing and reads nothing after the candidate, and speaks of a candidate that reads 못하 as the kind given (the adjective: not the verb 못하, not 않, not a new word of the letters)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNegKindRule(t *testing.T) {
	r := negKindRule("열쇠:형용사", []string{"열쇠:동사"}, whyNotVerbBefore)
	if r.name != ruleNegKind || r.force != forceUngrammatical || r.bonus != 0 || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "못하", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다", Pred: true}, true},
		{rule.Near{Stem: "못하", Kind: "열쇠:동사", Last: "맺음:다", Marks: "맺음:다", Pred: true}, false},
		{rule.Near{Stem: "않", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다", Pred: true}, false},
		{rule.Near{Stem: "못하", Kind: "", Last: "맺음:다", Marks: "맺음:다", Pred: true}, false},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
