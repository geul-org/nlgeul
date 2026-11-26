//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamSplit — with a word table an eojeol written with the letters of a closed-set word reads as that word first and as the table's predicate stem and its closing beside it (하지만: 하+지만, 그리고: 그리+고), also with a particle or a sign after it; after -기는 too the word 하지만 stays first (the case exception table), after -어야 the stem and closing are first; with no word table the stem and closing are not read
package parse

import (
	"slices"
	"testing"
)

func TestReaderStreamSplit(t *testing.T) {
	r := streamReader(t, "words-split.tsv")
	cases := []struct{ sentence, letters, first, also string }{
		{"하지만 달이 밝다.", "하지만", "하지만[접속 부사:하지만]", "하[열쇠:동사]+지만[맺음:지만]"},
		{"그리고 달이 밝다.", "그리고", "그리고[접속 부사:그리고]", "그리[열쇠:동사]+고[맺음:고]"},
		{"밝다 하지만, 달이 밝다.", "하지만", "하지만[접속 부사:하지만]", "하[열쇠:동사]+지만[맺음:지만]"},
		{"달을 그리고 하다.", "그리고", "그리고[접속 부사:그리고]", "그리[열쇠:동사]+고[맺음:고]"},
		{"달이 밝기는 하지만 밝다.", "하지만", "하지만[접속 부사:하지만]", "하[열쇠:동사]+지만[맺음:지만]"},
		{"달이 밝아야 하지만 밝다.", "하지만", "하[열쇠:동사]+지만[맺음:지만]", "하지만[접속 부사:하지만]"},
	}
	for _, c := range cases {
		if got := stretchFirst(r, c.sentence, c.letters); got != c.first {
			t.Errorf("%s: first %q, want %q", c.sentence, got, c.first)
		}
		if lines := streamLines(r, c.sentence); !slices.Contains(lines, c.also) {
			t.Errorf("%s: no %q in %q", c.sentence, c.also, lines)
		}
	}
	if lines := streamLines(streamReader(t, ""), "하지만 달이 밝다."); slices.Contains(lines, "하+지만[맺음:지만]") {
		t.Fatal("with no word table the closed-set word alone", lines)
	}
}
