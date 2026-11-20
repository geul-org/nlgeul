//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLinkAux — end to end with the word table of testdata, an auxiliary particle written onto the linking closings 다가, 고서 and 려고 (으려고) is a piece of the same word (먹+었+다가+도, 하+고서+야, 가+려고+는, 먹+으려고+도); a case particle is not (먹다가를), and after 는데 the particle is not read onto the closing (먹는데도 has no 는데+도)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamLinkAux(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹었다가도": "\n먹[열쇠:동사]+었[선어말:었]+다가[맺음:다가]+도[보조사:도]\n",
		"하고서야":  "\n하[열쇠:동사]+고서[맺음:고서]+야[보조사:야]\n",
		"가려고는":  "\n가[열쇠:동사]+려고[맺음:려고]+는[보조사:는]\n",
		"먹으려고도": "\n먹[열쇠:동사]+으려고[맺음:으려고]+도[보조사:도]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"먹다가를": "다가[맺음:다가]+를",
		"먹는데도": "는데[맺음:는데]+도",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
