//ff:func feature=token type=reader control=sequence
//ff:what TestModifierWritten — a new word whose letters split at a place inside them into a modifier and a known word it modifies: 못와 as a whole, 두손 before the particle 으로 (the place inside the stem), 할 before the closing 게 (the place where the tail begins: 하+ㄹ and 게, 거+이, which is read so where the scan has had a word end in a modifying closing before it); not a tail that names its own stem or stands for a number, not a word of one syllable (the place at its end is not asked), and not letters that hold no such place (곰곰)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestModifierWritten(t *testing.T) {
	verb := word{kinds: kindVerb, count: 10000, by: [6]int{0, 10000}}
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}
	whole := wordsScan("못와", layerCost)
	whole.ix.lex["오"] = verb
	if !whole.modifierWritten(0, 2, tail{start: 2}) {
		t.Fatal("못와 as a whole")
	}
	if whole.modifierWritten(0, 2, tail{start: 2, only: "못와"}) || whole.modifierWritten(0, 2, tail{start: 2, as: "수"}) {
		t.Fatal("a tail that names its own stem or stands for a number is not asked")
	}
	ro := []piece{{text: "으로", names: "부사격:으로", verdict: 1, tail: rule.TailCase}}
	hands := wordsScan("두손으로", layerCost)
	hands.ix.lex["손"] = noun
	if !hands.modifierWritten(0, 4, tail{start: 2, pieces: ro}) {
		t.Fatal("두손 before 으로")
	}
	ge := []piece{{text: "게", names: "맺음:게", verdict: 1, tail: rule.TailPred}}
	will := wordsScan("할게", layerCost)
	if will.modifierWritten(0, 2, tail{start: 1, pieces: ge}) {
		t.Fatal("게 is read as 거+이 only where a word that ends in a modifying closing ends before it")
	}
	will.adnEnds[1] = true
	if !will.modifierWritten(0, 2, tail{start: 1, pieces: ge}) {
		t.Fatal("할 before the closing 게: the modifying form and 게 (거+이)")
	}
	if wordsScan("못", layerCost).modifierWritten(0, 1, tail{start: 1}) {
		t.Fatal("a word of one syllable")
	}
	if wordsScan("곰곰", layerCost).modifierWritten(0, 2, tail{start: 2}) {
		t.Fatal("letters that hold no such place")
	}
}
