//ff:func feature=token type=reader control=sequence
//ff:what TestKeptAt — the candidates before a sign between letters are those kept there with the words of no markers that end there (a new word stands though no change of script shows its end), worked out once for a place
package token

import "testing"

func TestKeptAt(t *testing.T) {
	sc := scanOf("구름·돌")
	best := [][]*path{{{verdict: 1}}, nil, nil, nil, nil}
	got := sc.keptAt(best, 2)
	if len(got) != 1 || got[0].sig != "구름" {
		t.Fatal(got)
	}
	if again := sc.keptAt(nil, 2); len(again) != 1 || again[0] != got[0] {
		t.Fatal("worked out once", again)
	}
}
