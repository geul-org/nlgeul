//ff:func feature=token type=reader control=sequence
//ff:what TestDeterminerOneRuleReader — read with words-one.tsv: at the head of a sentence (한 번 웃었다) the first candidate of 한 is the determiner with bonusOne more than it has with no rule, 하다 with ㄴ is not raised and the candidates are as many as with no rule; after an object (일을 한 사람) 하다 with ㄴ stays first, the determiner is not raised and its trace line says 기각: whyHaTakes; 한 written onto the word before it (사람한 번) is not asked about
package token

import (
	"strings"
	"testing"
)

func TestDeterminerOneRuleReader(t *testing.T) {
	det, ha := "한[관형사:한]", "하[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]"
	got, notes := oneCandidates(t, "한 번 웃었다", true, true)
	bare, _ := oneCandidates(t, "한 번 웃었다", false, false)
	if bare[0].String() != ha || got[0].String() != det || len(got) != len(bare) {
		t.Fatalf("%q first of %d (%q of %d with no rule)", got[0].String(), len(got), bare[0].String(), len(bare))
	}
	if !almost(auxLinkedPick(t, got, det).Pref-auxLinkedPick(t, bare, det).Pref, bonusOne) || !almost(auxLinkedPick(t, got, ha).Pref, auxLinkedPick(t, bare, ha).Pref) {
		t.Fatal("preferences at the head of a sentence")
	}
	auxLinkedHas(t, notes, "  · "+det+"\t"+ruleDeterminerOne+"\t1.00\t통과")
	got, notes = oneCandidates(t, "일을 한 사람", true, true)
	bare, _ = oneCandidates(t, "일을 한 사람", false, false)
	if got[0].String() != ha || !almost(auxLinkedPick(t, got, det).Pref, auxLinkedPick(t, bare, det).Pref) {
		t.Fatalf("after an object: %q first", got[0].String())
	}
	auxLinkedHas(t, notes, "  · "+det+"\t"+ruleDeterminerOne+"\t-1.00\t기각: "+whyHaTakes)
	_, notes = oneCandidates(t, "사람한 번", true, true)
	if strings.Contains(strings.Join(notes, "\n"), ruleDeterminerOne) {
		t.Fatal("한 written onto the word before it is asked about", notes)
	}
}
