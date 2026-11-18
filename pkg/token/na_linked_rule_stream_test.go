//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNaLinkedRuleStream — read as a stream with words-na-linked.tsv (the noun 나서 and the pronoun 나 with 는 read before the verb 나 by less than bonusShortTaken): with no rule 나서 after 먹고 is the noun and 난 after 끝나고 the pronoun with 는; with the rule both are 나다 (나+(아)서, 나+(ㄴ)) and the trace says 통과; 나서 after 집에 stays the noun (기각: whyNotNaLinked), 나는 after 먹고 stays the pronoun with 는 (기각: whyNaTopic) and 남은 after 먹고 stays 남+은 (기각: whyNaNamed)
package token

import "testing"

func TestNaLinkedRuleStream(t *testing.T) {
	r := naLinkedRule()
	line := func(cand, v, why string) string {
		return "  · " + cand + "\t" + ruleNaLinked + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, word, bare, ruled, note string
	}{
		{"먹고 나서", "나서", "나서[열쇠:체언]", "나[열쇠:동사]+(아)서[맺음:어서(합침)]", line("나[열쇠:동사]+(아)서[맺음:어서(합침)]", "1.00", "통과")},
		{"끝나고 난 뒤", "난", "나[열쇠:체언]+(는)[보조사:는]", "나[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]", line("나[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]", "1.00", "통과")},
		{"집에 나서", "나서", "나서[열쇠:체언]", "나서[열쇠:체언]", line("나[열쇠:동사]+(아)서[맺음:어서(합침)]", "-1.00", "기각: "+whyNotNaLinked)},
		{"먹고 나는 집에", "나는", "나[열쇠:체언]+는[보조사:는]", "나[열쇠:체언]+는[보조사:는]", line("나[열쇠:동사]+는[꾸밈:는]", "-1.00", "기각: "+whyNaTopic)},
		{"먹고 남은 집", "남은", "남[열쇠:체언]+은[보조사:은]", "남[열쇠:체언]+은[보조사:은]", line("나[열쇠:동사]+(ㅁ)[이름:ㅁ]+은[보조사:은]", "-1.00", "기각: "+whyNaNamed)},
	} {
		if got, _ := ruledFirst(t, "words-na-linked.tsv", c.sentence, c.word); got != c.bare {
			t.Errorf("%s: %s with no rule is %s, want %s", c.sentence, c.word, got, c.bare)
		}
		got, notes := ruledFirst(t, "words-na-linked.tsv", c.sentence, c.word, r)
		if got != c.ruled {
			t.Errorf("%s: %s with the rule is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
