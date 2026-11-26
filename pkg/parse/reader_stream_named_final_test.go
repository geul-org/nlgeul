//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNamedFinal — end to end with the word table of testdata, 만 and 그려 written onto a final closing are pieces of the same word (먹+습니다+만, 먹+습니다+그려), and so is the 가 of emphasis after the linking 지 (먹+지+가); 그려 and that 가 are read nowhere else: not after 지 or 고, not right on a nominal
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNamedFinal(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹습니다만":  "\n먹[열쇠:동사]+습니다[맺음:습니다]+만[보조사:만]\n",
		"먹습니다그려": "\n먹[열쇠:동사]+습니다[맺음:습니다]+그려[보조사:그려]\n",
		"먹지가":    "\n먹[열쇠:동사]+지[맺음:지]+가[보조사:가]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for _, in := range []string{"먹지그려", "학교그려", "먹고가", "학교가"} {
		if got := streamText(r, in); strings.Contains(got, "그려[보조사:그려]") || strings.Contains(got, "가[보조사:가]") {
			t.Errorf("%s: a particle read only after the closings of its row\n%s", in, got)
		}
	}
}
