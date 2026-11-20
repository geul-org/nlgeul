//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDative — end to end with the word table of testdata, the particle 더러 (the person one speaks to) is read after a noun (우리+더러, 고양이+더러+는) and not after a verb's stem (먹더러)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDative(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"우리더러":   "\n우리[열쇠:체언]+더러[부사격:더러]\n",
		"고양이더러는": "\n고양이[열쇠:체언]+더러[부사격:더러]+는[보조사:는]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "먹더러"); strings.Contains(got, "먹[열쇠:동사]+더러") {
		t.Errorf("더러 after a verb's stem\n%s", got)
	}
}
