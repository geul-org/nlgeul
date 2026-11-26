//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestDroppedAfterRunRule — the rule is named as the trace prints it, adds bonusShortTaken, and speaks only of a candidate with the dropped copula stem whose stem is a run of Latin letters and digits (not of a noun of the table, nor of the closing read right on the run)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDroppedAfterRunRule(t *testing.T) {
	r := droppedAfterRunRule()
	if r.name != ruleDropRun || r.bonus != bonusShortTaken || r.graph == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "abc", Kind: runName, Last: "맺음:다", Marks: "서술격:이(빠짐)+맺음:다"}, true},
		{rule.Near{Stem: "abc", Kind: runName, Last: "맺음:다", Marks: "맺음:다"}, false},
		{rule.Near{Stem: "학교", Kind: "열쇠:체언", Last: "맺음:다", Marks: "서술격:이(빠짐)+맺음:다"}, false},
		{rule.Near{Stem: "3", Kind: "수", Last: "맺음:다", Marks: "서술격:이(빠짐)+맺음:다"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
}
