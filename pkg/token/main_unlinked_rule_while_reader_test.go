//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestMainUnlinkedRuleWhileReader — read with words-link-before-aux.tsv: after 먹다 the main verb 보 of 보니 is not raised (기각: 앞 어절이 다·다가로 끝난다), so the auxiliary 보 of the same lambda stays first with the preference it has with no rule; after 집에 보니 is the main verb, raised by bonusMainFirst; after 먹다 the main verb of 보아 (no 니, 니까, 면) and the main 있 of 있으니 (not 보) are raised as ever
package token

import "testing"

func TestMainUnlinkedRuleWhileReader(t *testing.T) {
	for _, c := range []struct {
		in, first, other string
		gain             float64
	}{
		{"먹다 보니", "보[보조 용언:보]+니[맺음:니]", "보[열쇠:동사]+니[맺음:니]", 0},
		{"집에 보니", "보[열쇠:동사]+니[맺음:니]", "보[보조 용언:보]+니[맺음:니]", bonusMainFirst},
		{"먹다 보아", "보[열쇠:동사]+아[맺음:아]", "보[보조 용언:보]+아[맺음:아]", bonusMainFirst},
		{"먹다 있으니", "있[열쇠:형용사]+으[이음:으]+니[맺음:니]", "있[보조 용언:있]+으[이음:으]+니[맺음:니]", bonusMainFirst},
	} {
		got, notes := linkBeforeWords(t, c.in, true)
		bare, _ := linkBeforeWords(t, c.in, false)
		last, none := got[len(got)-1], bare[len(bare)-1]
		if last[0].String() != c.first || len(last) != len(none) {
			t.Fatalf("%q: %q first", c.in, last[0].String())
		}
		if !almost(last[0].Pref-auxLinkedPick(t, none, c.first).Pref, c.gain) || !almost(auxLinkedPick(t, last, c.other).Pref, auxLinkedPick(t, none, c.other).Pref) {
			t.Fatalf("%q: preferences", c.in)
		}
		if c.in == "먹다 보니" {
			auxLinkedHas(t, notes, "  · "+c.other+"\t"+ruleMainUnlinked+"\t-1.00\t기각: 앞 어절이 다·다가로 끝난다")
		}
	}
}
