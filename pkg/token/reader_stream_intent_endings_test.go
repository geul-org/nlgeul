//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamIntentEndings — end to end with the word table of testdata, the endings of intent and guess and the other linking and final endings added to the end of closings.tsv: 려나, 려니, 려다가, 려거든 and 랴 (with 으 after a consonant), 으되, 건마는, ㄹ진대 and 을진대, ㄹ는지 and 을는지, ㄴ대야 and 는대야, ㄴ담, ㄴ고 after the copula's stem, 다네, the modifying 으라는; 려다가, 려거든, 는대야 and ㄴ담 after a verb's stem and not after an adjective's
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamIntentEndings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"가려나":   "\n가[열쇠:동사]+려나[맺음:려나]\n",
		"먹으려나":  "\n먹[열쇠:동사]+으려나[맺음:으려나]\n",
		"가려니":   "\n가[열쇠:동사]+려니[맺음:려니]\n",
		"먹으려니":  "\n먹[열쇠:동사]+으려니[맺음:으려니]\n",
		"가려다가":  "\n가[열쇠:동사]+려다가[맺음:려다가]\n",
		"먹으려다가": "\n먹[열쇠:동사]+으려다가[맺음:으려다가]\n",
		"가려거든":  "\n가[열쇠:동사]+려거든[맺음:려거든]\n",
		"먹으려거든": "\n먹[열쇠:동사]+으려거든[맺음:으려거든]\n",
		"가랴":    "\n가[열쇠:동사]+랴[맺음:랴]\n",
		"먹으랴":   "\n먹[열쇠:동사]+으랴[맺음:으랴]\n",
		"먹으되":   "\n먹[열쇠:동사]+으되[맺음:으되]\n",
		"가건마는":  "\n가[열쇠:동사]+건마는[맺음:건마는]\n",
		"갈진대":   "\n가[열쇠:동사]+(ㄹ)진대[맺음:ㄹ진대]\n",
		"먹을진대":  "\n먹[열쇠:동사]+을진대[맺음:을진대]\n",
		"갈는지":   "\n가[열쇠:동사]+(ㄹ)는지[맺음:ㄹ는지]\n",
		"먹을는지":  "\n먹[열쇠:동사]+을는지[맺음:을는지]\n",
		"간대야":   "\n가[열쇠:동사]+(ㄴ)대야[맺음:ㄴ대야]\n",
		"먹는대야":  "\n먹[열쇠:동사]+는대야[맺음:는대야]\n",
		"간담":    "\n가[열쇠:동사]+(ㄴ)담[맺음:ㄴ담]\n",
		"학교인고":  "\n학교[열쇠:체언]+이[서술격:이]+(ㄴ)고[맺음:ㄴ고]\n",
		"작다네":   "\n작[열쇠:형용사]+다네[맺음:다네]\n",
		"먹으라는":  "\n먹[열쇠:동사]+으라는[꾸밈:으라는]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작으려다가": "작[열쇠:형용사]+으려다가",
		"작으려거든": "작[열쇠:형용사]+으려거든",
		"작는대야":  "작[열쇠:형용사]+는대야",
		"예쁜담":   "예쁘[열쇠:형용사]+(ㄴ)담",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
