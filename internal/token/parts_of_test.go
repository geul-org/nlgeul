//ff:func feature=token type=reader control=sequence
//ff:what TestPartsOf — a piece with no parts stays; with parts each has its verdict and tail, the first its preference
package token

import "testing"

func TestPartsOf(t *testing.T) {
	pc := piece{"죠", "맺음:죠", 0.5, -2, 4, false, ""}
	if got := partsOf(pc, nil); len(got) != 1 || got[0] != pc {
		t.Fatal(got)
	}
	got := partsOf(pc, splits["맺음:죠"])
	if len(got) != 2 || got[0] != (piece{"지", jiName, 0.5, -2, 4, false, ""}) || got[1] != (piece{"(요)", "보조사:요", 0.5, 0, 4, false, ""}) {
		t.Fatal(got)
	}
}
