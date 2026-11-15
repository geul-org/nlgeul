//ff:func feature=rule type=helper control=sequence
//ff:what TestPieceFactsUnitWord — the field UnitWord sets the flag PieceUnitWord alone, and adds up with the other flags (a number before it, the prefix 제)
package rule

import "testing"

func TestPieceFactsUnitWord(t *testing.T) {
	if got := pieceFacts(Piece{UnitWord: true}); got != PieceUnitWord {
		t.Fatalf("%b, want %b", got, PieceUnitWord)
	}
	if got := pieceFacts(Piece{Num: true, Order: true, UnitWord: true}); got != PieceNum|PieceOrder|PieceUnitWord {
		t.Fatalf("%b", got)
	}
	if pieceFacts(Piece{Form: "년간"})&PieceUnitWord != 0 {
		t.Fatal("a stem's letters alone do not set it")
	}
}
