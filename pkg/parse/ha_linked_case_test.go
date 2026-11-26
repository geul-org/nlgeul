//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestHaLinkedRuleCase — the case row 하지만 under haLinkedRule's name holds after a -기 form and not after -어야 (하 with the ending 지만 is preferred), on rule.Around and as a stream after -어야
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestHaLinkedRuleCase(t *testing.T) {
	self := rule.Near{Stem: "하", Text: "하지만", Kind: "열쇠:동사", Last: "맺음:지만", Marks: "맺음:지만", Pred: true}
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
		{"after -기는", rule.Around{Self: self, Before: []rule.Near{before("있", "이름:기+보조사:는", "보조사:는")}}, false},
		{"after -고", rule.Around{Self: self, Before: []rule.Near{before("먹", "맺음:고", "맺음:고")}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
	verb := "하[열쇠:동사]+지만[맺음:지만]"
	if got, _ := ruledFirst(t, "words-ha-linked.tsv", "돈이 있어야 하지만", "하지만", r); got != verb {
		t.Errorf("하지만 after -어야 is %s, want %s", got, verb)
	}
}
