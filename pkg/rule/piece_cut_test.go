//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceCut — applies when what follows is neither the unit's end nor a particle
package rule

import "testing"

func TestPieceCut(t *testing.T) {
	yes, _ := PieceCut(pieceCtx(Piece{Right: RightOther}), nil)
	no, _ := PieceCut(pieceCtx(Piece{Right: RightParticle}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
