//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamLinkedCopula — end to end with the word table of testdata, the copula right after the linking 어서 is a piece of the same word and the stem before it a verb (먹+어서+이+다, 가+(아)서+이+(ㄹ)까), also with its stem dropped (하+(여)서+(이)+다), while the copula after 고 is not read onto the closing
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamLinkedCopula(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹어서이다": "\n먹[열쇠:동사]+어서[맺음:어서]+이[서술격:이]+다[맺음:다]\n",
		"가서일까":  "\n가[열쇠:동사]+(아)서[맺음:어서(합침)]+이[서술격:이]+(ㄹ)까[맺음:ㄹ까]\n",
		"해서다":   "\n하[열쇠:동사]+(여)서[맺음:어서(합침)]+(이)[서술격:이(빠짐)]+다[맺음:다]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "먹고이다"); strings.Contains(got, "고[맺음:고]+이[서술격:이]") {
		t.Fatal("the copula after 고\n" + got)
	}
}
