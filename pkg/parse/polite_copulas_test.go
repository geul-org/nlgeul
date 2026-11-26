//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPoliteCopulas — the copula stem before a polite slot is read where it is written (학교+이+옵+니다) and not where it is left out; where the stem and the slot's syllable are a word of the word table as written (이온) the tail is not read, and a tail whose piece after the stem is no polite slot is kept
package parse

import "testing"

func TestPoliteCopulas(t *testing.T) {
	sc := wordsScan("학교이옵니다", layerCost)
	polite := sc.polite(sc.closings(6, false, false))
	got := sc.politeCopulas(polite, false)
	if len(got) == 0 {
		t.Fatal("no copula before the slot", polite)
	}
	for _, x := range got {
		if x.start != 2 || x.pieces[0].names != "서술격:이" || x.pieces[0].unseen || x.pieces[1].names != "선어말:옵" {
			t.Fatal(x)
		}
	}
	sc = wordsScan("학교이온", layerCost)
	polite = sc.polite(sc.closings(4, false, false))
	if got = sc.politeCopulas(polite, false); len(polite) == 0 || len(got) == 0 {
		t.Fatal("이온 is no word of this table", polite, got)
	}
	sc.ix.lex["이온"] = word{kinds: kindNominal}
	if got = sc.politeCopulas(polite, false); len(got) != 0 {
		t.Fatal("이온 is a word of the table", got)
	}
	sc.ix.polite = nil
	if got = sc.politeCopulas(polite, false); len(got) == 0 {
		t.Fatal("a piece that is no polite slot is not looked up", got)
	}
	sc = wordsScan("학교옵니다", layerCost)
	polite = sc.polite(sc.closings(5, false, false))
	if got = sc.politeCopulas(polite, false); len(polite) == 0 || len(got) != 0 {
		t.Fatal("a copula stem left out is not read before the slot", polite, got)
	}
	if got = sc.politeCopulas(nil, true); len(got) != 0 {
		t.Fatal(got)
	}
}
