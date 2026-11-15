//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamEdgeSlash — end to end with no word table: a slash at a unit's edge is /[빗금] after Latin letters (before a comma too), after a closing quote and alone; a slash between Latin letters stays inside the path (a/b[경로]) with no 빗금
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamEdgeSlash(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ sentence, has, not string }{
		{"docs/, tests/ 폴더", "\n/[빗금]\n", "\n/\n"},
		{"\"좋다\"/ 나쁘다", "\n/[빗금]\n", "\n/\n"},
		{"가 / 나", "\n/[빗금]\n", "\n/\n"},
		{"ab/cd를 연다.", "ab/cd[경로]+를[목적격:를]", "[빗금]"},
	}
	for _, c := range cases {
		got := streamText(r, c.sentence)
		if !strings.Contains(got, c.has) || strings.Contains(got, c.not) {
			t.Errorf("%q: want %q and no %q\n%s", c.sentence, c.has, c.not, got)
		}
	}
}
