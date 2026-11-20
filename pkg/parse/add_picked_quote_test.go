//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddPickedQuote — the graph of the quoting 라고 takes the closing 다, ㄴ다 or 는다 before it (a sentence quoted with no quotation marks) and turns down 고 and 지; with the word table of testdata 간다라고, 먹는다라고 and 먹었다라고 read the particle right after the closing
package parse

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddPickedQuote(t *testing.T) {
	ix := &index{}
	ix.addPicked()
	e := ix.pickedAux["라고"]
	if e == nil || e.name != pickedAuxName+"라고" {
		t.Fatal("the graph of 라고", e)
	}
	for before, want := range map[string]bool{"맺음:다": true, "맺음:ㄴ다": true, "맺음:는다": true, "맺음:고": false, "맺음:지": false} {
		if v, _ := verdictOf(e.g, rule.Piece{Form: "라고", Before: before, HasStem: true, Next: -1}, "", false); (v > 0) != want {
			t.Errorf("%s takes 라고: %v", before, v)
		}
	}
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"간다라고":  "가[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]+라고[부사격:라고]",
		"먹는다라고": "먹[열쇠:동사]+는다[맺음:는다]+라고[부사격:라고]",
		"먹었다라고": "먹[열쇠:동사]+었[선어말:었]+다[맺음:다]+라고[부사격:라고]",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
}
