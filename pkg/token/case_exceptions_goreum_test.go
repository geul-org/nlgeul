//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCaseExceptionsGoreum — the row of 고름 in the case exception table held in the program (graph 이름:ㅁ, ground observed): end to end with a word table of testdata, 고름 before a particle standing as a unit of its own is read first as the reading the row keeps (the noun) and no longer as 고르 with the naming closing ㅁ, which it is without the row; the row strikes, so 고름 written as the naming form of 고르다 is read as the noun too (옷을 고름에); the naming closing ㅁ after other stems stays first (됨을)
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestCaseExceptionsGoreum(t *testing.T) {
	rows := caseExceptions["이름:ㅁ"]
	if len(rows) != 1 || rows[0].Word != "고름" || rows[0].Reading != "고름[열쇠:체언]" || rows[0].GroundKind != groundSeen || rows[0].Row.Line != 3 {
		t.Fatal(caseExceptions)
	}
	r := streamReader(t, "words-goreum.tsv")
	for _, c := range []struct{ sentence, word, first, gone string }{
		{"고름을 짰다", "고름을", "고름[열쇠:체언]+을[목적격:을]", "고르[열쇠:동사]+[이름:ㅁ]+을[목적격:을]"},
		{"저고리 고름에 손이 갔다", "고름에", "고름[열쇠:체언]+에[부사격:에]", "고르[열쇠:동사]+[이름:ㅁ]+에[부사격:에]"},
		{"옷을 고름에 시간이 걸렸다", "고름에", "고름[열쇠:체언]+에[부사격:에]", "고르[열쇠:동사]+[이름:ㅁ]+에[부사격:에]"},
	} {
		w, ok := jsonWordAt(r.JSONSentence(c.sentence, 0), c.word)
		got := jsonWordShown(w)
		if !ok || got[0] != c.first || slices.Contains(got, c.gone) {
			t.Errorf("%s: %v", c.sentence, got)
		}
	}
	if got := strings.Join(r.BoxLines("고름을", true, 0), "\n"); !strings.Contains(got, "\t이름:ㅁ\t-1.00\t기각: 예외 표의 줄이다(고름) — 예외 표 3: 관찰 30/30, 고름+조사 57자리 25파일 중 30, 2026-11-15, rules/new-word b3cb6af\n") {
		t.Errorf("the trace names the row:\n%s", got)
	}
	w, ok := jsonWordAt(r.JSONSentence("무효가 됨을 알았다", 0), "됨을")
	if got := jsonWordShown(w); !ok || got[0] != "되[열쇠:동사]+[이름:ㅁ]+을[목적격:을]" {
		t.Errorf("됨을: %v, want the naming closing first", got)
	}
	old := caseExceptions
	t.Cleanup(func() { caseExceptions = old })
	caseExceptions = map[string][]caseException{}
	w, _ = jsonWordAt(streamReader(t, "words-goreum.tsv").JSONSentence("고름을 짰다", 0), "고름을")
	if got := jsonWordShown(w); got[0] != "고르[열쇠:동사]+[이름:ㅁ]+을[목적격:을]" {
		t.Fatal("without the row the naming closing's reading comes first", got)
	}
}
