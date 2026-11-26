//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLinkLiquid — with a word table the link vowel 으 written after a predicate stem's ㄹ against the orthography is read as a piece (살으시오: 살+으+시+오, 살으니), with a verdict below 1; a stem that changed ㄷ to ㄹ keeps its own candidate (들으니: 듣+으+니) with a higher verdict; with no word table 으 is not read after ㄹ
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderStreamLinkLiquid(t *testing.T) {
	r := streamReader(t, "words-hao.tsv")
	for _, c := range []struct{ sentence, want string }{
		{"여기 살으시오.", "살[열쇠:동사]+으[이음:으]+시[선어말:시]+오[맺음:오]"},
		{"여기 살으니", "살[열쇠:동사]+으[이음:으]+니[맺음:니]"},
		{"결과를 들으니", "듣[열쇠:동사]+으[이음:으]+니[맺음:니]"},
	} {
		if lines := streamLines(r, c.sentence); !slices.Contains(lines, c.want) {
			t.Errorf("%s: no %q in %q", c.sentence, c.want, lines)
		}
	}
	scored := strings.Join(scoredLines(r, "여기 살으니"), "\n")
	if !strings.Contains(scored, "살[열쇠:동사]+으[이음:으]+니[맺음:니]\t성립 0.") {
		t.Fatal("the form against the orthography is rebutted", scored)
	}
	if got := strings.Join(streamLines(streamReader(t, ""), "여기 살으니"), "\n"); strings.Contains(got, "이음:으") {
		t.Fatal("no word table", got)
	}
}
