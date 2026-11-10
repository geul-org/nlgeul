//ff:func feature=rule type=helper control=sequence
//ff:what pieceCtx — a rule context for one piece
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func pieceCtx(p Piece) toulmin.Context {
	ctx := toulmin.NewContext()
	ctx.Set(KeyPiece, p)
	return ctx
}
