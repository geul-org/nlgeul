//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestOutsideNounRuleGraph — the graph of outsideNounRule evaluated by hand on rule.Around: 밖+에 stands before a verb that is no word of negation and at the end of the sentence, and falls when the word right after it is 없다, 모르다 or the negative adverb 안, or when the word after that is 않다 (밖에 남지 않았다)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestOutsideNounRuleGraph(t *testing.T) {
	self := rule.Near{Stem: "밖", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}
	word := func(stem, kind string) rule.Near {
		return rule.Near{Stem: stem, Kind: kind, Pred: kind != "열쇠:부사"}
	}
	r := outsideNounRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a verb", rule.Around{Self: self, After: []rule.Near{word("나가", "열쇠:동사")}}, true},
		{"at the end", rule.Around{Self: self}, true},
		{"before 없다", rule.Around{Self: self, After: []rule.Near{word("없", "열쇠:형용사")}}, false},
		{"before 모르다", rule.Around{Self: self, After: []rule.Near{word("모르", "열쇠:동사")}}, false},
		{"before 안", rule.Around{Self: self, After: []rule.Near{word("안", "열쇠:부사"), word("되", "열쇠:동사")}}, false},
		{"before 남지 않다", rule.Around{Self: self, After: []rule.Near{word("남", "열쇠:동사"), word("않", "열쇠:동사")}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
