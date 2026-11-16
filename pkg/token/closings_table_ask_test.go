//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestGatedClosingsAsk — the asking closing 누 of gated-closings.tsv is read after a predicate's stem of the word table (먹+었+누, 가+누) and not without a word table, where nothing tells the stem
package token

import (
	"strings"
	"testing"
)

func TestGatedClosingsAsk(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹었누": "먹[열쇠:동사]+었[선어말:었]+누[맺음:누]",
		"가누":  "가[열쇠:동사]+누[맺음:누]",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	bare, _ := New(Options{})
	if got := streamText(bare, "먹었누"); strings.Contains(got, "[맺음:누]") {
		t.Fatal("누 without a word table\n" + got)
	}
}
