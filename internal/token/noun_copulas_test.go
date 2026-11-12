//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestNounCopulas — of the tails given, each one read after 테 gives its candidate in order and the others none (no closing, another letter); no tails give none; the predicate tails of a unit (closings) carry 터+(이) before 니 (갈테니), before ㄴ데 fused under 텐 (올텐데) and before 야 (할테야) where a word that ends in a modifying closing ends before the letter, and not where none does
package token

import "testing"

func TestNounCopulas(t *testing.T) {
	sc := scanOf("갈테니")
	sc.adnEnds[1] = true
	ni := []piece{{text: "니", names: "맺음:니"}}
	given := []tail{
		{start: 2, pieces: ni, close: &entry{name: "맺음:니"}},
		{start: 2, pieces: ni},
		{start: 3, base: '니', pieces: []piece{{text: "(어)", names: "맺음:어(합침)"}}, close: &entry{name: "맺음:어(합침)"}},
		{start: 2, pieces: ni, close: &entry{name: "맺음:니"}, noYo: true},
	}
	got := sc.nounCopulas(given)
	if len(got) != 2 || got[0].noYo || !got[1].noYo || got[0].close != given[0].close || got[1].close != given[3].close || got[0].label != "줄임:테" || got[1].label != "줄임:테" {
		t.Fatal(got)
	}
	if got = sc.nounCopulas(nil); got != nil {
		t.Fatal("no tails", got)
	}
	for in, want := range map[string]string{"갈테니": "니", "올텐데": "(ㄴ)데", "할테야": "야"} {
		n, read := 0, scanOf(in)
		read.adnEnds[1] = true
		for _, c := range read.closings(3, false, false) {
			if c.label == "줄임:테" && c.start == 2 && c.shown == "터" && len(c.pieces) == 2 && c.pieces[0].names == copulaName && c.pieces[1].text == want {
				n++
			}
		}
		bare := 0
		for _, c := range scanOf(in).closings(3, false, false) {
			if c.label == "줄임:테" {
				bare++
			}
		}
		if n != 1 || bare != 0 {
			t.Errorf("%s: %d after a modifying closing, %d with none", in, n, bare)
		}
	}
}
