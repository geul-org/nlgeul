//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDropRago — end to end with the word table of testdata, the copula's stem 이 is read as dropped before the quoting closing 라고 after a nominal of the table that ends in a vowel (사과라고, 세대라고), beside the particle reading 라고 that stays; not after a word no table has (고래라고), not after a consonant (입력라고), and not without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDropRago(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, _ := New(Options{})
	const dropped = "(이)[서술격:이(빠짐)]+라고[맺음:라고]"
	for _, in := range []string{"사과라고", "세대라고"} {
		noun := strings.TrimSuffix(in, "라고") + "[열쇠:체언]+"
		got := streamText(r, in)
		if !strings.Contains(got, "\n"+noun+dropped+"\n") || !strings.Contains(got, "\n"+noun+"라고[부사격:라고]\n") {
			t.Errorf("%s: the dropped stem before 라고 and the particle 라고\n%s", in, got)
		}
		if got = streamText(bare, in); strings.Contains(got, dropped) {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
	for _, in := range []string{"고래라고", "입력라고"} {
		if got := streamText(r, in); strings.Contains(got, dropped) {
			t.Errorf("%s: a dropped stem before 라고\n%s", in, got)
		}
	}
}
