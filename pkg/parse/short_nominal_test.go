//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestShortNominal — a candidate whose stem is one syllable named as a noun of the table, a dependent noun or a pronoun, with a particle or the copula 이 first among its markers written out after it, is such a noun; a longer stem, a stem alone, a closing or the dropped copula first, a verb's stem, a particle put under the stem's syllable (letters no longer than the stem) and a candidate with no stem are not
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestShortNominal(t *testing.T) {
	for _, c := range []struct {
		name string
		near rule.Near
		want bool
	}{
		{"noun and 의", rule.Near{Stem: "자", Kind: "열쇠:체언", Marks: "관형격:의", Text: "자의"}, true},
		{"noun and two particles", rule.Near{Stem: "자", Kind: "열쇠:체언", Marks: "부사격:와+관형격:의", Text: "자와의"}, true},
		{"noun and the copula", rule.Near{Stem: "일", Kind: "열쇠:체언", Marks: "서술격:이+꾸밈:ㄴ", Text: "일인"}, true},
		{"dependent noun", rule.Near{Stem: "간", Kind: "의존명사:간", Marks: "부사격:에", Text: "간에"}, true},
		{"pronoun", rule.Near{Stem: "나", Kind: "대명사:나", Marks: "주격:가", Text: "나가"}, true},
		{"two syllables", rule.Near{Stem: "자의", Kind: "열쇠:체언", Marks: "주격:가", Text: "자의가"}, false},
		{"stem alone", rule.Near{Stem: "자", Kind: "열쇠:체언", Text: "자"}, false},
		{"closing first", rule.Near{Stem: "치", Kind: "열쇠:체언", Marks: "꾸밈:는", Text: "치는"}, false},
		{"dropped copula", rule.Near{Stem: "지", Kind: "열쇠:체언", Marks: "서술격:이(빠짐)+맺음:지", Text: "지지"}, false},
		{"verb", rule.Near{Stem: "치", Kind: "열쇠:동사", Marks: "꾸밈:는", Text: "치는"}, false},
		{"particle under the syllable", rule.Near{Stem: "저", Kind: "열쇠:체언", Marks: "보조사:는", Text: "전"}, false},
		{"no stem", rule.Near{Marks: "관형격:의", Text: "의"}, false},
	} {
		if got := shortNominal(c.near); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
