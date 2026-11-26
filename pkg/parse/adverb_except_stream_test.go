//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAdverbExceptStream — end to end with the word table of testdata: an adverb reads the copula after it (빨리+이+다), an interjection the particle 요 (아이고+요), each beside the readings there were; an adverb with an auxiliary particle stands as before (빨리+도); a case particle after an adverb is not read on the adverb (빨리가), nor a case particle after an interjection (아이고를)
package parse

import (
	"strings"
	"testing"
)

func TestAdverbExceptStream(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"빨리이다": "\n빨리[열쇠:부사]+이[서술격:이]+다[맺음:다]\n",
		"아이고요": "\n아이고[열쇠:홀로]+요[보조사:요]\n",
		"빨리도":  "\n빨리[열쇠:부사]+도[보조사:도]\n",
		"빨리요":  "\n빨리[열쇠:부사]+요[보조사:요]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"빨리가":  "빨리[열쇠:부사]+가",
		"아이고를": "아이고[열쇠:홀로]+를",
		"아이고도": "아이고[열쇠:홀로]+도",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
	if got := streamText(r, "빨리이다"); !strings.Contains(got, "\n빨리이다\n") {
		t.Fatal("the new word of the same letters stays\n" + got)
	}
}
