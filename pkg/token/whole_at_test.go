//ff:func feature=token type=reader control=sequence
//ff:what TestWholeAt — the function is worked through by the net over the sentences of the tests and made-up long units, and each is read to at least one stretch; where only some places may begin a word (starts), letters from a place that is not one of them grow nothing and the candidates given come back as they are, while from a place that is one of them they grow (K팝 with no 팝 in the table: the whole new word)
package token

import "testing"

func TestWholeAt(t *testing.T) {
	driveAll(t)

	given := []*path{{sig: "앞", verdict: 1}}
	shut := wordsScan("K팝", layerCost)
	shut.starts = map[int]bool{1: true}
	if got := shut.wholeAt(given, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2, false); len(got) != 1 || got[0] != given[0] {
		t.Fatal("K팝 from a place no word may begin at: the candidates given, as they are", got)
	}
	open := wordsScan("K팝", layerCost)
	open.starts = map[int]bool{0: true}
	if got := open.wholeAt(nil, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2, false); len(got) != 1 || got[0].sig != "K팝" || got[0].verdict <= 0 {
		t.Fatal("K팝 from a place a word may begin at: the whole new word", got)
	}
}
