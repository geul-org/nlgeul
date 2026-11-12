//ff:func feature=token type=reader control=sequence
//ff:what TestWholeDet — where a determiner ends as a word of no markers is kept (어떤), and not where another word does (사과); a word that only stands alone (아이고) grows only the kept readings whose last word neither ends in a modifying closing nor is a determiner, while a noun grows them all
package token

import "testing"

func TestWholeDet(t *testing.T) {
	sc := wordsScan("어떤사과", layerCost)
	best := [][]*path{{{verdict: 1}}, nil, nil, nil, nil}
	best[2] = sc.whole(nil, best, 2, false)
	if len(best[2]) != 1 || !best[2][0].last.det || !sc.detEnds[2] {
		t.Fatal("어떤, a determiner", best[2], sc.detEnds)
	}
	if got := sc.whole(nil, best, 4, false); len(got) == 0 || got[0].last.det || sc.detEnds[4] {
		t.Fatal("사과, no determiner", got, sc.detEnds)
	}
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}, {verdict: 1, last: step{det: true}}}
	alone := wordsScan("먹을아이고", layerCost)
	alone.adnEnds[2], alone.detEnds[2] = true, true
	got := alone.whole(nil, [][]*path{nil, nil, kept, nil, nil, nil}, 5, false)
	if len(got) != 1 || got[0].prev != kept[0] || got[0].last.sig != "아이고[열쇠:홀로]" {
		t.Fatal("아이고 right after what modifies a nominal", got)
	}
	noun := wordsScan("먹을사과", layerCost)
	noun.adnEnds[2], noun.detEnds[2] = true, true
	if got := noun.whole(nil, [][]*path{nil, nil, kept, nil, nil}, 4, false); len(got) != 3 {
		t.Fatal("사과 grows every reading", got)
	}
}
