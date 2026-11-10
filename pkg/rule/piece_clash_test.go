//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceClash — applies when a piece that begins with 어/아 does not fit the vowel before it; not when it fits, when nothing stands before, or for a piece with no letters of its own
package rule

import "testing"

func TestPieceClash(t *testing.T) {
	yes, _ := PieceClash(pieceCtx(Piece{Form: "아서", Prev: '먹'}), nil)
	no, _ := PieceClash(pieceCtx(Piece{Form: "어서", Prev: '먹'}), nil)
	none, _ := PieceClash(pieceCtx(Piece{Form: "아"}), nil)
	bare, _ := PieceClash(pieceCtx(Piece{Prev: '먹'}), nil)
	if yes != 1 || no != -1 || none != -1 || bare != -1 {
		t.Fatal(yes, no, none, bare)
	}
}
