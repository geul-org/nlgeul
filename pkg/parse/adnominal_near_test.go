//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestAdnominalNear — a candidate as the rules of a sentence read it is an adnominal when its last piece is a modifying closing or when it is a determiner (관형사:이, 열쇠:관형사) with no marker; not a determiner's letters with a marker, a noun, a predicate with another closing, nor an empty one
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAdnominalNear(t *testing.T) {
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "크", Kind: "열쇠:형용사", Last: "꾸밈:ㄴ", Marks: "꾸밈:ㄴ", Pred: true}, true},
		{rule.Near{Stem: "이", Kind: "관형사:이"}, true},
		{rule.Near{Stem: "새", Kind: "열쇠:관형사"}, true},
		{rule.Near{Stem: "이", Kind: "관형사:이", Last: "보조사:는", Marks: "보조사:는"}, false},
		{rule.Near{Stem: "책", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "크", Kind: "열쇠:형용사", Last: "맺음:다", Marks: "맺음:다", Pred: true}, false},
		{rule.Near{}, false},
	} {
		if got := adnominalNear(c.near); got != c.want {
			t.Errorf("%+v: %v", c.near, got)
		}
	}
}
