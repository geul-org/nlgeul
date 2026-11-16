//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestGiveLinkedRuleStream — read as a stream with words-give.tsv (the adjective 다르 counted more than the verb 달): with no rule 달라 is 다르다 with 아 fused; with the rule 달라 after the 어/아 form (도와 달라) is 달다 with 라 and the trace says 통과, and after 고 (읽고 달라) it stays 다르다 and the trace says 기각: whyNotConn
package token

import "testing"

func TestGiveLinkedRuleStream(t *testing.T) {
	r := giveLinkedRule()
	give, other := "달[열쇠:동사]+라[맺음:라]", "다르[열쇠:형용사]+(아)[맺음:어(합침)]"
	line := func(v, why string) string {
		return "  · " + give + "\t" + ruleGiveLinked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"도와 달라", give, line("1.00", "통과")},
		{"읽고 달라", other, line("-1.00", "기각: "+whyNotConn)},
	} {
		if got, _ := giveFirst(t, c.sentence, "달라"); got != other {
			t.Errorf("%s: 달라 with no rule is %s", c.sentence, got)
		}
		got, notes := giveFirst(t, c.sentence, "달라", r)
		if got != c.ruled {
			t.Errorf("%s: 달라 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
