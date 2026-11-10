//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceNoStem — applies when no stem stands before the piece
package rule

import "testing"

func TestPieceNoStem(t *testing.T) {
	no, _ := PieceNoStem(pieceCtx(Piece{}), nil)
	yes, _ := PieceNoStem(pieceCtx(Piece{HasStem: true}), nil)
	if no != 1 || yes != -1 {
		t.Fatal(no, yes)
	}
}
