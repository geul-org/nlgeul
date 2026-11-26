//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNamedQuoted — end to end with the word table of testdata, a case particle written onto the final closing of a sentence quoted as it stands is a piece of the same word (가+(ㄴ)다+를, 먹+는다+의, 가+(ㅂ)니다+가, 학교+이+다+를); not after a copula stem dropped after a vowel (학교다를), and 다 with 가 is the closing 다가 only
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNamedQuoted(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"간다를":   "\n가[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]+를[목적격:를]\n",
		"먹는다의":  "\n먹[열쇠:동사]+는다[맺음:는다]+의[관형격:의]\n",
		"갑니다가":  "\n가[열쇠:동사]+(ㅂ)니다[맺음:ㅂ니다]+가[주격:가]\n",
		"학교이다를": "\n학교[열쇠:체언]+이[서술격:이]+다[맺음:다]+를[목적격:를]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "학교다를"); strings.Contains(got, "다[맺음:다]+를") {
		t.Fatal("a case particle after a dropped copula stem\n" + got)
	}
	if got := streamText(r, "먹다가"); strings.Contains(got, "다[맺음:다]+가") {
		t.Fatal("다 and 가 are the closing 다가\n" + got)
	}
}
