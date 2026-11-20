//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamCopulaParticle — end to end with the word table of testdata, the copula's stem stands after the particles 대로, 까지 and 부터, dropped after a vowel (문제+대로+(이)+지, 문제+대로+(이)+다, 문제+까지+(이)+다) said after a final consonant (문제+부터+이+다), and after the particle 뿐 written onto a nominal (문제+뿐+이+다)
package parse

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
		"문제대로지":  "\n문제[열쇠:체언]+대로[부사격:대로]+(이)[서술격:이(빠짐)]+지[맺음:지]\n",
		"문제대로다":  "\n문제[열쇠:체언]+대로[부사격:대로]+(이)[서술격:이(빠짐)]+다[맺음:다]\n",
		"문제까지다":  "\n문제[열쇠:체언]+까지[보조사:까지]+(이)[서술격:이(빠짐)]+다[맺음:다]\n",
		"문제부터이다": "\n문제[열쇠:체언]+부터[보조사:부터]+이[서술격:이]+다[맺음:다]\n",
		"문제뿐이다":  "\n문제[열쇠:체언]+뿐[보조사:뿐]+이[서술격:이]+다[맺음:다]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
}
