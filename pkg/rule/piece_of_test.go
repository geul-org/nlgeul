//ff:func feature=rule type=helper control=sequence
//ff:what TestPieceOf — the piece set in the context comes back; without one, an empty piece
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceOf(t *testing.T) {
	if PieceOf(toulmin.NewContext()).Form != "" || PieceOf(pieceCtx(Piece{Form: "은"})).Form != "은" {
		t.Fatal("PieceOf")
	}
}
