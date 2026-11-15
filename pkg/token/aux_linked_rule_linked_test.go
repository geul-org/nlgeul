//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAuxLinkedRuleLinked — read with a word table that has 있 (verb and adjective), 보 and 주: after a word that ends in 어 (먹어), 고 (먹고) or the fused 해 the first candidate of 있다, 보다 and 주다 is the auxiliary (보조 용언), whether a space stands between (먹어 보다) or not (먹어보다), and its preference is exactly bonusSentence above that of the main verb of the same letters and above its own with no rule; no candidate is dropped and the other candidates keep their preference
package token

import "testing"

func TestAuxLinkedRuleLinked(t *testing.T) {
	for _, c := range []struct{ in, aux, main string }{
		{"먹어 있다", "있[보조 용언:있]+다[맺음:다]", "있[열쇠:동사]+다[맺음:다]"},
		{"먹고 있다", "있[보조 용언:있]+다[맺음:다]", "있[열쇠:형용사]+다[맺음:다]"},
		{"해 있다", "있[보조 용언:있]+다[맺음:다]", "있[열쇠:동사]+다[맺음:다]"},
		{"먹어있다", "있[보조 용언:있]+다[맺음:다]", "있[열쇠:형용사]+다[맺음:다]"},
		{"먹어 보다", "보[보조 용언:보]+다[맺음:다]", "보[열쇠:동사]+다[맺음:다]"},
		{"먹어보다", "보[보조 용언:보]+다[맺음:다]", "보[열쇠:동사]+다[맺음:다]"},
		{"먹고 보다", "보[보조 용언:보]+다[맺음:다]", "보[열쇠:동사]+다[맺음:다]"},
		{"해 보다", "보[보조 용언:보]+다[맺음:다]", "보[열쇠:동사]+다[맺음:다]"},
		{"먹어 주다", "주[보조 용언:주]+다[맺음:다]", "주[열쇠:동사]+다[맺음:다]"},
		{"해 주다", "주[보조 용언:주]+다[맺음:다]", "주[열쇠:동사]+다[맺음:다]"},
	} {
		got, _ := auxLinkedLast(t, c.in, true, false)
		bare, _ := auxLinkedLast(t, c.in, false, false)
		if got[0].String() != c.aux || len(got) != len(bare) {
			t.Fatalf("%q: %q first of %d (%d with no rule)", c.in, got[0].String(), len(got), len(bare))
		}
		aux, main := auxLinkedPick(t, got, c.aux), auxLinkedPick(t, got, c.main)
		if !almost(aux.Pref-main.Pref, bonusSentence) || !almost(aux.Pref-auxLinkedPick(t, bare, c.aux).Pref, bonusSentence) {
			t.Fatalf("%q: the auxiliary stands %v above the main verb, %v above its own", c.in, aux.Pref-main.Pref, aux.Pref-auxLinkedPick(t, bare, c.aux).Pref)
		}
		if !almost(main.Pref, auxLinkedPick(t, bare, c.main).Pref) {
			t.Fatalf("%q: the main verb is not raised", c.in)
		}
	}
}
