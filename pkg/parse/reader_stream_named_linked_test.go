//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNamedLinked — end to end with the word table of testdata, the plural 들 written onto a linking closing 어, 게, 지 or 고 is a piece of the same word (먹+어+들, 먹+고+들, 마시+(어)+들 with the closing fused into the stem); after a linking closing that is not of its row it is not read so (먹으면들)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNamedLinked(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹어들": "\n먹[열쇠:동사]+어[맺음:어]+들[접미:들]\n",
		"먹고들": "\n먹[열쇠:동사]+고[맺음:고]+들[접미:들]\n",
		"마셔들": "\n마시[열쇠:동사]+(어)[맺음:어(합침)]+들[접미:들]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "먹으면들"); strings.Contains(got, "면[맺음:으면]+들") || strings.Contains(got, "면[맺음:면]+들") {
		t.Fatal("들 after a closing that is not of its row\n" + got)
	}
}
