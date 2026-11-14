//ff:func feature=token type=rule control=sequence
//ff:what TestFact — the criterion of a piece's facts with the mask given; rule.PieceIs under it holds when the piece has one of those facts and not when it has none of them, turned round with not; a mask of no fact does not validate
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFact(t *testing.T) {
	f := fact(rule.PieceNum | rule.PieceWhole)
	if f.Of != "piece" || f.Mask != rule.PieceNum|rule.PieceWhole || f.SpecName() != "bits:piece" || f.Validate() != nil {
		t.Fatal(f)
	}
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyPiece, rule.Piece{Form: "명", Num: true})
	if v, _ := rule.PieceIs(ctx, toulmin.Specs{f}); v != 1 {
		t.Fatal("a number right before it is one of the facts", v)
	}
	if v, _ := rule.PieceIs(ctx, toulmin.Specs{fact(rule.PieceWhole | rule.PieceLast)}); v != -1 {
		t.Fatal("none of the facts", v)
	}
	if v, _ := rule.PieceIs(ctx, toulmin.Specs{fact(rule.PieceWhole), not}); v != 1 {
		t.Fatal("turned round with not", v)
	}
	if fact(0).Validate() == nil {
		t.Fatal("a mask of no fact")
	}
}
