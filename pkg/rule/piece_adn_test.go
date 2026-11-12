//ff:func feature=rule type=rule control=sequence
//ff:what TestPieceAdn — applies when the word before the piece ends in a modifying closing, not when it does not or when the context has no piece
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceAdn(t *testing.T) {
	yes, yv := PieceAdn(pieceCtx(Piece{Form: "가르", Adn: true}), nil)
	no, nv := PieceAdn(pieceCtx(Piece{Form: "가르"}), nil)
	none, _ := PieceAdn(toulmin.NewContext(), nil)
	if yes != 1 || no != -1 || none != -1 || yv != nil || nv != nil {
		t.Fatal(yes, no, none, yv, nv)
	}
}
