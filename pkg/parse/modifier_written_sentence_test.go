//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestModifierWrittenSentence — end to end, over the small word table with the words of the sentences added and its total made a hundred times as large: a modifier written onto a known word it modifies is read as two words first: 못가 (못 and 가+아), 두손으로 (두 and 손+으로), 좀더 (좀 and 더), 갈거야 (가+ㄹ and 거+이+야), and the new word of the letters whole stays a candidate there, with its particle after it (두손+으로) and with the lower verdict of the rebuttal; letters whose part after the modifier is no known word (두루미) are a new word with a new word's own verdict
package parse

import (
	"strings"
	"testing"
)

func TestModifierWrittenSentence(t *testing.T) {
	r := streamReader(t, "words.tsv")
	r.ix.lex["손"] = word{kinds: kindNominal, count: 5000, by: [6]int{5000}}
	r.ix.lex["좀"] = word{kinds: kindAdverb, count: 9000, by: [6]int{0, 0, 0, 9000}}
	r.ix.lex["더"] = word{kinds: kindAdverb, count: 9000, by: [6]int{0, 0, 0, 9000}}
	r.ix.lex["못"] = word{kinds: kindAdverb, count: 9000, by: [6]int{0, 0, 0, 9000}}
	r.ix.lexTotal *= 100
	for _, c := range []struct{ in, want, behind string }{
		{"오늘은 못가.", "못[열쇠:부사] │ 가[열쇠:동사]+(아)[맺음:어(합침)]", "못가"},
		{"두손으로 들었다.", "두[관형사:두] │ 손[열쇠:체언]+으로[부사격:으로]", "두손+으로[부사격:으로]"},
		{"좀더 빨리.", "좀[열쇠:부사] │ 더[열쇠:부사]", "좀더"},
		{"내일 갈거야.", "가[열쇠:동사]+(ㄹ)[꾸밈:ㄹ] │ 거[의존명사:거]+(이)[서술격:이(빠짐)]+야[맺음:야]", "갈거"},
	} {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		if !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
		if !netHas(r, c.in, c.behind) {
			t.Errorf("%s: the new word %s is no candidate", c.in, c.behind)
		}
	}
	if !netHas(r, "두루미", "두루미") {
		t.Fatal("두루미: no known word after 두, the new word stays")
	}
}
