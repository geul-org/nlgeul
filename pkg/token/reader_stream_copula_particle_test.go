//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamCopulaParticle — end to end with the word table of testdata, the copula's stem stands after the particle 대로, dropped after its vowel (문제+대로+(이)+지, 문제+대로+(이)+다), and not after 뿐 or 까지 (문제뿐이다, 문제까지다)
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamCopulaParticle(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"문제대로지": "\n문제[열쇠:체언]+대로[부사격:대로]+(이)[서술격:이(빠짐)]+지[맺음:지]\n",
		"문제대로다": "\n문제[열쇠:체언]+대로[부사격:대로]+(이)[서술격:이(빠짐)]+다[맺음:다]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"문제뿐이다": "뿐[보조사:뿐]+이[서술격:이]",
		"문제까지다": "까지[보조사:까지]+(이)[서술격:이(빠짐)]",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
