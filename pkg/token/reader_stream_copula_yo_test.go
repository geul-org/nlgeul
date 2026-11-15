//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamCopulaYo — end to end with the word table of testdata, the linking ending 요 added to the end of closings-more.tsv stands after the copula's stem 이 and after 아니, and not after a verb's stem or right after a noun with the copula's stem dropped
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamCopulaYo(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"학교이요": "\n학교[열쇠:체언]+이[서술격:이]+요[맺음:요]\n",
		"문제이요": "\n문제[열쇠:체언]+이[서술격:이]+요[맺음:요]\n",
		"아니요":  "\n아니+요[맺음:요]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"먹요":  "[맺음:요]",
		"학교요": "[맺음:요]",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
