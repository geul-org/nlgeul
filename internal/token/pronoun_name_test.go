//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestPronounName — 이, 그 and 저 named as determiners are renamed pronouns before particles or the copula; alone they stay determiners; a determiner that is no pronoun, another set's name of the same letters, and a name of other letters keep their names
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestPronounName(t *testing.T) {
	cases := []struct {
		why, name, stem string
		tail            int
		want            string
	}{
		{"그 before a case particle", "관형사:그", "그", rule.TailCase, "대명사:그"},
		{"이 before an auxiliary particle", "관형사:이", "이", rule.TailAux, "대명사:이"},
		{"저 before the copula", "관형사:저", "저", rule.TailCopula, "대명사:저"},
		{"그 alone", "관형사:그", "그", rule.TailBare, "관형사:그"},
		{"이 alone", "관형사:이", "이", rule.TailBare, "관형사:이"},
		{"a determiner that is no pronoun", "관형사:어떤", "어떤", rule.TailCase, "관형사:어떤"},
		{"a determiner of two sets that is no pronoun", "관형사:두|보조 용언:두", "두", rule.TailCase, "관형사:두|보조 용언:두"},
		{"only the determiner's name of two sets", "의존명사:이|관형사:이", "이", rule.TailCase, "의존명사:이|대명사:이"},
		{"a name of 이 that is no determiner", "의존명사:이", "이", rule.TailCase, "의존명사:이"},
		{"a name of other letters than the stem", "관형사:그", "이", rule.TailCase, "관형사:그"},
	}
	for _, c := range cases {
		if got := pronounName(c.name, c.stem, c.tail); got != c.want {
			t.Fatal(c.why, got, c.want)
		}
	}
}
