//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAskClosingsWonder — the closings ㄹ는지 and 을는지 take the particles of askParticles as ㄹ지 and 는지 do: 갔을는지도, 갈는지를, 먹었을는지가
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAskClosingsWonder(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	for _, c := range [][2]string{{"맺음:ㄹ는지", "도"}, {"맺음:을는지", "를"}, {"맺음:을는지", "가"}} {
		if v, _ := verdictOf(ix.askAux.g, rule.Piece{Form: c[1], Before: c[0], HasStem: true, Next: -1}, "", false); v <= 0 {
			t.Errorf("%s takes %s", c[0], c[1])
		}
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	for in, want := range map[string]string{
		"갔을는지도":  "가[열쇠:동사]+(았)[선어말:ㅆ]+을는지[맺음:을는지]+도[보조사:도]",
		"갈는지를":   "가[열쇠:동사]+(ㄹ)는지[맺음:ㄹ는지]+를[목적격:를]",
		"먹었을는지가": "먹[열쇠:동사]+었[선어말:었]+을는지[맺음:을는지]+가[주격:가]",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
}
