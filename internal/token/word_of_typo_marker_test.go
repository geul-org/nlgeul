//ff:func feature=token type=reader control=sequence
//ff:what TestWordOfTypoMarker — a marker written wrong is a part of the word with the right closing as its letters, the letters it is written with in Written, and the word is a reading with letters written wrong (typo), printed with the mark; a word whose markers are written right is not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWordOfTypoMarker(t *testing.T) {
	sc := typoScan("할께", layerCost)
	pc := piece{"(ㄹ)게", "맺음:ㄹ게", 0.6, -3, rule.TailPred, false, "ㄹ께"}
	e := sc.wordOf(0, 2, tail{start: 1, base: '하', pieces: []piece{pc}})
	last := e.word.Parts[1]
	if last.Text != "게" || last.Hidden != "ㄹ" || last.Written != "ㄹ께" || last.Stem || len(last.Names) != 1 || last.Names[0] != (Name{"맺음", "ㄹ게"}) {
		t.Fatal(last)
	}
	if !e.typo || e.sig != "하[열쇠:동사]+(ㄹ)게[맺음:ㄹ게,오타정정:ㄹ께]" || e.word.Parts[0].Written != "" {
		t.Fatal(e.typo, e.sig)
	}
	pc.wrote = ""
	if right := sc.wordOf(0, 2, tail{start: 1, base: '하', pieces: []piece{pc}}); right.typo || right.sig != "하[열쇠:동사]+(ㄹ)게[맺음:ㄹ게]" {
		t.Fatal(right.typo, right.sig)
	}
	if stem := typoScan("맞겨", layerCost).wordOf(0, 2, tail{start: 2, base: '기', shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{"(어)", "맺음:어(합침)", 1, 0, rule.TailPred, false, ""}}}); !stem.typo {
		t.Fatal("a stem written wrong is a reading with letters written wrong too")
	}
}
