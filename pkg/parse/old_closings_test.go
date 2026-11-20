//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestOldClosings — end to end with a word table, the closings of the dictionary added at the end of the closing table stand after a predicate's stem or the copula's: ㄹ소냐 and 을소냐 (갈소냐, 먹을소냐), 노라 (가노라), 니라 (학교이니라), 오니 (학교이오니) and ㄴ다더니 (간다더니)
package parse

import (
	"strings"
	"testing"
)

func TestOldClosings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-hao.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ in, want string }{
		{"갈소냐", "가[열쇠:동사]+(ㄹ)소냐[맺음:ㄹ소냐]"},
		{"먹을소냐", "먹[열쇠:동사]+을소냐[맺음:을소냐]"},
		{"가노라", "가[열쇠:동사]+노라[맺음:노라]"},
		{"학교이니라", "학교[열쇠:체언]+이[서술격:이]+니라[맺음:니라]"},
		{"학교이오니", "학교[열쇠:체언]+이[서술격:이]+오니[맺음:오니]"},
		{"간다더니", "가[열쇠:동사]+(ㄴ)다더니[맺음:ㄴ다더니]"},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.want {
			t.Errorf("%s: want first %q\n%s", c.in, c.want, strings.Join(lines, "\n"))
		}
	}
}
