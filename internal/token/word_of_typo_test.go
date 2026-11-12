//ff:func feature=token type=reader control=sequence
//ff:what TestWordOfTypo — a stem written wrong is printed and looked up as the right form with the letters it is written with kept apart (Written) and its readings those of the right form's kind alone; a word written right and a new word have no written form; a misspelt stem the right form does not take carries no written form
package token

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWordOfTypo(t *testing.T) {
	sc := typoScan("맞겨", layerCost)
	conn := []piece{{"(어)", "맺음:어(합침)", 1, 0, rule.TailPred, false, ""}}
	e := sc.wordOf(0, 2, tail{start: 2, base: '기', shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: conn})
	stem := e.word.Parts[0]
	if stem.Text != "맡기" || stem.Written != "맞기" || !stem.Stem || !reflect.DeepEqual(stem.Names, []Name{{"열쇠", "동사"}}) {
		t.Fatal(stem)
	}
	if e.sig != "맡기[열쇠:동사,오타정정:맞기]+(어)[맺음:어(합침)]" || e.verdict < 0.54 || e.verdict > 0.56 {
		t.Fatal(e.sig, e.verdict)
	}
	if e.word.Parts[1].Written != "" {
		t.Fatal("a marker is written right", e.word.Parts[1])
	}
	right := typoScan("맡겨", layerCost).wordOf(0, 2, tail{start: 2, base: '기', pieces: conn})
	if right.sig != "맡기[열쇠:동사]+(어)[맺음:어(합침)]" || right.word.Parts[0].Written != "" || right.verdict <= e.verdict {
		t.Fatal(right.sig, right.verdict)
	}
	if e = sc.wordOf(0, 2, tail{start: 2, shown: "맡기", wrote: "맞기", kinds: kindVerb, pieces: []piece{{"이", "주격:이", 1, -1, rule.TailCase, false, ""}}}); e.word.Parts[0].Written != "" || e.word.Parts[0].Names != nil {
		t.Fatal("a misspelt verb stem before a particle", e)
	}
}
