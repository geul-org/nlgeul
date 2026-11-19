//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNegKindRuleStream — read as a stream with words-neg-kind.tsv (the adjective 못하 counted more than the verb) and the rules as the index puts them: with no rule 못했다 is the adjective after 가지 and after 편하지; with the rules it is the verb after the verb's 가지 (the adjective's rebuttal 통과) and stays the adjective after 편하지 (the adjective's rebuttal 기각: whyNotVerbBefore, the verb's rebuttal 통과); the rebutted candidate is kept, with a third of its verdict
package token

import "testing"

func TestNegKindRuleStream(t *testing.T) {
	rules := []sentenceRule{negKindRule("열쇠:형용사", []string{"열쇠:동사"}, whyNotVerbBefore), negKindRule("열쇠:동사", []string{"열쇠:형용사"}, whyNotAdjBefore)}
	const adj, verb = "못하[열쇠:형용사]+(였)[선어말:ㅆ]+다[맺음:다]", "못하[열쇠:동사]+(였)[선어말:ㅆ]+다[맺음:다]"
	line := func(cand, v, why string) string { return "  · " + cand + "\t" + ruleNegKind + "\t" + v + "\t" + why }
	for _, c := range []struct {
		sentence, ruled, adjNote, verbNote string
	}{
		{"가지 못했다", verb, line(adj, "1.00", "통과"), line(verb, "-1.00", "기각: "+whyNotAdjBefore)},
		{"편하지 못했다", adj, line(adj, "-1.00", "기각: "+whyNotVerbBefore), line(verb, "1.00", "통과")},
	} {
		if got, _ := ruledFirst(t, "words-neg-kind.tsv", c.sentence, "못했다"); got != adj {
			t.Errorf("%s: 못했다 with no rule is %s", c.sentence, got)
		}
		got, notes := ruledFirst(t, "words-neg-kind.tsv", c.sentence, "못했다", rules...)
		if got != c.ruled {
			t.Errorf("%s: 못했다 with the rules is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.adjNote, c.verbNote)
	}
	second := func(ruled bool) (Branch, int) {
		st := streamReader(t, "words-neg-kind.tsv").Stream("가지 못했다")
		st.rules = nil
		if ruled {
			st.rules = rules
		}
		st.Next()
		seg, _ := st.Next()
		return branchNamed(seg.Nodes[0], adj)
	}
	was, _ := second(false)
	if now, k := second(true); k < 1 || !almost(now.Verdict, was.Verdict/3) || now.Pref != was.Pref {
		t.Fatal("the rebutted adjective is kept behind with a third of its verdict", k, now.Verdict, was.Verdict)
	}
}
