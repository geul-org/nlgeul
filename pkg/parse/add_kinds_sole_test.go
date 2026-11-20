//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddKindsSole — on each kind's graph a one-syllable form seen enough times stands fully when its word is the whole of its unit (Sole), above where it is not, as the rebuttal that it is one syllable long does not hold against it; a two-syllable form stands the same either way, seen enough or once; and a one-syllable form seen once stands below 1 even as a sole word, where a two-syllable form seen once stands (the rebuttal that it was rarely seen still holds)
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddKindsSole(t *testing.T) {
	ix := &index{lexTotal: 1e8}
	ix.addKinds()
	fits := map[int]int{kindNominal: rule.TailCase, kindVerb: rule.TailPred, kindAdj: rule.TailAdj, kindAdverb: rule.TailAux, kindDeterminer: rule.TailBare, kindAlone: rule.TailBare}
	judge := func(kind int, form string, count int, sole bool) float64 {
		v, _ := verdictOf(ix.kinds[kind], rule.Piece{Form: form, Tail: fits[kind], Count: count, Sole: sole}, "", false)
		return v
	}
	for _, kind := range kindOrder {
		in, sole := judge(kind, "솥", ix.seenEnough(), false), judge(kind, "솥", ix.seenEnough(), true)
		if sole != 1 || in >= sole || in <= 0 {
			t.Fatal(kindLabel[kind], "a one-syllable form seen enough", in, sole)
		}
		if judge(kind, "가르", ix.seenEnough(), true) != judge(kind, "가르", ix.seenEnough(), false) || judge(kind, "가르", 1, true) != judge(kind, "가르", 1, false) {
			t.Fatal(kindLabel[kind], "a two-syllable form")
		}
		few := judge(kind, "솥", 1, true)
		if few <= 0 || few >= 1 || few != judge(kind, "가르", 1, false) || few <= judge(kind, "솥", 1, false) {
			t.Fatal(kindLabel[kind], "a one-syllable form seen once", few)
		}
	}
}
