//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestMalLinkedRuleReader — read with words-link-before-aux.tsv: after 울지 and 울고 the first candidate of 마라 is 말+(아)라 (말다 without its ㄹ) with exactly bonusShortTaken more than it has with no rule, ahead of the noun 마라; after 집 it is not raised and the noun stays first; no candidate is dropped
package token

import "testing"

func TestMalLinkedRuleReader(t *testing.T) {
	mal, noun := "말[열쇠:동사]+(아)라[맺음:어라(합침)]", "마라[열쇠:체언]"
	for _, c := range []struct {
		in     string
		raised bool
	}{
		{"울지 마라", true},
		{"울고 마라", true},
		{"집 마라", false},
	} {
		got, _ := linkBeforeWords(t, c.in, true)
		bare, _ := linkBeforeWords(t, c.in, false)
		gain := auxLinkedPick(t, got[1], mal).Pref - auxLinkedPick(t, bare[1], mal).Pref
		first := mal
		if !c.raised {
			first = noun
		}
		if c.raised && !almost(gain, bonusShortTaken) || !c.raised && !almost(gain, 0) || got[1][0].String() != first || len(got[1]) != len(bare[1]) {
			t.Fatalf("%q: %q first, raised by %v", c.in, got[1][0].String(), gain)
		}
	}
}
