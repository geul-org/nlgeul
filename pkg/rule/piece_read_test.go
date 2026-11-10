//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceRead — always applies; the evidence is the form
package rule

import "testing"

func TestPieceRead(t *testing.T) {
	if v, ev := PieceRead(pieceCtx(Piece{Form: "은"}), nil); v != 1 || ev != "은" {
		t.Fatal(v, ev)
	}
}
