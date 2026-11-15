//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAskClosingsQuestion — the closings 냐, 느냐 and 는가 take the particles of askParticles as 는지 does, 의 among them (는가+의), not 과 or 로: 옳으냐에, 버티느냐가, 하느냐를, 가느냐도; 옳은가에 is not read so (ㄴ가 is not among them)
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAskClosingsQuestion(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	for _, c := range [][2]string{{"맺음:냐", "에"}, {"맺음:느냐", "가"}, {"맺음:느냐", "를"}, {"맺음:는가", "에"}, {"맺음:냐", "와"}, {"맺음:는가", "도"}, {"맺음:는가", "의"}} {
		if v, _ := verdictOf(ix.askAux.g, rule.Piece{Form: c[1], Before: c[0], HasStem: true, Next: -1}, "", false); v <= 0 {
			t.Errorf("%s takes %s", c[0], c[1])
		}
	}
	for _, c := range [][2]string{{"맺음:느냐", "과"}, {"맺음:냐", "로"}, {"맺음:ㄴ가", "에"}} {
		if v, _ := verdictOf(ix.askAux.g, rule.Piece{Form: c[1], Before: c[0], HasStem: true, Next: -1}, "", false); v > 0 {
			t.Errorf("%s takes no %s", c[0], c[1])
		}
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["옳"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	r.ix.lex["버티"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	for in, want := range map[string]string{
		"옳으냐에":  "옳[열쇠:형용사]+으[이음:으]+냐[맺음:냐]+에[부사격:에]",
		"버티느냐가": "버티[열쇠:동사]+느냐[맺음:느냐]+가[주격:가]",
		"하느냐를":  "하[열쇠:동사]+느냐[맺음:느냐]+를[목적격:를]",
		"가느냐도":  "가[열쇠:동사]+느냐[맺음:느냐]+도[보조사:도]",
		"가는가에":  "가[열쇠:동사]+는가[맺음:는가]+에[부사격:에]",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	if got := streamText(r, "옳은가에"); strings.Contains(got, "은가[맺음:은가]+에") || strings.Contains(got, "+(ㄴ)가[맺음:ㄴ가]+에") {
		t.Fatal("ㄴ가 and 은가 take no particle\n" + got)
	}
}
