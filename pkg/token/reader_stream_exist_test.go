//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamExist — end to end with the word table of testdata, whose 작 is an adjective: 느냐 does not stand after an adjective's stem (작느냐), but does after the slot 었 or 겠 (작았느냐, 작겠느냐) and after a verb's stem (먹느냐); a closing only verbs take still does not stand after the slot (작았는다 has no 작+았+는다)
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamExist(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"작았느냐": "\n작[열쇠:형용사]+았[선어말:았]+느냐[맺음:느냐]\n",
		"작겠느냐": "\n작[열쇠:형용사]+겠[선어말:겠]+느냐[맺음:느냐]\n",
		"먹느냐":  "\n먹[열쇠:동사]+느냐[맺음:느냐]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작느냐":  "작[열쇠:형용사]+느냐",
		"작았는다": "작[열쇠:형용사]+았[선어말:았]+는다[맺음:는다]",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
