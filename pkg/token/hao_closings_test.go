//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestHaoClosings — end to end with a word table, the 하오체 closings stand after a predicate's stem of the table (먹소, 가오, 받으오, 하겠소, 학교이오, 아니올시다), 가시오 and 먹으시오 are read as the honorific and 오, and the letters of nouns (장소를, 연구소에, 감소는, 라디오를) and a stem the table does not have (접골시술소) are not read as a closing 소 or 오
package token

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
		{"먹소", "  먹[열쇠:동사]+소[맺음:소]\t"},
		{"하겠소", "  하[열쇠:동사]+겠[선어말:겠]+소[맺음:소]\t"},
		{"가오", "  가[열쇠:동사]+오[맺음:오]\t"},
		{"받으오", "  받[열쇠:동사]+으오[맺음:으오]\t"},
		{"학교이오", "  학교[열쇠:체언]+이[서술격:이]+오[맺음:오]\t"},
		{"아니올시다", "  아니[열쇠:형용사]+올시다[맺음:올시다]\t"},
		{"가시오", "  가[열쇠:동사]+시[선어말:시]+오[맺음:오]\t"},
		{"먹으시오", "  먹[열쇠:동사]+으[이음:으]+시[선어말:시]+오[맺음:오]\t"},
	} {
		lines := r.Lines(c.in, false)
		if len(lines) < 2 || !strings.HasPrefix(lines[1], c.want) {
			t.Errorf("%s: want first %q\n%s", c.in, c.want, strings.Join(lines, "\n"))
		}
	}
	for _, c := range []struct{ in, want string }{
		{"장소를", "  장소[열쇠:체언]+를[목적격:를]\t"},
		{"연구소에", "  연구소[열쇠:체언]+에[부사격:에]\t"},
		{"감소는", "  감소[열쇠:체언]+는[보조사:는]\t"},
		{"라디오를", "  라디오[열쇠:체언]+를[목적격:를]\t"},
	} {
		lines := r.Lines(c.in, false)
		if len(lines) < 2 || !strings.HasPrefix(lines[1], c.want) || strings.Contains(strings.Join(lines, "\n"), "맺음:소]") {
			t.Errorf("%s: read as a closing\n%s", c.in, strings.Join(lines, "\n"))
		}
	}
	if got := strings.Join(r.Lines("접골시술소", false), "\n"); strings.Contains(got, "맺음:소]") {
		t.Errorf("접골시술소: a closing after a stem not in the table\n%s", got)
	}
}
