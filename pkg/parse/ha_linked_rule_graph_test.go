//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestHaLinkedRuleGraph — the graph of haLinkedRule evaluated by hand on rule.Around: 하다 stands right after -어야 (written as it is or fused) and right after a word with the naming ending -기 and a particle (-기는, -기로), and falls after -고, after a noun with 는, after a -기 form two words before and at the head of the sentence
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHaLinkedRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "하", Kind: "열쇠:동사", Last: "맺음:나", Marks: "맺음:나", Pred: true}
	before := func(stem, marks, last string) rule.Near {
		return rule.Near{Stem: stem, Kind: "열쇠:동사", Last: last, Marks: marks, Pred: true}
	}
	r := haLinkedRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after -어야", rule.Around{Self: self, Before: []rule.Near{before("갖추", "맺음:어야", "맺음:어야")}}, true},
		{"after -어야 fused", rule.Around{Self: self, Before: []rule.Near{before("하", "맺음:어야(합침)", "맺음:어야(합침)")}}, true},
		{"after -기는", rule.Around{Self: self, Before: []rule.Near{before("있", "이름:기+보조사:는", "보조사:는")}}, true},
		{"after -기로", rule.Around{Self: self, Before: []rule.Near{before("가", "이름:기+부사격:로", "부사격:로")}}, true},
		{"after -고", rule.Around{Self: self, Before: []rule.Near{before("먹", "맺음:고", "맺음:고")}}, false},
		{"after a noun with 는", rule.Around{Self: self, Before: []rule.Near{{Stem: "돈", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}}}, false},
		{"after -기는 two words before", rule.Around{Self: self, Before: []rule.Near{{Stem: "다시", Kind: "열쇠:부사"}, before("있", "이름:기+보조사:는", "보조사:는")}}, false},
		{"at the head", rule.Around{Self: self}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
