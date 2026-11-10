//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceRight — applies when what stands right after the piece is of the kind asked about
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceRight(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "right", Mask: RightParticle}}
	yes, _ := PieceRight(pieceCtx(Piece{Right: RightParticle}), specs)
	no, _ := PieceRight(pieceCtx(Piece{Right: RightOther}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
