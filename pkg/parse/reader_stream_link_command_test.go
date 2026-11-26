//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLinkCommand — the commanding closing 라 takes the link vowel 으 after a stem's final consonant (먹으라: 먹+으+라, 받으라), as the tables' 으라고 does; after a stem with no final consonant 라 stands alone (가라)
package parse

import (
	"slices"
	"testing"
)

func TestReaderStreamLinkCommand(t *testing.T) {
	r := streamReader(t, "words-hao.tsv")
	for _, c := range []struct{ sentence, want string }{
		{"물을 먹으라!", "먹[열쇠:동사]+으[이음:으]+라[맺음:라]"},
		{"결과를 받으라.", "받[열쇠:동사]+으[이음:으]+라[맺음:라]"},
		{"학교에 가라.", "가[열쇠:동사]+라[맺음:라]"},
	} {
		if lines := streamLines(r, c.sentence); !slices.Contains(lines, c.want) {
			t.Errorf("%s: no %q in %q", c.sentence, c.want, lines)
		}
	}
	if !slices.Contains(linkTakers, "맺음:라") {
		t.Fatal(linkTakers)
	}
}
