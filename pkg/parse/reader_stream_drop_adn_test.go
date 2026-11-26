//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDropAdn — end to end with the word table of testdata, the copula's stem is read as dropped before the modifying ㄴ put under the last syllable of a nominal of the table (문젠, 세댄), not under a word no table has (고랜), and not without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDropAdn(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, _ := New(Options{})
	const dropped = "(이)[서술격:이(빠짐)]+(ㄴ)[꾸밈:ㄴ]"
	for in, noun := range map[string]string{"문젠": "문제", "세댄": "세대"} {
		if got := streamText(r, in); !strings.Contains(got, "\n"+noun+"[열쇠:체언]+"+dropped+"\n") {
			t.Errorf("%s: no dropped stem before ㄴ\n%s", in, got)
		}
		if got := streamText(bare, in); strings.Contains(got, dropped) {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
	if got := streamText(r, "고랜"); strings.Contains(got, dropped) {
		t.Errorf("고랜: a dropped stem under a word no table has\n%s", got)
	}
}
