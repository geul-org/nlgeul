//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestHaoClosings — end to end with a word table, the 하오체 closings stand after a predicate's stem of the table (먹소, 가오, 받으오, 하겠소, 학교이오, 아니올시다, and 학교올시다 with the copula's stem dropped after a vowel), 가시오 and 먹으시오 are read as the honorific and 오, and the letters of nouns (장소를, 연구소에, 감소는, 라디오를) and a stem the table does not have (접골시술소) are not read as a closing 소 or 오
package parse

import (
	"strings"
	"testing"
)

func TestHaoClosings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-hao.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ in, want string }{
		{"먹소", "먹[열쇠:동사]+소[맺음:소]"},
		{"하겠소", "하[열쇠:동사]+겠[선어말:겠]+소[맺음:소]"},
		{"가오", "가[열쇠:동사]+오[맺음:오]"},
		{"받으오", "받[열쇠:동사]+으오[맺음:으오]"},
		{"학교이오", "학교[열쇠:체언]+이[서술격:이]+오[맺음:오]"},
		{"아니올시다", "아니[열쇠:형용사]+올시다[맺음:올시다]"},
		{"학교올시다", "학교[열쇠:체언]+(이)[서술격:이(빠짐)]+올시다[맺음:올시다]"},
		{"가시오", "가[열쇠:동사]+시[선어말:시]+오[맺음:오]"},
		{"먹으시오", "먹[열쇠:동사]+으[이음:으]+시[선어말:시]+오[맺음:오]"},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.want {
			t.Errorf("%s: want first %q\n%s", c.in, c.want, strings.Join(lines, "\n"))
		}
	}
	for _, c := range []struct{ in, want string }{
		{"장소를", "장소[열쇠:체언]+를[목적격:를]"},
		{"연구소에", "연구소[열쇠:체언]+에[부사격:에]"},
		{"감소는", "감소[열쇠:체언]+는[보조사:는]"},
		{"라디오를", "라디오[열쇠:체언]+를[목적격:를]"},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.want || strings.Contains(strings.Join(lines, "\n"), "맺음:소]") {
			t.Errorf("%s: read as a closing\n%s", c.in, strings.Join(lines, "\n"))
		}
	}
	if got := streamText(r, "접골시술소"); strings.Contains(got, "맺음:소]") {
		t.Errorf("접골시술소: a closing after a stem not in the table\n%s", got)
	}
}
