//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestNounUnheldRuleGraph — the graph of nounUnheldRule evaluated by hand on rule.Around: the noun 시 stands after a particle, a linking closing, an adverb and a comma, and falls after a modifying closing, the adnominal 의, a determiner of the closed set, a nominal, a suffix, the 여 of a rough number, a closing sign, at the head of a sentence and right after an opening sign
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounUnheldRuleGraph(t *testing.T) {
	si := rule.Near{Stem: "시", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}
	before := func(n rule.Near) rule.Around { return rule.Around{Self: si, Before: []rule.Near{n}} }
	r := nounUnheldRule(unheldBound)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a particle", before(rule.Near{Stem: "사람", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}), true},
		{"after a linking closing", before(rule.Near{Stem: "따르", Kind: "열쇠:동사", Last: "맺음:아", Pred: true, Marks: "맺음:아"}), true},
		{"after an adverb", before(rule.Near{Stem: "또는", Kind: "열쇠:부사"}), true},
		{"after a comma", before(rule.Near{Stem: "특히", Kind: "열쇠:부사", Last: "쉼표", Marks: "쉼표"}), true},
		{"after a modifying closing", before(rule.Near{Stem: "쓰", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true, Marks: "꾸밈:ㄹ"}), false},
		{"after 의", before(rule.Near{Stem: "사람", Kind: "열쇠:체언", Last: "관형격:의", Marks: "관형격:의"}), false},
		{"after a determiner", before(rule.Near{Stem: "그", Kind: "관형사:그"}), false},
		{"after a nominal", before(rule.Near{Stem: "인구", Kind: "열쇠:체언", Nominal: true}), false},
		{"after a suffix", before(rule.Near{Stem: "사람", Kind: "열쇠:체언", Last: "접미:들", Marks: "접미:들"}), false},
		{"after a rough number", before(rule.Near{Stem: "5", Kind: "수", Last: "어림:여", Marks: "단위:분+어림:여"}), false},
		{"after a closing sign", before(rule.Near{Stem: ")", Kind: "닫기:1"}), false},
		{"at the head", rule.Around{Self: si}, false},
		{"after an opening sign", before(rule.Near{Stem: "“", Kind: "열기:1"}), false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
