//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamLinkMore — end to end with the word table of testdata, the auxiliary particles 까지, 라도 and 은 written onto a linking closing are pieces of the same word (먹+게+까지, 하+면서+라도, 먹+으면+은, 가+면+은), and so is one after 면 (가+면+도); a case particle after 면 is not read onto it (가면을)
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamLinkMore(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹게까지":  "\n먹[열쇠:동사]+게[맺음:게]+까지[보조사:까지]\n",
		"하면서라도": "\n하[열쇠:동사]+면서[맺음:면서]+라도[보조사:라도]\n",
		"먹으면은":  "\n먹[열쇠:동사]+으면[맺음:으면]+은[보조사:은]\n",
		"가면은":   "\n가[열쇠:동사]+면[맺음:면]+은[보조사:은]\n",
		"가면도":   "\n가[열쇠:동사]+면[맺음:면]+도[보조사:도]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "가면을"); strings.Contains(got, "면[맺음:면]+을") {
		t.Fatal("a case particle after 면\n" + got)
	}
}
