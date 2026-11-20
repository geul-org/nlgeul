//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestNumberStopSentence — end to end with the word table of testdata: in a date written with dots the last number's dot with a particle after it is a piece of the number's word, read first (2012. 12. 31.까지: 31[수]+.[마침표]+까지[보조사:까지]), while the dots before it stand alone; the dot of a decimal stays in the number (1.5까지), and a full stop after a word still closes that word (했다.그는)
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestNumberStopSentence(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"2012. 12. 31.까지", []string{"{0 4} 2012[수]", "{4 5} .[마침표]", "{5 7} 12[수]", "{7 8} .[마침표]", "{8 13} 31[수]+.[마침표]+까지[보조사:까지]"}},
		{"1.5까지", []string{"{0 5} 1.5[소수점]+까지[보조사:까지]"}},
	} {
		if got := flatNodes(r, c.in, true); !slices.Equal(got, c.want) {
			t.Errorf("%s: %q, want %q", c.in, got, c.want)
		}
	}
	got := flatNodes(r, "했다.그는", true)
	if len(got) != 2 || got[0] != "{0 3} 하[열쇠:동사]+(였)[선어말:ㅆ]+다[맺음:다]+.[마침표]" || !strings.HasPrefix(got[1], "{3 5} 그") {
		t.Fatalf("했다.그는: %q", got)
	}
}
