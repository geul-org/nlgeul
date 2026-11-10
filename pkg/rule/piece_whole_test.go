//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceWhole — applies when the number and the piece together are a word
package rule

import "testing"

func TestPieceWhole(t *testing.T) {
	yes, _ := PieceWhole(pieceCtx(Piece{Whole: true}), nil)
	no, _ := PieceWhole(pieceCtx(Piece{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
