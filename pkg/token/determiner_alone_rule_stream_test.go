//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDeterminerAloneRuleStream — read as a stream with words-determiner.tsv (the pronouns 이·그 counted more than the determiners): with no rule the pronoun comes first; with the rule 이 before a noun (이 책을), 그 right after an opening quotation mark (“그 책을) and 그 right before one (그 “책”을) are the determiner and the trace says 통과, and 그 before an adverb (그 다시) and 이 at the end of the sentence stay the pronoun and the trace says 기각: whyNoModifiableNext
package token

import "testing"

func TestDeterminerAloneRuleStream(t *testing.T) {
	r := determinerAloneRule()
	line := func(stem, v, why string) string {
		return "  · " + stem + "[관형사:" + stem + "]\t" + ruleDeterminerAlone + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, word, ruled, note string
	}{
		{"이 책을 읽었다", "이", "이[관형사:이]", line("이", "1.00", "통과")},
		{"“그 책을 읽었다", "그", "그[관형사:그]", line("그", "1.00", "통과")},
		{"그 “책”을 읽었다", "그", "그[관형사:그]", line("그", "1.00", "통과")},
		{"그 다시 읽었다", "그", "그[열쇠:체언]", line("그", "-1.00", "기각: "+whyNoModifiableNext)},
		{"책을 읽은 이", "이", "이[열쇠:체언]", line("이", "-1.00", "기각: "+whyNoModifiableNext)},
	} {
		if got, _ := determinerFirst(t, c.sentence, c.word); got != c.word+"[열쇠:체언]" {
			t.Errorf("%s: %s with no rule is %s", c.sentence, c.word, got)
		}
		got, notes := determinerFirst(t, c.sentence, c.word, r)
		if got != c.ruled {
			t.Errorf("%s: %s with the rule is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
