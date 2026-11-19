//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestMainBeforeAuxSentence — end to end, over the small word table with the verbs 빼 and 꾸 added and its total made a hundred times as large: a one-syllable main verb that begins its unit with an auxiliary verb written onto its 어/아 form is read first as the main verb's 어/아 form and the auxiliary, and its node stands fully (the rebuttal that it is one syllable long does not hold): 가봐라, 꿔줄; 빼둬 has that way too, but with the values fitted on 2026-11-19 the new word 빼둬 comes first; the same verb away from the unit's head keeps the rebuttal
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestMainBeforeAuxSentence(t *testing.T) {
	r := streamReader(t, "words.tsv")
	enough := 10000
	verb := word{kinds: kindVerb, count: enough, by: [6]int{0, enough}}
	for _, w := range []string{"빼", "꾸"} {
		r.ix.lex[w] = verb
	}
	r.ix.lexTotal *= 100
	for _, c := range []struct{ in, want, main string }{
		{"빨리 가봐라.", "가[열쇠:동사]+(아)[맺음:어(합침)] │ 보[보조 용언:보]+(아)라[맺음:어라(합침)]", "가[열쇠:동사]+(아)[맺음:어(합침)]"},
		{"돈을 꿔줄 사람", "꾸[열쇠:동사]+(어)[맺음:어(합침)] │ 주[보조 용언:주]+(ㄹ)[꾸밈:ㄹ]", "꾸[열쇠:동사]+(어)[맺음:어(합침)]"},
	} {
		var words []string
		main := -1.0
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
			if b.String() == c.main {
				main = b.Verdict
			}
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		if !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
		if main != 1 {
			t.Errorf("%s: the node %s stands at %v", c.in, c.main, main)
		}
	}
	if !slices.Contains(streamLines(r, "그걸 빼둬."), "빼[열쇠:동사]+(어)[맺음:어(합침)] 두[보조 용언:두]+(어)[맺음:어(합침)]") {
		t.Error("빼둬: no way of the main verb and the auxiliary")
	}
	for _, b := range bestNodes(r, "또가봐라.") {
		if strings.HasPrefix(b.String(), "가[열쇠:동사]") && b.Verdict >= 1 {
			t.Errorf("away from the unit's head %s stands at %v", b.String(), b.Verdict)
		}
	}
}
