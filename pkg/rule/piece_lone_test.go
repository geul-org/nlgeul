//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceLone — applies when a lone Sino-Korean digit stands before the piece with no word table to tell
package rule

import "testing"

func TestPieceLone(t *testing.T) {
	yes, _ := PieceLone(pieceCtx(Piece{Lone: true}), nil)
	no, _ := PieceLone(pieceCtx(Piece{}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
