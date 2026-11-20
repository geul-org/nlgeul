//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMainUnlinkedRuleAskedReader — read with words-aux-linked.tsv: after 오나, 좋은가, 먹는가, 갈까 and 먹을까 the main verb 보[열쇠:동사] of 보다 is not raised and its trace line says 기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다 under 문장:연결 꼴 뒤가 아닌 본용언, so the auxiliary 보[보조 용언:보] of the same lambda stays first with the preference it has with no rule, and the candidates are as many as with no rule; after 먹어 보다 the auxiliary is first by the first rule, as ever
package parse

import "testing"

func TestMainUnlinkedRuleAskedReader(t *testing.T) {
	aux, main := "보[보조 용언:보]+다[맺음:다]", "보[열쇠:동사]+다[맺음:다]"
	for _, in := range []string{"오나 보다", "좋은가 보다", "먹는가 보다", "갈까 보다", "먹을까 보다"} {
		got, notes := auxLinkedLast(t, in, true, true)
		bare, _ := auxLinkedLast(t, in, false, false)
		if got[0].String() != aux || len(got) != len(bare) {
			t.Fatalf("%q: %q first of %d (%d with no rule)", in, got[0].String(), len(got), len(bare))
		}
		if !almost(auxLinkedPick(t, got, aux).Pref, auxLinkedPick(t, bare, aux).Pref) || !almost(auxLinkedPick(t, got, main).Pref, auxLinkedPick(t, bare, main).Pref) {
			t.Fatalf("%q: preferences", in)
		}
		auxLinkedHas(t, notes,
			"  · "+main+"\t"+ruleMainUnlinked+"\t-1.00\t기각: 앞 어절이 물음 꼴(나, ㄴ가, ㄹ까)로 끝난다",
			"  · "+aux+"\t"+ruleAuxLinked+"\t-1.00\t기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다")
	}
	got, _ := auxLinkedLast(t, "먹어 보다", true, false)
	if got[0].String() != aux {
		t.Fatal("after a linking form 보다 is the auxiliary by the first rule, as ever")
	}
}
