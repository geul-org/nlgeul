//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestOldParticles — end to end with a word table, the auxiliary particles of the dictionary added at the end of the particle table stand after a nominal: 라야 after a vowel (학교라야) and 인즉 after a final consonant (마당인즉) first; 이라야 after a final consonant is read (마당이라야), beside the suffix 이 of a name with 라야 of the same lambda
package token

import (
	"strings"
	"testing"
)

func TestOldParticles(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-hao.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ in, want string }{
		{"학교라야", "학교[열쇠:체언]+라야[보조사:라야]"},
		{"마당인즉", "마당[열쇠:체언]+인즉[보조사:인즉]"},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.want {
			t.Errorf("%s: want first %q\n%s", c.in, c.want, strings.Join(lines, "\n"))
		}
	}
	if got := strings.Join(streamLines(r, "마당이라야"), "\n"); !strings.Contains(got, "마당[열쇠:체언]+이라야[보조사:이라야]") {
		t.Errorf("마당이라야: no 이라야\n%s", got)
	}
}
