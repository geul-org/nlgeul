//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamMoreEndings — end to end with the word table of testdata, the endings added to the end of closings-more.tsv: the imperative 거라 and 너라 after a verb's stem and not after an adjective's (작거라) or the copula's stem; the 하소서체 나이다, 나이까, 더이다 and 소서 (으소서 after a consonant); 노라고, 노라면, 거니와, 던지, 다거나, 렴 (으렴), 오리다 (으오리다), ㄴ걸 after the copula; 로구나 after the copula's stem and not right after a noun or a verb's stem
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamMoreEndings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"가거라":    "\n가[열쇠:동사]+거라[맺음:거라]\n",
		"가너라":    "\n가[열쇠:동사]+너라[맺음:너라]\n",
		"먹었나이다":  "\n먹[열쇠:동사]+었[선어말:었]+나이다[맺음:나이다]\n",
		"먹나이까":   "\n먹[열쇠:동사]+나이까[맺음:나이까]\n",
		"하더이다":   "\n하[열쇠:동사]+더이다[맺음:더이다]\n",
		"먹으소서":   "\n먹[열쇠:동사]+으소서[맺음:으소서]\n",
		"하노라고":   "\n하[열쇠:동사]+노라고[맺음:노라고]\n",
		"하노라면":   "\n하[열쇠:동사]+노라면[맺음:노라면]\n",
		"하거니와":   "\n하[열쇠:동사]+거니와[맺음:거니와]\n",
		"먹던지":    "\n먹[열쇠:동사]+던지[맺음:던지]\n",
		"작다거나":   "\n작[열쇠:형용사]+다거나[맺음:다거나]\n",
		"먹으렴":    "\n먹[열쇠:동사]+으렴[맺음:으렴]\n",
		"가오리다":   "\n가[열쇠:동사]+오리다[맺음:오리다]\n",
		"먹으오리다":  "\n먹[열쇠:동사]+으오리다[맺음:으오리다]\n",
		"학교인걸":   "\n학교[열쇠:체언]+이[서술격:이]+(ㄴ)걸[맺음:ㄴ걸]\n",
		"학교이로구나": "\n학교[열쇠:체언]+이[서술격:이]+로구나[맺음:로구나]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작거라":   "작[열쇠:형용사]+거라",
		"학교이거라": "[서술격:이]+거라",
		"학교로구나": "학교[열쇠:체언]+로구나",
		"먹로구나":  "먹[열쇠:동사]+로구나",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
