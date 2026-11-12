//ff:func feature=token type=reader control=sequence
//ff:what TestJoinPiecesTypo — two readings of letters written wrong join into a piece that keeps the letters it is written with
package token

import "testing"

func TestJoinPiecesTypo(t *testing.T) {
	a := piece{"(ㄹ)게", "맺음:ㄹ게", 0.6, -2, 4, false, "ㄹ께"}
	b := piece{"(ㄹ)게", "맺음:ㄹ게", 0.5, -1, 4, false, "ㄹ께"}
	if got := joinPieces(a, b); got.wrote != "ㄹ께" || got.names != "맺음:ㄹ게|맺음:ㄹ게" || got.verdict != 0.6 || got.pref != -1 {
		t.Fatal(got)
	}
	if got := joinPieces(piece{}, a); got != a {
		t.Fatal("the first reading stays as it is", got)
	}
}
