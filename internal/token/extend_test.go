//ff:func feature=token type=reader control=sequence
//ff:what TestExtend — a tail grows every reading kept before its stem, from each place before it
package token

import "testing"

func TestExtend(t *testing.T) {
	sc := scanOf("곰이")
	best := [][]*path{{{verdict: 1}}, nil, nil}
	if got := sc.extend(nil, best, tail{start: 2, pieces: []piece{{text: "x"}}}, 2); len(got) != 1 {
		t.Fatal("a place with no reading grown", got)
	}
	got := sc.extend(nil, best, tail{start: 1, pieces: []piece{{"이", "주격:이", 1, -1, 0, false}}}, 2)
	if len(got) != 1 || got[0].sig != "곰+이[주격:이]" {
		t.Fatal(got)
	}
}
