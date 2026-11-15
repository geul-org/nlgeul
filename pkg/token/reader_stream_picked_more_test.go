//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamPickedMore — end to end with the word table of testdata, 와 written onto the linking 어서 is a piece of the same word (먹+어서+와), not onto 고 (먹고와); 의 right after a closing that asks is one too (가+는가+의, 먹+는지+의)
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamPickedMore(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹어서와": "\n먹[열쇠:동사]+어서[맺음:어서]+와[부사격:와]\n",
		"가는가의": "\n가[열쇠:동사]+는가[맺음:는가]+의[관형격:의]\n",
		"먹는지의": "\n먹[열쇠:동사]+는지[맺음:는지]+의[관형격:의]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "먹고와"); strings.Contains(got, "고[맺음:고]+와") {
		t.Fatal("와 after 고\n" + got)
	}
}
