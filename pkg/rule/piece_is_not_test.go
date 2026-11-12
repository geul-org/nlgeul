//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestPieceIsNot — with Not it applies when the piece has none of the facts asked for (no stem before it, no number right before it, not after a modifying closing, a piece follows), and not when it has one
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceIsNot(t *testing.T) {
	cases := []struct {
		p    Piece
		mask int
		want float64
	}{
		{Piece{}, PieceStem, 1},
		{Piece{HasStem: true}, PieceStem, -1},
		{Piece{}, PieceNum, 1},
		{Piece{Num: true}, PieceNum, -1},
		{Piece{}, PieceAdn, 1},
		{Piece{Adn: true}, PieceAdn, -1},
		{Piece{Next: 3}, PieceLast, 1},
		{Piece{Next: -1}, PieceLast, -1},
		{Piece{Num: true}, PieceStem | PieceAdn, 1},
		{Piece{Adn: true}, PieceStem | PieceAdn, -1},
	}
	for _, c := range cases {
		got, ev := PieceIs(pieceCtx(c.p), toulmin.Specs{&Bits{Of: "piece", Mask: c.mask}, &Not{}, &Claim{Name: "a", Concludes: "b"}})
		if got != c.want || ev != nil {
			t.Errorf("PieceIs(%+v, %b, not) = %v, %v", c.p, c.mask, got, ev)
		}
	}
}
