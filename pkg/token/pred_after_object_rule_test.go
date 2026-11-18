//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPredAfterObjectRule — the rule is named as the trace prints it, adds bonusOverCount, does not wait, and speaks of a candidate of a verb or an adjective of the word table with a closing (not one with no closing, not a noun with the copula, not an auxiliary verb, not a new word), but not of one whose word is a negative adverb's letters (안, 못, and 아니 read as 알다 with -니)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPredAfterObjectRule(t *testing.T) {
	r := predAfterObjectRule()
	if r.name != rulePredAfterObject || r.bonus != bonusOverCount || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "보", Kind: "열쇠:동사", Last: "맺음:니", Marks: "맺음:니", Pred: true, Text: "보니"}, true},
		{rule.Near{Stem: "검", Kind: "열쇠:형용사", Last: "꾸밈:은", Marks: "꾸밈:은", Pred: true, Text: "검은"}, true},
		{rule.Near{Stem: "보니", Kind: "열쇠:체언", Text: "보니"}, false},
		{rule.Near{Stem: "보", Kind: "열쇠:동사", Text: "보"}, false},
		{rule.Near{Stem: "거리", Kind: "열쇠:체언", Last: "맺음:다", Marks: "서술격:이(빠짐)+맺음:다", Pred: true, Text: "거리다"}, false},
		{rule.Near{Stem: "주", Kind: "보조 용언:주", Last: "맺음:다", Marks: "맺음:다", Pred: true, Text: "주다"}, false},
		{rule.Near{Stem: "보", Kind: "", Last: "맺음:니", Marks: "맺음:니", Pred: true, Text: "보니"}, false},
		{rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true, Text: "안"}, false},
		{rule.Near{Stem: "모", Kind: "열쇠:동사", Last: "맺음:ㅅ", Marks: "맺음:ㅅ", Pred: true, Text: "못"}, false},
		{rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "맺음:니", Marks: "맺음:니", Pred: true, Text: "아니"}, false},
		{rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "맺음:니", Marks: "맺음:니", Pred: true, Text: "알니"}, true},
	} {
		if r.about(c.near) != c.want {
			t.Errorf("about(%+v), want %v", c.near, c.want)
		}
	}
}
