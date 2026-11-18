//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNaLinkedRuleGraph — the graph of naLinkedRule evaluated by hand on rule.Around: 나다 with a closing stands right after a word that ends in -고, and falls after -지, after a nominal with a particle and at the head of the sentence; after -고 it falls when it ends in the modifying 는 (the pronoun 나 with 는 begins the next clause) and when it has the naming -ㅁ or -기 (the letters of 남다 and of the noun 남)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNaLinkedRuleGraph(t *testing.T) {
	self := func(last string, marks string) rule.Near {
		return rule.Near{Stem: "나", Kind: "열쇠:동사", Last: last, Marks: marks, Pred: true}
	}
	ko := []rule.Near{{Stem: "먹", Kind: "열쇠:동사", Last: "맺음:고", Marks: "맺음:고", Pred: true}}
	r := naLinkedRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"나서 after -고", rule.Around{Self: self("맺음:어서(합침)", "맺음:어서(합침)"), Before: ko}, true},
		{"난 after -고", rule.Around{Self: self("꾸밈:ㄴ", "꾸밈:ㄴ"), Before: ko}, true},
		{"나서 after -지", rule.Around{Self: self("맺음:어서(합침)", "맺음:어서(합침)"), Before: []rule.Near{{Stem: "먹", Kind: "열쇠:동사", Last: "맺음:지", Marks: "맺음:지", Pred: true}}}, false},
		{"나서 after 에", rule.Around{Self: self("맺음:어서(합침)", "맺음:어서(합침)"), Before: []rule.Near{{Stem: "집", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}}}, false},
		{"나서 at the head", rule.Around{Self: self("맺음:어서(합침)", "맺음:어서(합침)")}, false},
		{"나는 after -고", rule.Around{Self: self("꾸밈:는", "꾸밈:는"), Before: ko}, false},
		{"남은 after -고", rule.Around{Self: self("보조사:은", "이름:ㅁ+보조사:은"), Before: ko}, false},
		{"나기 after -고", rule.Around{Self: self("이름:기", "이름:기"), Before: ko}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
