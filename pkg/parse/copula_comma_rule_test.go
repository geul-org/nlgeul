//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCopulaCommaRule — the rule is named as the trace prints it, adds bonusShortTaken, waits for the words after a candidate it speaks of, and speaks only of a candidate with the copula 이 written out whose last piece is the ending 나 (not of the auxiliary particle 이나, nor of 나 after the dropped copula, nor of 이 with another ending)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCopulaCommaRule(t *testing.T) {
	r := copulaCommaRule()
	if r.name != ruleCopulaComma || r.bonus != bonusShortTaken || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Stem: "것", Kind: "의존명사:것", Last: "맺음:나", Marks: "서술격:이+맺음:나"}, true},
		{rule.Near{Stem: "것", Kind: "의존명사:것", Last: "보조사:이나", Marks: "보조사:이나"}, false},
		{rule.Near{Stem: "누구", Last: "맺음:나", Marks: "서술격:이(빠짐)+맺음:나"}, false},
		{rule.Near{Stem: "것", Last: "맺음:다", Marks: "서술격:이+맺음:다"}, false},
		{rule.Near{Stem: "가", Kind: "열쇠:동사", Last: "맺음:나", Marks: "맺음:나"}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
		if got := r.waits(c.near); got != c.want {
			t.Errorf("waits(%+v) = %v", c.near, got)
		}
	}
}
