//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamEmph — end to end with the word table of testdata, 를 of emphasis after 에 and 으로 is a piece of the same word (고양이+에+를, 고양이+로+를), not after 에서 or 서, nor right on a stem, and the particle 일랑 is read after a final consonant (책+일랑)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamEmph(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"고양이에를": "\n고양이[열쇠:체언]+에[부사격:에]+를[보조사:를]\n",
		"고양이로를": "\n고양이[열쇠:체언]+로[부사격:로]+를[보조사:를]\n",
		"책일랑":   "\n책+일랑[보조사:일랑]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for _, in := range []string{"고양이에서를", "고양이서를", "고양이를"} {
		if got := streamText(r, in); strings.Contains(got, "[보조사:를]") {
			t.Errorf("%s: 를 of emphasis\n%s", in, got)
		}
	}
}
