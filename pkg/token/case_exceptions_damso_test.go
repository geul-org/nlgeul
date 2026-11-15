//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCaseExceptionsDamso — the row of 담소 in the case exception table held in the program (graph 맺음:소, ground observed): end to end with a word table of testdata, 담소 alone and in a sentence is read first as the reading the row keeps (the noun) and is no longer read with the 하오체 closing 소, which it is without the row; the 하오체 closings after other stems stay first (적소, 하오, 가오)
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestCaseExceptionsDamso(t *testing.T) {
	rows := caseExceptions["맺음:소"]
	if len(rows) != 1 || rows[0].Word != "담소" || rows[0].Reading != "담소[열쇠:체언]" || rows[0].GroundKind != groundSeen || rows[0].Row.Line != 2 {
		t.Fatal(caseExceptions)
	}
	r := streamReader(t, "words-damso.tsv")
	for _, s := range []string{"담소", "두 사람은 오랜 담소 끝에 헤어졌다."} {
		w, ok := jsonWordAt(r.JSONSentence(s, 0), "담소")
		got := jsonWordShown(w)
		if !ok || got[0] != rows[0].Reading || slices.Contains(got, "담[열쇠:동사]+소[맺음:소]") {
			t.Errorf("%s: %v", s, got)
		}
	}
	if got := strings.Join(r.BoxLines("담소", true, 0), "\n"); !strings.Contains(got, "\t맺음:소\t-1.00\t기각: 예외 표의 줄이다(담소) — 예외 표 2: 관찰 46/46, 29파일, 2026-11-15, rules/new-word 2483d7f\n") {
		t.Errorf("the trace names the row:\n%s", got)
	}
	for _, c := range []struct{ sentence, word, first string }{
		{"내 이름을 여기 적소.", "적소", "+소[맺음:소]"},
		{"그만 하오.", "하오", "하[열쇠:동사]+오[맺음:오]"},
		{"나는 이제 가오.", "가오", "가[열쇠:동사]+오[맺음:오]"},
	} {
		w, ok := jsonWordAt(r.JSONSentence(c.sentence, 0), c.word)
		if got := jsonWordShown(w); !ok || !strings.HasSuffix(got[0], c.first) || !strings.HasPrefix(got[0], string([]rune(c.word)[0])+"[열쇠:") {
			t.Errorf("%s: %v, want the 하오체 first", c.sentence, got)
		}
	}
	old := caseExceptions
	t.Cleanup(func() { caseExceptions = old })
	caseExceptions = map[string][]caseException{}
	w, _ := jsonWordAt(streamReader(t, "words-damso.tsv").JSONSentence("담소", 0), "담소")
	if got := jsonWordShown(w); got[0] != "담[열쇠:동사]+소[맺음:소]" {
		t.Fatal("without the row the 하오체 reading comes first", got)
	}
}
