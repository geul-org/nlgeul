//ff:func feature=token type=rule control=sequence
//ff:what TestAddAfter — a piece is kept by its last letter and stands only right after a number and not with letters running on after it; after a lone Sino-Korean digit with no word table it does not stand, and when the number and the piece together are a word it stands lower; a greater number has neither of the last two
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddAfter(t *testing.T) {
	ix := &index{after: map[rune][]*entry{}}
	ix.addAfter("단위:개월", "개월", placeUnit)
	e := ix.after['월'][0]
	yes, _ := verdictOf(e.g, rule.Piece{Num: true, Right: rule.RightDone}, "")
	on, _ := verdictOf(e.g, rule.Piece{Num: true, Right: rule.RightOther}, "")
	no, _ := verdictOf(e.g, rule.Piece{}, "")
	lone, _ := verdictOf(e.g, rule.Piece{Num: true, Right: rule.RightDone, Lone: true}, "")
	whole, _ := verdictOf(e.g, rule.Piece{Num: true, Right: rule.RightDone, Whole: true}, "")
	if e.form != "개월" || e.slot != placeUnit || yes <= 0 || no > 0 || on > 0 || ix.graphs != 1 || ix.edges != 4 {
		t.Fatal(e, yes, no)
	}
	if lone > 0 || whole <= 0 || whole >= yes {
		t.Fatal(lone, whole)
	}
	ix.addAfter(bigName+"만", "만", placeBig)
	if big, _ := verdictOf(ix.after['만'][0].g, rule.Piece{Num: true, Right: rule.RightDone, Lone: true, Whole: true}, ""); big != yes || ix.edges != 6 {
		t.Fatal(big, ix.edges)
	}
}
