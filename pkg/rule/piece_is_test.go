//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestPieceIs — applies when the piece has one of the facts asked for (a lone digit, a whole word with its number, after a modifying closing, a modifying or naming form, a marked new word, the last piece, a stem keeping its vowel, a stem whose word is the whole of its unit, after a determiner); not when it has none of them, or when the context has no piece
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPieceIs(t *testing.T) {
	cases := []struct {
		p    Piece
		mask int
		want float64
	}{
		{Piece{Lone: true}, PieceLone, 1},
		{Piece{}, PieceLone, -1},
		{Piece{Whole: true}, PieceWhole, 1},
		{Piece{}, PieceWhole, -1},
		{Piece{Form: "가르", Adn: true}, PieceAdn, 1},
		{Piece{Form: "가르"}, PieceAdn, -1},
		{Piece{Turned: true}, PieceTurned, 1},
		{Piece{Adn: true}, PieceTurned, -1},
		{Piece{Det: true}, PieceDet, 1},
		{Piece{Adn: true}, PieceDet, -1},
		{Piece{Det: true}, PieceAdn | PieceDet, 1},
		{Piece{Marked: true}, PieceMarked, 1},
		{Piece{}, PieceMarked, -1},
		{Piece{Next: -1}, PieceLast, 1},
		{Piece{Next: 3}, PieceLast, -1},
		{Piece{Keeps: true}, PieceKeeps, 1},
		{Piece{HasStem: true}, PieceKeeps, -1},
		{Piece{Form: "솥", Sole: true}, PieceSole, 1},
		{Piece{Form: "솥", HasStem: true, Next: -1}, PieceSole, -1},
		{Piece{Whole: true}, PieceNum | PieceWhole, 1},
		{Piece{HasStem: true, Adn: true, Next: -1}, PieceLone | PieceMarked, -1},
	}
	for _, c := range cases {
		got, ev := PieceIs(pieceCtx(c.p), toulmin.Specs{&Bits{Of: "piece", Mask: c.mask}})
		if got != c.want || ev != nil {
			t.Errorf("PieceIs(%+v, %b) = %v, %v", c.p, c.mask, got, ev)
		}
	}
	if empty, _ := PieceIs(toulmin.NewContext(), toulmin.Specs{&Bits{Of: "piece", Mask: PieceStem | PieceLast}}); empty != -1 {
		t.Error("empty context", empty)
	}
}
