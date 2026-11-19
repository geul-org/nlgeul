//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAttachedMarks — markers alone that attach to the word before: a particle, the copula and a closing first (with more after them); not a word with a stem (a noun, a new word), a punctuation mark first, nor no marker at all
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAttachedMarks(t *testing.T) {
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Last: "주격:이", Marks: "주격:이"}, true},
		{rule.Near{Last: "보조사:는", Marks: "보조사:은+보조사:는"}, true},
		{rule.Near{Last: "꾸밈:라는", Marks: "서술격:이+꾸밈:라는"}, true},
		{rule.Near{Last: "맺음:다", Marks: "맺음:다"}, true},
		{rule.Near{Stem: "을", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "을", Last: "주격:이", Marks: "주격:이"}, false},
		{rule.Near{Last: "마침표", Marks: "마침표"}, false},
		{rule.Near{}, false},
	} {
		if got := attachedMarks(c.near); got != c.want {
			t.Errorf("%+v: %v", c.near, got)
		}
	}
}
