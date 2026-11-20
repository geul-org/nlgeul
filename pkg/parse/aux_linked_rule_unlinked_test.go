//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAuxLinkedRuleUnlinked — where no 어/아 or 고 form stands before (집에 있다, 영화도 보다, 있다 and 보다 as the first word of a sentence, 밥도 주다), the first candidate of 있다, 보다 and 주다 is the main predicate (있 the adjective, 보 and 주 the verb), the auxiliary candidate is not raised (its preference is the one it has with no rule) while the main one gains bonusMainFirst, and the candidates are as many as with no rule
package parse

import "testing"

func TestAuxLinkedRuleUnlinked(t *testing.T) {
	for _, c := range []struct{ in, main, aux string }{
		{"집에 있다", "있[열쇠:형용사]+다[맺음:다]", "있[보조 용언:있]+다[맺음:다]"},
		{"있다", "있[열쇠:형용사]+다[맺음:다]", "있[보조 용언:있]+다[맺음:다]"},
		{"영화도 보다", "보[열쇠:동사]+다[맺음:다]", "보[보조 용언:보]+다[맺음:다]"},
		{"보다", "보[열쇠:동사]+다[맺음:다]", "보[보조 용언:보]+다[맺음:다]"},
		{"밥도 주다", "주[열쇠:동사]+다[맺음:다]", "주[보조 용언:주]+다[맺음:다]"},
	} {
		got, _ := auxLinkedLast(t, c.in, true, false)
		bare, _ := auxLinkedLast(t, c.in, false, false)
		if got[0].String() != c.main || len(got) != len(bare) {
			t.Fatalf("%q: %q first of %d (%d with no rule)", c.in, got[0].String(), len(got), len(bare))
		}
		if !almost(auxLinkedPick(t, got, c.aux).Pref, auxLinkedPick(t, bare, c.aux).Pref) {
			t.Fatalf("%q: the auxiliary is raised", c.in)
		}
		if !almost(auxLinkedPick(t, got, c.main).Pref-auxLinkedPick(t, bare, c.main).Pref, bonusMainFirst) {
			t.Fatalf("%q: the main predicate gains other than bonusMainFirst", c.in)
		}
	}
}
