//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestAdnNoun — a noun of the word table of two syllables or more that stands alone and ends in 인 or 일 is such a noun; one that ends otherwise, one of one syllable, one with a marker, and a candidate under another name or with no name are not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdnNoun(t *testing.T) {
	for _, c := range []struct {
		name string
		near rule.Near
		want bool
	}{
		{"ends in 인", rule.Near{Stem: "감정인", Kind: "열쇠:체언"}, true},
		{"ends in 일", rule.Near{Stem: "시행일", Kind: "열쇠:체언"}, true},
		{"two syllables", rule.Near{Stem: "법인", Kind: "열쇠:체언"}, true},
		{"ends otherwise", rule.Near{Stem: "감정", Kind: "열쇠:체언"}, false},
		{"one syllable", rule.Near{Stem: "일", Kind: "열쇠:체언"}, false},
		{"with a marker", rule.Near{Stem: "감정인", Kind: "열쇠:체언", Last: "주격:이", Marks: "주격:이"}, false},
		{"the copula reading", rule.Near{Stem: "감정", Kind: "열쇠:체언", Last: "꾸밈:ㄴ", Marks: "서술격:이+꾸밈:ㄴ", Pred: true}, false},
		{"a new word", rule.Near{Stem: "소외인"}, false},
		{"a dependent noun", rule.Near{Stem: "따위일", Kind: "의존명사:따위일"}, false},
	} {
		if got := adnNoun(c.near); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
