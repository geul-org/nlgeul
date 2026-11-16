//ff:func feature=token type=reader control=sequence
//ff:what TestLimitOneRuleReader — read with words-one.tsv: after the modifying 는 (살아 있는 한 웃는다) the first candidate of 한 is the noun 한[열쇠:체언] with bonusOne more than it has with no rule, the determiner and 하다 with ㄴ are not raised and the candidates are as many as with no rule; at the head of a sentence the noun is not raised
package token

import "testing"

func TestLimitOneRuleReader(t *testing.T) {
	noun, det, ha := "한[열쇠:체언]", "한[관형사:한]", "하[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]"
	got, notes := oneCandidates(t, "살아 있는 한 웃는다", true, true)
	bare, _ := oneCandidates(t, "살아 있는 한 웃는다", false, false)
	if bare[0].String() != ha || got[0].String() != noun || len(got) != len(bare) {
		t.Fatalf("%q first of %d (%q of %d with no rule)", got[0].String(), len(got), bare[0].String(), len(bare))
	}
	if !almost(auxLinkedPick(t, got, noun).Pref-auxLinkedPick(t, bare, noun).Pref, bonusOne) || !almost(auxLinkedPick(t, got, det).Pref, auxLinkedPick(t, bare, det).Pref) || !almost(auxLinkedPick(t, got, ha).Pref, auxLinkedPick(t, bare, ha).Pref) {
		t.Fatal("preferences after 는")
	}
	auxLinkedHas(t, notes, "  · "+noun+"\t"+ruleLimitOne+"\t1.00\t통과")
	got, _ = oneCandidates(t, "한 번 웃었다", true, false)
	bare, _ = oneCandidates(t, "한 번 웃었다", false, false)
	if !almost(auxLinkedPick(t, got, noun).Pref, auxLinkedPick(t, bare, noun).Pref) {
		t.Fatal("the noun is raised at the head of a sentence")
	}
}
