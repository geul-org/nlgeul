//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceBase — applies when the syllable under a fused piece is not in the list
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceBase(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("bases", []string{"하", "되"})}
	out, _ := PieceBase(pieceCtx(Piece{Base: "크"}), specs)
	in, _ := PieceBase(pieceCtx(Piece{Base: "하"}), specs)
	if out != 1 || in != -1 {
		t.Fatal(out, in)
	}
}
