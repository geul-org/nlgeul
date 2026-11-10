//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceNoNumber — applies when no number stands right before the piece
package rule

import "testing"

func TestPieceNoNumber(t *testing.T) {
	yes, _ := PieceNoNumber(pieceCtx(Piece{}), nil)
	no, _ := PieceNoNumber(pieceCtx(Piece{Num: true}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
