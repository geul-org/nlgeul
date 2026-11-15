//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestNegAdverbRuleSentence — end to end, over the small word table with 안 as a noun and an adverb of one count (as the word table has it), 먹 and 하 as verbs: 안 right before a verb, alone or written onto it, is read first as the adverb (밥을 안 먹었다, 안해), which it is not without the rules of a sentence; 안 with a particle (집 안에) and 안 right before a noun stay the noun
package token

import (
	"strings"
	"testing"
)

func TestNegAdverbRuleSentence(t *testing.T) {
	r := streamReader(t, "words.tsv")
	r.ix.lex["안"] = word{kinds: kindNominal | kindAdverb, count: 9000, by: [6]int{9000, 0, 0, 9000}}
	r.ix.lex["밥"] = word{kinds: kindNominal, count: 3000, by: [6]int{3000}}
	r.ix.lex["집"] = word{kinds: kindNominal, count: 4000, by: [6]int{4000}}
	for _, c := range []struct{ in, want string }{
		{"밥을 안 먹었다.", "안[열쇠:부사]"},
		{"안해.", "안[열쇠:부사]"},
		{"집 안에 있다.", "안[열쇠:체언]+에[부사격:에]"},
		{"안 학교.", "안[열쇠:체언]"},
	} {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		if !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
	}
	saved := r.ix.sentence
	r.ix.sentence = nil
	defer func() { r.ix.sentence = saved }()
	if got := bestNodes(r, "밥을 안 먹었다."); len(got) < 2 || got[1].String() != "안[열쇠:체언]" {
		t.Fatal("without the rules of a sentence the noun comes first", got)
	}
}
