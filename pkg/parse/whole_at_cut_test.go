//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestWholeAtCut — where a path ends in a slash before Hangul the path stands as a word of no markers (the slash shows the boundary: ~/a/b/[경로] and then a word), and the whole unit stands as a path with that Hangul as its name beside the new word of the same letters; a path with a name's letter before the Hangul has no such word, and a path is not cut from Hangul that reads as markers on it (~/a/b/에)
package parse

import (
	"strings"
	"testing"
)

func TestWholeAtCut(t *testing.T) {
	r := streamReader(t, "words.tsv")
	lines := "\n" + strings.Join(streamLines(r, "~/a/b/결과는 좋다"), "\n") + "\n"
	for _, want := range []string{"\n~/a/b/결과는[경로]\n", "\n~/a/b/[경로] 결과[열쇠:체언]+는[보조사:는]\n", "\n~/a/b/결과는\n", "\n~/a/b/결과+는[보조사:는]\n"} {
		if !strings.Contains(lines, want) {
			t.Errorf("%q is not in%s", want, lines)
		}
	}
	lines = "\n" + strings.Join(streamLines(r, "~/a/b/에 둔다"), "\n") + "\n"
	if strings.Contains(lines, "\n~/a/b/[경로]\n") || !strings.Contains(lines, "\n~/a/b/[경로]+에[부사격:에]\n") {
		t.Error(lines)
	}
	lines = "\n" + strings.Join(streamLines(r, "~/a/b결과는 좋다"), "\n") + "\n"
	if strings.Contains(lines, "결과는[경로]") || !strings.Contains(lines, "\n~/a/b[경로]\n결과[열쇠:체언]+는[보조사:는]\n") {
		t.Error(lines)
	}
}
