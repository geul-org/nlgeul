//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestPieceFacts — each yes/no field of a piece sets its own flag (Next below 0 sets PieceLast), an empty piece has none, and the flags add up
package rule

import "testing"

func TestPieceFacts(t *testing.T) {
	cases := []struct {
		p    Piece
		want int
	}{
		{Piece{}, 0},
		{Piece{Form: "가르", Prev: '먹', Slot: 2, Next: 3, Base: "하", Before: "꾸밈:는", After: "맺음:다", Stem: 6}, 0},
		{Piece{HasStem: true}, PieceStem},
		{Piece{Num: true}, PieceNum},
		{Piece{Lone: true}, PieceLone},
		{Piece{Whole: true}, PieceWhole},
		{Piece{Adn: true}, PieceAdn},
		{Piece{Turned: true}, PieceTurned},
		{Piece{Marked: true}, PieceMarked},
		{Piece{Next: -1}, PieceLast},
		{Piece{Keeps: true}, PieceKeeps},
		{Piece{Sole: true}, PieceSole},
		{Piece{Det: true}, PieceDet},
		{Piece{Closed: true}, PieceClosed},
		{Piece{Same: true}, PieceSame},
		{Piece{Misspelt: true}, PieceMisspelt},
		{Piece{YoOpen: true}, PieceYoOpen},
		{Piece{Adn: true, Det: true}, PieceAdn | PieceDet},
		{Piece{HasStem: true, Adn: true, Next: -1}, PieceStem | PieceAdn | PieceLast},
		{Piece{Keeps: true, Sole: true}, PieceKeeps | PieceSole},
		{Piece{HasStem: true, Num: true, Lone: true, Whole: true, Adn: true, Turned: true, Marked: true, Next: -1, Keeps: true, Sole: true, Det: true, Closed: true, Same: true, Misspelt: true, YoOpen: true},
			PieceStem | PieceNum | PieceLone | PieceWhole | PieceAdn | PieceTurned | PieceMarked | PieceLast | PieceKeeps | PieceSole | PieceDet | PieceClosed | PieceSame | PieceMisspelt | PieceYoOpen},
	}
	for _, c := range cases {
		if got := pieceFacts(c.p); got != c.want {
			t.Errorf("pieceFacts(%+v) = %b, want %b", c.p, got, c.want)
		}
	}
}
