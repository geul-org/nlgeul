//ff:func feature=token type=rule control=sequence
//ff:what TestKnownCandidates — a stem of two kinds is read once for each kind whose graph lets it stand, each with that kind's own verdict; a kind whose graph turns it down is not read; in the first layer every kind stands whole; a closed-set word that stands as the kind of its set is read under the set's name only; a stem no table lists has no candidate
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestKnownCandidates(t *testing.T) {
	ix := wordsScan("x", layerCost).ix
	ix.lex["지금"] = word{kinds: kindNominal | kindAdverb, count: 100, by: [6]int{60, 0, 0, 40}}
	got := ix.knownCandidates("지금", rule.TailBare, false)
	if len(got) != 2 || got[0].kind != kindNominal || got[0].name != "열쇠:체언" || got[1].kind != kindAdverb || got[1].name != "열쇠:부사" {
		t.Fatal(got)
	}
	if want := ix.stemVerdict(kindNominal, rule.Piece{Form: "지금", Tail: rule.TailBare, Count: 60}); got[0].v != want || got[0].v <= 0 {
		t.Fatal("the verdict of the kind's own graph", got[0].v, want)
	}
	if want := ix.stemVerdict(kindAdverb, rule.Piece{Form: "지금", Tail: rule.TailBare, Count: 40}); got[1].v != want || got[1].v == got[0].v {
		t.Fatal("each kind has its own value", got[1].v, want)
	}
	if got = ix.knownCandidates("지금", rule.TailPred, false); len(got) != 0 {
		t.Fatal("no kind takes a closing", got)
	}
	if got = ix.knownCandidates("지금", rule.TailCase, false); len(got) != 1 || got[0].kind != kindNominal {
		t.Fatal("only the kind that takes a particle is read", got)
	}
	first := wordsScan("x", layerKnown).ix.knownCandidates("작", rule.TailCase, false)
	if len(first) != 2 || first[0].v != 1 || first[1].v != 1 || first[0].name != "열쇠:체언" || first[1].name != "열쇠:형용사" {
		t.Fatal("in the first layer every kind stands whole", first)
	}
	if got = ix.knownCandidates("이", rule.TailBare, false); len(got) != 1 || got[0].name != "관형사:이" || got[0].kind != kindDeterminer {
		t.Fatal("a closed-set word under its set's name", got)
	}
	if got = ix.knownCandidates("이", rule.TailCase, false); len(got) != 1 || got[0].name != "열쇠:체언" {
		t.Fatal("named by the table's kind where it does not stand as its set's", got)
	}
	if got = ix.knownCandidates("없는말", rule.TailBare, false); len(got) != 0 {
		t.Fatal(got)
	}
	none := scanOf("x").ix
	if got = none.knownCandidates("이", rule.TailBare, false); len(got) != 1 || got[0].kind != 0 || got[0].name != "관형사:이" {
		t.Fatal("with no word table only the closed-set words are known, under their set's names", got)
	}
	if got = none.knownCandidates("고양이", rule.TailBare, false); len(got) != 0 {
		t.Fatal("a word only the table knows", got)
	}
}
