//ff:func feature=rule type=rule control=sequence
//ff:what TestStemTail — applies when what follows the stem is none of what its kind takes
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestStemTail(t *testing.T) {
	noun := toulmin.Specs{&Bits{Of: "tail", Mask: TailBare | TailAux | TailCase}}
	bad, _ := StemTail(pieceCtx(Piece{Tail: TailPred}), noun)
	good, _ := StemTail(pieceCtx(Piece{Tail: TailCase}), noun)
	if bad != 1 || good != -1 {
		t.Fatal(bad, good)
	}
}
