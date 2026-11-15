//ff:func feature=token type=rule control=sequence
//ff:what TestSaisiotCandidates — with the small word table of 사이시옷: 전셋값 before a case particle is one nominal named 사이시옷:ㅅ, its verdict the weaker of 전세 standing alone and 값 before the particle (both with no rebuttal "one syllable long"), above a new word's; 전깃불이 standing alone is 전기 + 불이; none for 아랫층 (not two nominals bound by 사이시옷) and for 전셋값 before a predicate tail (a nominal does not take it)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSaisiotCandidates(t *testing.T) {
	ix := streamReader(t, "words-saisiot.tsv").ix
	front, back := ix.nominalVerdict("전세", rule.TailBare, true), ix.nominalVerdict("값", rule.TailCase, true)
	if front == back {
		t.Fatal("the two words of the test stand differently", front)
	}
	got := ix.saisiotCandidates("전셋값", rule.TailCase)
	if want := min(front, back); len(got) != 1 || got[0].kind != kindNominal || got[0].name != "사이시옷:ㅅ" || got[0].v != want || want <= 1.0/3 {
		t.Fatal(got, front, back)
	}
	if got := ix.saisiotCandidates("전깃불이", rule.TailBare); len(got) != 1 || got[0].v != min(ix.nominalVerdict("전기", rule.TailBare, true), ix.nominalVerdict("불이", rule.TailBare, true)) {
		t.Fatal("전깃불이", got)
	}
	if got := ix.saisiotCandidates("아랫층", rule.TailCase); got != nil {
		t.Fatal("아랫층", got)
	}
	if got := ix.saisiotCandidates("전셋값", rule.TailPred); got != nil {
		t.Fatal("before a predicate tail", got)
	}
}
