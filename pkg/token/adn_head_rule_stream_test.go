//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAdnHeadRuleStream — read as a stream with words-adn-head.tsv and the two rules as the index puts them: with no rule 이 before 역시 is the determiner and 채 after 스무 the adverb; with the rules 이 is the pronoun (the determiner's rebuttal 통과, kept behind with a third of its verdict) and 채 the dependent noun (the adverb's rebuttal 통과); before 가장 큰 the determiner stays (기각: whyAdnAdverbNext, whyAdnAfterNext)
package token

import "testing"

func TestAdnHeadRuleStream(t *testing.T) {
	rules := []sentenceRule{adnHeadRule(), adnHeadAfterRule()}
	line := func(cand, v, why string) string { return "  · " + cand + "\t" + ruleAdnHead + "\t" + v + "\t" + why }
	for _, c := range []struct {
		sentence, word, plain, ruled, note string
	}{
		{"이 역시 어렵다.", "이", "이[관형사:이]", "이[열쇠:체언]", line("이[관형사:이]", "1.00", "통과")},
		{"스무 채.", "채", "채[열쇠:부사]", "채[의존명사:채]", line("채[열쇠:부사]", "1.00", "통과")},
		{"이 가장 큰 집", "이", "이[관형사:이]", "이[관형사:이]", line("이[관형사:이]", "-1.00", "기각: "+whyAdnAdverbNext+", "+whyAdnAfterNext)},
	} {
		if got, _ := ruledFirst(t, "words-adn-head.tsv", c.sentence, c.word); got != c.plain {
			t.Errorf("%s: %s with no rule is %s", c.sentence, c.word, got)
		}
		got, notes := ruledFirst(t, "words-adn-head.tsv", c.sentence, c.word, rules...)
		if got != c.ruled {
			t.Errorf("%s: %s with the rules is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
	first := func(ruled bool) (Branch, int) {
		st := streamReader(t, "words-adn-head.tsv").Stream("이 역시 어렵다.")
		st.rules = nil
		if ruled {
			st.rules = rules
		}
		seg, _ := st.Next()
		return branchNamed(seg.Nodes[0], "이[관형사:이]")
	}
	was, _ := first(false)
	if now, k := first(true); k < 1 || !almost(now.Verdict, was.Verdict/3) || now.Pref != was.Pref {
		t.Fatal("the rebutted determiner is kept behind with a third of its verdict", k, now.Verdict, was.Verdict)
	}
}
