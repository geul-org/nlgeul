//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceMarked — applies when the letters of a whole new word after a boundary read as markers
package rule

import "testing"

func TestPieceMarked(t *testing.T) {
	yes, _ := PieceMarked(pieceCtx(Piece{Marked: true}), nil)
	no, _ := PieceMarked(pieceCtx(Piece{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
