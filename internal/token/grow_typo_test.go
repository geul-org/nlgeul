//ff:func feature=token type=reader control=sequence
//ff:what TestGrowTypo — each kept candidate is grown by a word written wrong; a word of no verdict grows none; the candidates that end in a modifying closing or a determiner are not grown by a right form that only stands alone; a determiner marks where it ends
package token

import "testing"

func TestGrowTypo(t *testing.T) {
	kept := []*path{{verdict: 1}, {verdict: 1, last: step{adn: true}}, {verdict: 1, last: step{det: true}}}
	sc := typoScan("역활", layerCost)
	tl := sc.typoReads(nil, "역활", tail{start: 2})[0]
	e := sc.stepsOf(0, 2, tl)[0]
	got := sc.growTypo(nil, kept, tl, e, 0, 2)
	if len(got) != 3 || got[0].prev != kept[0] || got[2].prev != kept[2] || got[0].last.sig != "역할[열쇠:체언,오타정정:역활]" || sc.detEnds[2] {
		t.Fatal("a nominal grows every candidate", got)
	}
	dead := e
	dead.verdict = 0
	if got := sc.growTypo(nil, kept, tl, dead, 0, 2); len(got) != 0 {
		t.Fatal("a word of no verdict", got)
	}
	det := e
	det.det = true
	if got := sc.growTypo(nil, kept[:1], tl, det, 0, 2); len(got) != 1 || !sc.detEnds[2] {
		t.Fatal("a determiner is marked where it ends", got, sc.detEnds)
	}
	ix, err := newIndex(Options{Words: "testdata/words-typo-alone.tsv", Layers: layerCost})
	if err != nil {
		t.Fatal(err)
	}
	alone := &scan{ix: ix, s: []rune("아이구")}
	alone.prepare()
	alone.adnEnds[0], alone.detEnds[0] = true, true
	at := tail{start: 3, boundTypo: false}
	reads := alone.typoReads(nil, "아이구", at)
	if len(reads) != 1 {
		t.Fatal(reads)
	}
	e = alone.stepsOf(0, 3, reads[0])[0]
	got = alone.growTypo(nil, kept, reads[0], e, 0, 3)
	if len(got) != 1 || got[0].prev != kept[0] || got[0].last.sig != "아이고[열쇠:홀로,오타정정:아이구]" {
		t.Fatal("a right form that only stands alone grows only the candidates that end in no modifying closing and no determiner", got)
	}
}
