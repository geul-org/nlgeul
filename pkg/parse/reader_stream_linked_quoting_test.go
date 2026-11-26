//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLinkedQuoting — end to end with the word table of testdata, after the linking 어서 of a predicate the copula's stem is read as dropped before a quoting closing too (먹어서라는, 먹어서라, 해서라는), as it is before 다 (먹어서다): the gate of a nominal or a new word that such a closing asks of the word before the stem is not asked of the predicate's stem
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamLinkedQuoting(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹어서라는": "먹[열쇠:동사]+어서[맺음:어서]+(이)[서술격:이(빠짐)]+라는[꾸밈:라는]",
		"먹어서라":  "먹[열쇠:동사]+어서[맺음:어서]+(이)[서술격:이(빠짐)]+라[맺음:라]",
		"해서라는":  "하[열쇠:동사]+(여)서[맺음:어서(합침)]+(이)[서술격:이(빠짐)]+라는[꾸밈:라는]",
		"먹어서다":  "먹[열쇠:동사]+어서[맺음:어서]+(이)[서술격:이(빠짐)]+다[맺음:다]",
	} {
		if got := streamText(r, in); !strings.Contains(got, "\n"+want+"\n") {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
}
