//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLoneUnit — end to end with the small table of testdata: a unit written in letters that are no Hangul syllables and apart from its number (700 m, 5 km) keeps the reading of the word table first (m[열쇠:체언]) and has the readings as a unit among the candidates (m[단위:미터], m[단위:분], km[단위:킬로미터]); a unit in Hangul syllables apart from its number (3 명) is not read so
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamLoneUnit(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence, line string
		first, want    bool
	}{
		{"700 m 높다", "m[열쇠:체언]", true, true},
		{"700 m 높다", "m[단위:미터]", false, true},
		{"700 m 높다", "m[단위:분]", false, true},
		{"5 km 높다", "km[열쇠:체언]", true, true},
		{"5 km 높다", "km[단위:킬로미터]", false, true},
		{"3 명 높다", "[단위:", false, false},
	}
	for _, c := range cases {
		text := "\n" + streamText(r, c.sentence)
		if c.first {
			text = "\n" + strings.Join(bestLines(r, c.sentence), "\n") + "\n"
		}
		if got := strings.Contains(text, c.line); got != c.want {
			t.Errorf("%s: %q read %v, want %v (first %v)\n%s", c.sentence, c.line, got, c.want, c.first, text)
		}
	}
}
