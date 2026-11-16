//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNounHeadedRuleGraph — the graph of nounHeadedRule evaluated by hand on rule.Around: the noun 시 stands at the head of a sentence, right after an opening sign and right before a middle dot whatever is before it, falls after a noun, after a new word and after a closing sign with no middle dot after it (before a word, before a comma, at the end); the noun 뿐 at the head falls right before 아니다 (뿐만 아니라)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounHeadedRuleGraph(t *testing.T) {
	si := rule.Near{Stem: "시", Kind: "열쇠:체언", Last: "보조사:는", Marks: "보조사:는"}
	dotted := rule.Near{Stem: "시", Kind: "열쇠:체언"}
	ppun := rule.Near{Stem: "뿐", Kind: "열쇠:체언", Last: "보조사:만", Marks: "보조사:만"}
	r := nounHeadedRule(boundNext)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"at the head", rule.Around{Self: si}, true},
		{"after an opening sign", rule.Around{Self: si, Before: []rule.Near{{Stem: "“", Kind: "열기:1"}, {Stem: "그", Kind: "열쇠:체언"}}}, true},
		{"after a noun", rule.Around{Self: si, Before: []rule.Near{{Stem: "인구", Kind: "열쇠:체언"}}}, false},
		{"after a new word", rule.Around{Self: si, Before: []rule.Near{{Stem: "ᄒᆞᆫ"}}}, false},
		{"after a closing sign", rule.Around{Self: si, Before: []rule.Near{{Stem: "”", Kind: "닫기:1"}}}, false},
		{"after a noun before a middle dot", rule.Around{Self: dotted, Before: []rule.Near{{Stem: "인구", Kind: "열쇠:체언"}}, After: []rule.Near{{Last: middleDot, Marks: middleDot}, {Stem: "군", Kind: "열쇠:체언"}}}, true},
		{"after a noun before a word", rule.Around{Self: dotted, Before: []rule.Near{{Stem: "인구", Kind: "열쇠:체언"}}, After: []rule.Near{{Stem: "군", Kind: "열쇠:체언"}}}, false},
		{"after a noun before a comma", rule.Around{Self: dotted, Before: []rule.Near{{Stem: "인구", Kind: "열쇠:체언"}}, After: []rule.Near{{Last: "쉼표", Marks: "쉼표"}}}, false},
		{"뿐 at the head before 아니다", rule.Around{Self: ppun, After: []rule.Near{{Stem: "아니", Kind: "열쇠:형용사", Last: "맺음:라", Pred: true}}}, false},
		{"뿐 at the head before another word", rule.Around{Self: ppun, After: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "맺음:다", Pred: true}}}, true},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
