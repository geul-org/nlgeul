//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestLinkBeforeNegRuleGraph — the graph of linkBeforeNegRule evaluated by hand on rule.Around: the -지 form stands right before 않다, 못하다 and 말다, and falls before another verb, before a noun, before a sign and at the end of the sentence
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLinkBeforeNegRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:지", Marks: "맺음:지", Pred: true}
	verb := func(stem string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: "열쇠:동사", Last: "맺음:다", Marks: "맺음:다", Pred: true}}
	}
	r := linkBeforeNegRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before 않다", rule.Around{Self: self, After: verb("않")}, true},
		{"before 못하다", rule.Around{Self: self, After: verb("못하")}, true},
		{"before 말다", rule.Around{Self: self, After: verb("말")}, true},
		{"before another verb", rule.Around{Self: self, After: verb("먹")}, false},
		{"before a noun", rule.Around{Self: self, After: []rule.Near{{Stem: "책", Kind: "열쇠:체언"}}}, false},
		{"before a sign", rule.Around{Self: self, After: []rule.Near{{Kind: "", Last: "마침표"}}}, false},
		{"at the end", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
