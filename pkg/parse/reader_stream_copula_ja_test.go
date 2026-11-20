//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamCopulaJa — end to end with the word table of testdata: the linking closing 자 stands after the copula's stem (학교+이+자, the dictionary's 자 that joins two standings of one thing) and after a verb's stem (먹+자), but not after an adjective's stem (작+자 is no stem of the table), and 자마자 still does not stand after the copula's stem
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamCopulaJa(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"학교이자": "\n학교[열쇠:체언]+이[서술격:이]+자[맺음:자]\n",
		"먹자":   "\n먹[열쇠:동사]+자[맺음:자]\n",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작자":     "작[열쇠:형용사]+자[맺음:자]",
		"학교이자마자": "[서술격:이]+자마자[맺음:자마자]",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %q read\n%s", in, not, got)
		}
	}
}
