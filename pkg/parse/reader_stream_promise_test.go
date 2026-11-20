//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPromise — end to end, the closing 마 of a promise is read after a verb's stem of the word table (가마, 먹으마 with 으 after a consonant, 살마 after a stem's ㄹ) and not after an adjective's stem (작으마), not after a stem that is not in the table (고양이가마당 gives no 고양이가+마), and not without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamPromise(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"가마":  "\n가[열쇠:동사]+마[맺음:마]\n",
		"먹으마": "\n먹[열쇠:동사]+으마[맺음:으마]\n",
		"살마":  "\n살[열쇠:동사]+마[맺음:마]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작으마":    "작[열쇠:형용사]+으마",
		"고양이가마당": "고양이가+마",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	if got := streamText(bare, "가마"); strings.Contains(got, "[맺음:마]") {
		t.Errorf("마 read without a word table\n%s", got)
	}
}
