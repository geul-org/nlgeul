//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceNoAdn — applies when the word before the piece does not end in a modifying closing
package rule

import "testing"

func TestPieceNoAdn(t *testing.T) {
	yes, _ := PieceNoAdn(pieceCtx(Piece{}), nil)
	no, _ := PieceNoAdn(pieceCtx(Piece{Adn: true}), nil)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
