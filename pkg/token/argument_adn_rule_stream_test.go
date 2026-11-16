//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestArgumentAdnRuleStream — read as a stream with words-argument.tsv (the noun 파 counted more than the verb 팔): with no rule 파는 is the noun with the topic particle; with the rule 파는 between an object and a noun (물을 파는 가게) is 팔다's modifying form and the trace says 통과, and after a topic (물은 파는 가게) or before an adverb (물을 파는 빨리) it stays the noun and the trace says why the modifying form falls
package token

import "testing"

func TestArgumentAdnRuleStream(t *testing.T) {
	r := argumentAdnRule()
	adn := "팔[열쇠:동사]+는[꾸밈:는]"
	line := func(v, why string) string {
		return "  · " + adn + "\t" + ruleArgumentAdn + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, ruled, note string
	}{
		{"물을 파는 가게", adn, line("1.00", "통과")},
		{"물은 파는 가게", "파[열쇠:체언]+는[보조사:는]", line("-1.00", "기각: "+whyNoArgumentBefore)},
		{"물을 파는 빨리", "파[열쇠:체언]+는[보조사:는]", line("-1.00", "기각: "+whyNoModifiableNext)},
	} {
		if got, _ := argumentFirst(t, c.sentence, "파는"); got != "파[열쇠:체언]+는[보조사:는]" {
			t.Errorf("%s: 파는 with no rule is %s", c.sentence, got)
		}
		got, notes := argumentFirst(t, c.sentence, "파는", r)
		if got != c.ruled {
			t.Errorf("%s: 파는 with the rule is %s, want %s", c.sentence, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
