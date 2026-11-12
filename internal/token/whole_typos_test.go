//ff:func feature=token type=reader control=sequence
//ff:what TestWholeTypos — the letters from b to j that are a wrong way to write a noun of the word table are read as the right form with no marker, from every candidate kept at b, with the letters kept; a bound stem with a word written wrong is read as the right form; not from a place no candidate reaches, not for a verb stem written wrong (맞기), not for letters that are not misspelt, and not after a candidate that ends in a modifying closing when the right form only stands alone
package token

import "testing"

func TestWholeTypos(t *testing.T) {
	sc := typoScan("역활", layerCost)
	best := [][]*path{{{verdict: 1}}, nil, nil}
	got := sc.wholeTypos(nil, best, 0, 2)
	if len(got) != 1 || got[0].sig != "역할[열쇠:체언,오타정정:역활]" {
		t.Fatal(got)
	}
	if got := sc.wholeTypos(nil, [][]*path{nil, nil, nil}, 0, 2); len(got) != 0 {
		t.Fatal("no candidate reaches b", got)
	}
	if got := typoScan("맞기", layerCost).wholeTypos(nil, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2); len(got) != 0 {
		t.Fatal("a verb stem written wrong stands with no marker nowhere", got)
	}
	if got := typoScan("역할", layerCost).wholeTypos(nil, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2); len(got) != 0 {
		t.Fatal("letters that are not misspelt", got)
	}
	two := typoScan("가역활", layerCost)
	if got := two.wholeTypos(nil, [][]*path{{{verdict: 1}}, {{verdict: 1, sig: "가"}}, nil, nil}, 1, 3); len(got) != 1 || got[0].sig != "가 역할[열쇠:체언,오타정정:역활]" {
		t.Fatal(got)
	}
	bound := typoScan("입력/역활", layerCost)
	if got := bound.wholeTypos(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil, nil}, 0, 5); len(got) != 1 || got[0].sig != "입력/역할[대비:/,오타정정:입력/역활]" {
		t.Fatal(got)
	}
	ix, err := newIndex(Options{Words: "testdata/words-typo-alone.tsv", Layers: layerCost})
	if err != nil {
		t.Fatal(err)
	}
	alone := &scan{ix: ix, s: []rune("아이구")}
	alone.prepare()
	alone.adnEnds[0], alone.detEnds[0] = true, true
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}, {verdict: 1, last: step{det: true}}}
	got = alone.wholeTypos(nil, [][]*path{kept, nil, nil, nil}, 0, 3)
	if len(got) != 1 || got[0].prev != kept[0] || got[0].last.sig != "아이고[열쇠:홀로,오타정정:아이구]" {
		t.Fatal("a right form that only stands alone grows only the candidates that end in no modifying closing and no determiner", got)
	}
}
