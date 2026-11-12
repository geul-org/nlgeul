//ff:func feature=token type=reader control=sequence
//ff:what TestGrowAlone — each kept reading is grown by the word that stands alone (a number written in binary has two readings); a word that is a determiner marks where it ends; the readings that end in a modifying closing or a determiner are not grown by a word that only stands alone
package token

import "testing"

func TestGrowAlone(t *testing.T) {
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}, {verdict: 1, last: step{det: true}}}
	alone := wordsScan("먹을아이고", layerCost)
	alone.adnEnds[2], alone.detEnds[2] = true, true
	e := alone.stepsOf(2, 5, tail{start: 5})[0]
	got := alone.growAlone(nil, kept, e, 2, 5)
	if len(got) != 1 || got[0].prev != kept[0] || got[0].last.sig != "아이고[열쇠:홀로]" {
		t.Fatal("아이고 right after what modifies a nominal", got)
	}
	noun := wordsScan("먹을사과", layerCost)
	noun.adnEnds[2], noun.detEnds[2] = true, true
	e = noun.stepsOf(2, 4, tail{start: 4})[0]
	if got := noun.growAlone(nil, kept, e, 2, 4); len(got) != 3 {
		t.Fatal("사과 grows every reading", got)
	}
	det := wordsScan("어떤사과", layerCost)
	e = det.stepsOf(0, 2, tail{start: 2})[0]
	if got := det.growAlone(nil, kept[:1], e, 0, 2); len(got) != 1 || !det.detEnds[2] {
		t.Fatal("어떤, a determiner, is marked where it ends", got, det.detEnds)
	}
	e = det.stepsOf(2, 4, tail{start: 4})[0]
	if got := det.growAlone(nil, kept[:1], e, 2, 4); len(got) != 1 || det.detEnds[4] {
		t.Fatal("사과, no determiner", got, det.detEnds)
	}
	bin := scanOf("1010")
	e = bin.stepsOf(0, 4, tail{start: 4})[0]
	if got := bin.growAlone(nil, kept[:1], e, 0, 4); len(got) != 2 || got[0].last.sig != "1010[수]" || got[1].last.sig != "1010[수:2진법]" {
		t.Fatal("a number in binary has two readings", got)
	}
	if got := bin.growAlone(nil, nil, e, 0, 4); len(got) != 0 {
		t.Fatal("nothing kept", got)
	}
}
