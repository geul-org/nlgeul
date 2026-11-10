//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceRunsOn — applies when letters run on right after the piece
package rule

import "testing"

func TestPieceRunsOn(t *testing.T) {
	yes, _ := PieceRunsOn(pieceCtx(Piece{Right: RightOther}), nil)
	no, _ := PieceRunsOn(pieceCtx(Piece{Right: RightDone}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
