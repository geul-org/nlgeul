//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamShortEndings — end to end with the word table of testdata, the short endings added to the end of closings-more.tsv: 매 (으매 after a consonant), 람 after the copula's stem, the linking 든, 어다 (and 아다, and fused into the stem's last syllable: 가+(아)다)
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamShortEndings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"가매":   "\n가[열쇠:동사]+매[맺음:매]\n",
		"먹으매":  "\n먹[열쇠:동사]+으매[맺음:으매]\n",
		"학교이람": "\n학교[열쇠:체언]+이[서술격:이]+람[맺음:람]\n",
		"먹든":   "\n먹[열쇠:동사]+든[맺음:든]\n",
		"먹어다":  "\n먹[열쇠:동사]+어다[맺음:어다]\n",
		"가다":   "\n가[열쇠:동사]+(아)다[맺음:어다(합침)]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
}
