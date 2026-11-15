//ff:func feature=token type=reader control=sequence
//ff:what TestWholeAtChange — letters no table lists whose new word only the place where the script changes inside them turns down (crossesChange) do not grow the whole new word but grow the other ways their stem stands (제1-1, the closed-set 제 before the atom 1-1: the atom's names 차례, 범위 and 빼기, as they stand with no word table, and no new word); with no other way to stand they grow none (K팝 with 팝 a noun of the table); a new word the place does not turn down stands (K팝 with no 팝 in the table); letters the new word's graph turns down for what unlistedPiece tells grow none (A씨 with 씨 a title of the table, written apart)
package token

import "testing"

func TestWholeAtChange(t *testing.T) {
	order := wordsScan("제1-1", layerCost)
	got := order.wholeAt(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil}, 0, 4, false)
	if len(got) != 3 || got[0].sig != "제1-1[차례]" || got[1].sig != "제1-1[범위]" || got[2].sig != "제1-1[빼기]" || got[0].verdict != 1 {
		t.Fatal("제1-1: the atom's names, and no new word", got)
	}
	bare := scanOf("제1-1").wholeAt(nil, [][]*path{{{verdict: 1}}, nil, nil, nil, nil}, 0, 4, false)
	if len(bare) != 3 || bare[0].sig != got[0].sig || bare[1].sig != got[1].sig || bare[2].sig != got[2].sig {
		t.Fatal("제1-1 with no word table", bare)
	}

	pop := wordsScan("K팝", layerCost)
	pop.ix.lex["팝"] = word{kinds: kindNominal, count: 500, by: [6]int{500}}
	if got := pop.wholeAt(nil, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2, false); len(got) != 0 {
		t.Fatal("K팝 with 팝 a noun of the table: the whole new word only, turned down", got)
	}
	if got := wordsScan("K팝", layerCost).wholeAt(nil, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2, false); len(got) != 1 || got[0].sig != "K팝" || got[0].verdict <= 0 {
		t.Fatal("K팝 with no 팝 in the table: the whole new word", got)
	}

	title := wordsScan("A씨", layerCost)
	title.ix.lex["씨"] = word{kinds: kindNominal, count: 500, by: [6]int{500}}
	if got := title.wholeAt(nil, [][]*path{{{verdict: 1}}, nil, nil}, 0, 2, false); len(got) != 0 {
		t.Fatal("A씨, written apart", got)
	}
}
