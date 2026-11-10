//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceLast — applies when nothing follows the piece
package rule

import "testing"

func TestPieceLast(t *testing.T) {
	last, _ := PieceLast(pieceCtx(Piece{Next: -1}), nil)
	more, _ := PieceLast(pieceCtx(Piece{Next: 3}), nil)
	if last != 1 || more != -1 {
		t.Fatal(last, more)
	}
}
