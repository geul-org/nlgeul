//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCaseExceptions — end to end with a case exception table of testdata and a word table: a word a row names is not read with the closing of the row's graph (먹소, 가오 lose 맺음:소 and 맺음:오, and so does 먹소 before a particle), the trace says why with the word and the row (line, kind of ground, ground), and the same closings stay after other words (하겠소, 먹었소, 가시오); a graph the table names that is no marker graph of the index is an error with a word table and none without one, and the table held in the program builds an index
package parse

import (
	"strings"
	"testing"
)

func TestCaseExceptions(t *testing.T) {
	if _, err := New(Options{Words: "testdata/words-hao.tsv"}); err != nil {
		t.Fatal("the table held in the program", err)
	}
	useCases(t, "testdata/case-exceptions.tsv")
	r, err := New(Options{Words: "testdata/words-hao.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ in, closing, note string }{
		{"먹소", "맺음:소]", "  · 소\t맺음:소\t-1.00\t기각: 예외 표의 줄이다(먹소) — 예외 표 3: 관찰 시험 3/3"},
		{"먹소를", "맺음:소]", "  · 소\t맺음:소\t-1.00\t기각: 예외 표의 줄이다(먹소) — 예외 표 3: 관찰 시험 3/3"},
		{"가오", "맺음:오]", "  · 오\t맺음:오\t-1.00\t기각: 예외 표의 줄이다(가오) — 예외 표 4: 규범 시험 조항"},
	} {
		got := strings.Join(r.BoxLines(c.in, true, 0), "\n")
		if strings.Contains(got, c.closing) || !strings.Contains(got, c.note+"\n") {
			t.Errorf("%s:\n%s", c.in, got)
		}
	}
	for _, c := range []struct{ in, want string }{
		{"하겠소", "하[열쇠:동사]+겠[선어말:겠]+소[맺음:소]"},
		{"먹었소", "먹[열쇠:동사]+었[선어말:었]+소[맺음:소]"},
		{"가시오", "가[열쇠:동사]+시[선어말:시]+오[맺음:오]"},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.want {
			t.Errorf("%s: want first %q\n%s", c.in, c.want, strings.Join(lines, "\n"))
		}
	}
	useCases(t, "testdata/case-exceptions-bad.tsv")
	if _, err := New(Options{Words: "testdata/words-hao.tsv"}); err == nil || err.Error() != `표 case-exceptions-bad.tsv 3번째 줄 1번째 칸(그래프) "맺음:소오": 색인에 이 이름의 표지 그래프도 문장 규칙도 없다` {
		t.Fatal(err)
	}
	if _, err := New(Options{}); err != nil {
		t.Fatal("without a word table the names are not asked", err)
	}
}
