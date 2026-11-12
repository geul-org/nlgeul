//ff:func feature=rule type=rule control=sequence
//ff:what TestStemKind — applies when the stem has one of the kinds asked for; not when it has none, or when the context has no piece
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStemKind(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "kinds", Mask: 6}, &Claim{Name: "kind", Concludes: "stem"}}
	none, nv := StemKind(pieceCtx(Piece{Stem: 1}), specs)
	one, ov := StemKind(pieceCtx(Piece{Stem: 3}), specs)
	all, _ := StemKind(pieceCtx(Piece{Stem: 6}), specs)
	empty, _ := StemKind(toulmin.NewContext(), specs)
	if none != -1 || one != 1 || all != 1 || empty != -1 || nv != nil || ov != nil {
		t.Fatal(none, one, all, empty, nv, ov)
	}
}
