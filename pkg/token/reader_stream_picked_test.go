//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamPicked — end to end with the word table of testdata, 를 written onto the linking closing 지 is a piece of the same word (먹+지+를, 가+지+를 beside the noun 가지 of the table), and 가 after the same closing is not read onto it
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamPicked(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"먹지를 않는다": "\n먹[열쇠:동사]+지[맺음:지]+를[목적격:를]\n",
		"가지를":     "\n가[열쇠:동사]+지[맺음:지]+를[목적격:를]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "먹지가 않다"); strings.Contains(got, "지[맺음:지]+가") {
		t.Fatal("가 after 지\n" + got)
	}
}
