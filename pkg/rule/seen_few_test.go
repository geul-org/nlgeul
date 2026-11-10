//ff:func feature=rule type=rule control=sequence
//ff:what TestSeenFew — one observation is as few as can be, the count named enough is not few, a count between is partly few, an unknown count says nothing
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSeenFew(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "enough", Mask: 10000}}
	one, _ := SeenFew(pieceCtx(Piece{Count: 1}), specs)
	mid, _ := SeenFew(pieceCtx(Piece{Count: 100}), specs)
	many, _ := SeenFew(pieceCtx(Piece{Count: 50000}), specs)
	none, _ := SeenFew(pieceCtx(Piece{}), specs)
	unknown, _ := SeenFew(pieceCtx(Piece{Count: -1}), specs)
	if one != 1 || mid != 0 || many != -1 || none != 1 || unknown != -1 {
		t.Fatal(one, mid, many, none, unknown)
	}
}
