//ff:func feature=rule type=helper control=sequence
//ff:what PieceOf — the piece under judgment (an empty piece when the context has none)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceOf(ctx toulmin.Context) Piece {
	v, _ := ctx.Get(KeyPiece)
	p, _ := v.(Piece)
	return p
}
