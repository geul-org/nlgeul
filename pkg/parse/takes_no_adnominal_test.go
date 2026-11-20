//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestTakesNoAdnominal — no adnominal modifies an adverb (of the word table, a conjunctive one of the closed set), an interjection, a verb, an adjective or an auxiliary verb that ends in another closing or stands bare; one does modify a predicate's modifying or naming form, a nominal, a determiner, a new word and a sign
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTakesNoAdnominal(t *testing.T) {
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "역시", Kind: "열쇠:부사"}, true},
		{rule.Near{Stem: "그리고", Kind: "접속 부사:그리고"}, true},
		{rule.Near{Stem: "아", Kind: "열쇠:홀로"}, true},
		{rule.Near{Stem: "어렵", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다", Pred: true}, true},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:고", Marks: "선어말:았+맺음:고", Pred: true}, true},
		{rule.Near{Stem: "있", Kind: "보조 용언:있", Last: "맺음:다", Marks: "맺음:다", Pred: true}, true},
		{rule.Near{Stem: "가", Kind: "열쇠:동사"}, true},
		{rule.Near{Stem: "크", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true}, false},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "주격:이", Marks: "이름:기+주격:이", Pred: true}, false},
		{rule.Near{Stem: "책", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "새", Kind: "열쇠:관형사"}, false},
		{rule.Near{Stem: "쇼"}, false},
		{rule.Near{Last: "마침표"}, false},
	} {
		if got := takesNoAdnominal(c.near); got != c.want {
			t.Errorf("%+v: %v", c.near, got)
		}
	}
}
