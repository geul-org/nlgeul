//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestDeterminerFormRuleStream — read as a stream with words-form.tsv (the adjectives 다르다·그렇다 counted far more than the determiners 다른·그런): with no rule the adjective's modifying form comes first; with the rule 다른 at the head of a sentence, 그런 after a word ending in 과 and 다른 right after an opening quotation mark are the determiner and the trace says 통과, and 다른 after a word ending in 과 stays the adjective's form and the trace says 기각: whyComparedBefore
package parse

import "testing"

func TestDeterminerFormRuleStream(t *testing.T) {
	r := determinerFormRule()
	line := func(stem, v, why string) string {
		return "  · " + stem + "[열쇠:관형사]\t" + ruleDeterminerForm + "\t" + v + "\t" + why
	}
	for _, c := range []struct {
		sentence, word, plain, ruled, note string
	}{
		{"다른 사람이 왔다", "다른", "다르[열쇠:형용사]+(ㄴ)[꾸밈:ㄴ]", "다른[열쇠:관형사]", line("다른", "1.00", "통과")},
		{"책과 그런 사람이 왔다", "그런", "그렇[열쇠:형용사]+(ㄴ)[꾸밈:ㄴ]", "그런[열쇠:관형사]", line("그런", "1.00", "통과")},
		{"“다른 책이다", "다른", "다르[열쇠:형용사]+(ㄴ)[꾸밈:ㄴ]", "다른[열쇠:관형사]", line("다른", "1.00", "통과")},
		{"책과 다른 사람이 왔다", "다른", "다르[열쇠:형용사]+(ㄴ)[꾸밈:ㄴ]", "다르[열쇠:형용사]+(ㄴ)[꾸밈:ㄴ]", line("다른", "-1.00", "기각: "+whyComparedBefore)},
	} {
		if got, _ := formFirst(t, c.sentence, c.word); got != c.plain {
			t.Errorf("%s: %s with no rule is %s, want %s", c.sentence, c.word, got, c.plain)
		}
		got, notes := formFirst(t, c.sentence, c.word, r)
		if got != c.ruled {
			t.Errorf("%s: %s with the rule is %s, want %s", c.sentence, c.word, got, c.ruled)
		}
		auxLinkedHas(t, notes, c.note)
	}
}
