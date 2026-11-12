//ff:func feature=token type=reader control=sequence
//ff:what TestJoinPieces — the first reading stands alone; a second adds its name and keeps the higher values
package token

import "testing"

func TestJoinPieces(t *testing.T) {
	a := piece{"이", "주격:이", 1, -1.2, 0, false, ""}
	if joinPieces(piece{}, a) != a {
		t.Fatal("first")
	}
	if got := joinPieces(a, piece{"이", "보격:이", 0.5, -1.0, 0, false, ""}); got != (piece{"이", "주격:이|보격:이", 1, -1.0, 0, false, ""}) {
		t.Fatal(got)
	}
}
