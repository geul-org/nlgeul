//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNamed — end to end with the word table of testdata, the auxiliary 야 written onto an ending is a piece of the same word: after the linking 다면 that ends in a consonant (가+(ㄴ)다면+야), after the closing 다 (먹+겠+다+야) and after the copula's 라 (학교+(이)+라+야); right after a naming closing that ends in a consonant it is not read (먹음야)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNamed(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"간다면야": "\n가[열쇠:동사]+(ㄴ)다면[맺음:ㄴ다면]+야[보조사:야]\n",
		"먹겠다야": "\n먹[열쇠:동사]+겠[선어말:겠]+다[맺음:다]+야[보조사:야]\n",
		"학교라야": "\n학교[열쇠:체언]+(이)[서술격:이(빠짐)]+라[맺음:라]+야[보조사:야]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "먹음야"); strings.Contains(got, "야[보조사:야]") {
		t.Fatal("야 after a naming closing\n" + got)
	}
}
