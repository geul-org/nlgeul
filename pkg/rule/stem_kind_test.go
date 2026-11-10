//ff:func feature=rule type=rule control=sequence
//ff:what TestStemKind — applies when the word before has none of the kinds asked for
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStemKind(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "kinds", Mask: 6}}
	yes, _ := StemKind(pieceCtx(Piece{Stem: 1}), specs)
	no, _ := StemKind(pieceCtx(Piece{Stem: 3}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
