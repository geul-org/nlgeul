//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestAuxMoreSentence — end to end, over the small word table with the verbs 오 and 달 added and its total made a hundred times as large: the auxiliary verbs 오다 (-어 오다) and 달다 (-어 달라) written onto the 어/아 form of a main verb are read first as the main verb's 어/아 form and the verb, not as one new word: 해온, 해달라고
package token

import (
	"strings"
	"testing"
)

func TestAuxMoreSentence(t *testing.T) {
	r := streamReader(t, "words.tsv")
	enough := 10000
	verb := word{kinds: kindVerb, count: enough, by: [6]int{0, enough}}
	for _, w := range []string{"오", "달"} {
		r.ix.lex[w] = verb
	}
	r.ix.lexTotal *= 100
	for _, c := range []struct{ in, want string }{
		{"일을 해온 사람", "하[열쇠:동사]+(여)[맺음:어(합침)] │ 오[열쇠:동사]+(ㄴ)[꾸밈:ㄴ]"},
		{"꼭 해달라고 했다.", "하[열쇠:동사]+(여)[맺음:어(합침)] │ 달[열쇠:동사]+라고[맺음:라고]"},
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
}
