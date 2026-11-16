//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestParticledNoun — a noun of the word table whose stem is one of the words given and whose first marker is a particle is one (측+은, 측+의+는), and not when the copula comes first (내+(이)+지), when it has no marker (측), when its stem is not among the words (측은), nor when it bears another name (측[의존명사:측]+은)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestParticledNoun(t *testing.T) {
	words := []string{"측", "내"}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "측", Kind: "열쇠:체언", Last: "보조사:은", Marks: "보조사:은"}, true},
		{rule.Near{Stem: "측", Kind: "열쇠:체언", Last: "보조사:는", Marks: "관형격:의+보조사:는"}, true},
		{rule.Near{Stem: "내", Kind: "열쇠:체언", Last: "맺음:지", Marks: "서술격:이(빠짐)+맺음:지"}, false},
		{rule.Near{Stem: "측", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "측은", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "측", Kind: "의존명사:측", Last: "보조사:은", Marks: "보조사:은"}, false},
	} {
		if got := particledNoun(c.near, words); got != c.want {
			t.Errorf("particledNoun(%+v) = %v", c.near, got)
		}
	}
}
