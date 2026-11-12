//ff:func feature=rule type=rule control=sequence
//ff:what TestStemKindNot — with Not it applies when the stem has none of the kinds asked for (the word before is not a stem of those kinds), and not when it has one
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStemKindNot(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "kinds", Mask: 6}, &Not{}}
	none, _ := StemKind(pieceCtx(Piece{Stem: 1}), specs)
	one, _ := StemKind(pieceCtx(Piece{Stem: 3}), specs)
	empty, _ := StemKind(toulmin.NewContext(), specs)
	if none != 1 || one != -1 || empty != 1 {
		t.Fatal(none, one, empty)
	}
}
