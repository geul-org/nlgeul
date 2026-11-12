//ff:func feature=token type=reader control=sequence
//ff:what TestStepsOf — a word whose stem stands as two kinds is two readings, each with its own kind, verdict and preference (the kind's count); in the first layer both stand whole; a closed-set word of two sets and an atom of two names are two readings of one verdict and preference; a word of one way is one reading
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStepsOf(t *testing.T) {
	sc := wordsScan("지금", layerCost)
	sc.ix.lex["지금"] = word{kinds: kindNominal | kindAdverb, count: 100, by: [6]int{60, 0, 0, 40}}
	got := sc.stepsOf(0, 2, tail{start: 2})
	if len(got) != 2 || got[0].sig != "지금[열쇠:체언]" || got[1].sig != "지금[열쇠:부사]" || got[0].kind != kindNominal || got[1].kind != kindAdverb {
		t.Fatal(got)
	}
	if got[0].verdict == got[1].verdict || got[0].pref <= got[1].pref {
		t.Fatal("each kind has its own verdict and count", got[0].verdict, got[1].verdict, got[0].pref, got[1].pref)
	}
	if one := sc.wordOf(0, 2, tail{start: 2}); one.sig != got[0].sig {
		t.Fatal("wordOf is the first", one.sig)
	}
	first := wordsScan("작", layerKnown).stepsOf(0, 1, tail{start: 1})
	if len(first) != 2 || first[0].verdict != 1 || first[1].verdict != 1 || first[0].sig != "작[열쇠:체언]" || first[1].sig != "작[열쇠:형용사]" {
		t.Fatal(first)
	}
	two := scanOf("두").stepsOf(0, 1, tail{start: 1, pieces: []piece{{"다", "맺음:다", 1, -4, rule.TailPred, false, ""}}})
	if len(two) != 2 || two[0].sig != "두[관형사:두]+다[맺음:다]" || two[1].sig != "두[보조 용언:두]+다[맺음:다]" || two[0].verdict != two[1].verdict || two[0].pref != two[1].pref {
		t.Fatal("a word of two sets", two)
	}
	atom := scanOf("1/3").stepsOf(0, 3, tail{start: 3})
	if len(atom) != 2 || atom[0].sig == atom[1].sig || atom[0].verdict != atom[1].verdict || atom[0].pref != atom[1].pref {
		t.Fatal("an atom of two names", atom)
	}
	if one := scanOf("고양이").stepsOf(0, 3, tail{start: 3}); len(one) != 1 || one[0].word.Parts[0].Name != (Name{}) {
		t.Fatal("a new word is one reading", one)
	}
}
