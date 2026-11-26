//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDropAfterI — end to end with the word table of testdata, after a nominal of the table that ends in the syllable 이 the copula's stem is read as dropped (고양이다, 고양이면, 고양이라는), beside the reading with the syllable as the stem said; not after a word no table has (뚜루이다), not when 이 begins the unit (이다), and not without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDropAfterI(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, _ := New(Options{})
	for in, want := range map[string]string{
		"고양이다":  "고양이[열쇠:체언]+(이)[서술격:이(빠짐)]+다[맺음:다]",
		"고양이면":  "고양이[열쇠:체언]+(이)[서술격:이(빠짐)]+면[맺음:면]",
		"고양이라는": "고양이[열쇠:체언]+(이)[서술격:이(빠짐)]+라는[꾸밈:라는]",
	} {
		got := streamText(r, in)
		if !strings.Contains(got, "\n"+want+"\n") || !strings.Contains(got, "+이[서술격:이]+") {
			t.Errorf("%s: no %q beside the stem said\n%s", in, want, got)
		}
		if got = streamText(bare, in); strings.Contains(got, "이+(이)[서술격:이(빠짐)]") {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
	for _, in := range []string{"뚜루이다", "이다"} {
		if got := streamText(r, in); strings.Contains(got, "이+(이)[서술격:이(빠짐)]") || strings.Contains(got, "이[열쇠:체언]+(이)[서술격:이(빠짐)]") {
			t.Errorf("%s: a stem dropped after 이\n%s", in, got)
		}
	}
}
