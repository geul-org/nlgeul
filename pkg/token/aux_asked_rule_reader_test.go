//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAuxAskedRuleReader — read with words-aux-linked.tsv: after 오나, 좋은가, 먹는가, 갈까 and 먹을까 the first candidate of 보다 is the auxiliary 보[보조 용언:보] with bonusSentence more than it has with no rule, the main verb 보[열쇠:동사] is not raised and its trace line says 기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다 under 문장:연결 꼴 뒤가 아닌 본용언, the candidates are as many as with no rule, and the auxiliary 있 after the same words is not raised
package token

import "testing"

func TestAuxAskedRuleReader(t *testing.T) {
	aux, main := "보[보조 용언:보]+다[맺음:다]", "보[열쇠:동사]+다[맺음:다]"
	for _, in := range []string{"오나 보다", "좋은가 보다", "먹는가 보다", "갈까 보다", "먹을까 보다"} {
		got, notes := auxLinkedLast(t, in, true, true)
		bare, _ := auxLinkedLast(t, in, false, false)
		if got[0].String() != aux || len(got) != len(bare) {
			t.Fatalf("%q: %q first of %d (%d with no rule)", in, got[0].String(), len(got), len(bare))
		}
		if !almost(auxLinkedPick(t, got, aux).Pref-auxLinkedPick(t, bare, aux).Pref, bonusSentence) || !almost(auxLinkedPick(t, got, main).Pref, auxLinkedPick(t, bare, main).Pref) {
			t.Fatalf("%q: preferences", in)
		}
		auxLinkedHas(t, notes,
			"  · "+aux+"\t"+ruleAuxAsked+"\t1.00\t통과",
			"  · "+main+"\t"+ruleMainUnlinked+"\t-1.00\t기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다",
			"  · "+aux+"\t"+ruleAuxLinked+"\t-1.00\t기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다")
	}
	got, _ := auxLinkedLast(t, "오나 있다", true, false)
	bare, _ := auxLinkedLast(t, "오나 있다", false, false)
	if a := "있[보조 용언:있]+다[맺음:다]"; !almost(auxLinkedPick(t, got, a).Pref, auxLinkedPick(t, bare, a).Pref) {
		t.Fatal("the auxiliary 있 is raised after a question form")
	}
	got, _ = auxLinkedLast(t, "먹어 보다", true, false)
	if got[0].String() != aux {
		t.Fatal("after a linking form 보다 is the auxiliary by the first rule, as ever")
	}
}
