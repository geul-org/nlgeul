//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceParticle — applies when a particle stands right after the piece
package rule

import "testing"

func TestPieceParticle(t *testing.T) {
	yes, _ := PieceParticle(pieceCtx(Piece{Right: RightParticle}), nil)
	no, _ := PieceParticle(pieceCtx(Piece{Right: RightDone}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
