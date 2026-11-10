//ff:func feature=token type=rule control=sequence
//ff:what TestAddAfter — a piece is kept by its last letter and stands only right after a number
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddAfter(t *testing.T) {
	ix := &index{after: map[rune][]*entry{}}
	ix.addAfter("단위:개월", "개월", placeUnit)
	e := ix.after['월'][0]
	yes, _ := verdictOf(e.g, rule.Piece{Num: true}, "")
	no, _ := verdictOf(e.g, rule.Piece{}, "")
	if e.form != "개월" || e.slot != placeUnit || yes <= 0 || no > 0 || ix.graphs != 1 || ix.edges != 1 {
		t.Fatal(e, yes, no)
	}
}
