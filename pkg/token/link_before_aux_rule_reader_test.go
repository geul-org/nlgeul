//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestLinkBeforeAuxRuleReader — read with words-link-before-aux.tsv: before 있다 the 어/아 form of a main verb whose letters are also a noun (남아, 서) comes first with exactly bonusShortTaken more than it has with no rule, whether a space stands between (남아 있다, 서 있다) or not (남아있다), and then 있다 is the auxiliary; before 먹다 the same form is not raised (남아 먹다, 서 먹다); no candidate is dropped
package token

import "testing"

func TestLinkBeforeAuxRuleReader(t *testing.T) {
	for _, c := range []struct {
		in, link string
		raised   bool
	}{
		{"남아 있다", "남[열쇠:동사]+아[맺음:아]", true},
		{"남아있다", "남[열쇠:동사]+아[맺음:아]", true},
		{"서 있다", "서[열쇠:동사]+(어)[맺음:어(합침)]", true},
		{"남아 먹다", "남[열쇠:동사]+아[맺음:아]", false},
		{"서 먹다", "서[열쇠:동사]+(어)[맺음:어(합침)]", false},
	} {
		got, notes := linkBeforeWords(t, c.in, true)
		bare, _ := linkBeforeWords(t, c.in, false)
		gain := auxLinkedPick(t, got[0], c.link).Pref - auxLinkedPick(t, bare[0], c.link).Pref
		want, note := 0.0, "-1.00\t기각: 바로 뒤 어절의 줄기가 보조 용언이 아니다"
		if c.raised {
			want, note = bonusShortTaken, "1.00\t통과"
		}
		if !almost(gain, want) || len(got[0]) != len(bare[0]) {
			t.Fatalf("%q: raised by %v, %d candidates (%d with no rule)", c.in, gain, len(got[0]), len(bare[0]))
		}
		auxLinkedHas(t, notes, "  · "+c.link+"\t"+ruleLinkBeforeAux+"\t"+note)
		if c.raised && (got[0][0].String() != c.link || got[1][0].String() != "있[보조 용언:있]+다[맺음:다]") {
			t.Fatalf("%q: %q first, then %q", c.in, got[0][0].String(), got[1][0].String())
		}
	}
}
