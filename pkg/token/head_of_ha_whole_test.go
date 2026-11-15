//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestHeadOfHaWhole — the sentences of clause 40, note 2, read as a stream (the best candidate of each word, by lambda): 생각건대 and 짐작건대 are one word, the stem with 하 and the closing 건대 as it stands; 도착지를 keeps the noun with its particle and has no reading with 하; 넉넉지 is left as it was (지 is not read so), with no reading with 하
package token

import (
	"strings"
	"testing"
)

func TestHeadOfHaWhole(t *testing.T) {
	r := streamReader(t, "words-ha.tsv")
	for _, c := range []struct{ sentence, word, first string }{
		{"생각건대 그 말이 맞다.", "생각건대", "생각하[열쇠:동사]+건대[맺음:건대]"},
		{"짐작건대 그는 떠났다.", "짐작건대", "짐작하[열쇠:동사]+건대[맺음:건대]"},
		{"도착지를 정했다.", "도착지를", "도착지[열쇠:체언]+를[목적격:를]"},
		{"넉넉지 않은 살림이다.", "넉넉지", "넉넉지"},
	} {
		w, ok := jsonWordAt(r.JSONSentence(c.sentence, 0), c.word)
		if !ok {
			t.Errorf("%s: no word %s", c.sentence, c.word)
			continue
		}
		got := jsonWordShown(w)
		if got[0] != c.first {
			t.Errorf("%s: %v, want %s first", c.sentence, got, c.first)
		}
		if !strings.HasSuffix(c.word, "건대") && strings.Contains(strings.Join(got, "\n"), "하[열쇠:") {
			t.Errorf("%s: a reading with 하: %v", c.sentence, got)
		}
	}
}
