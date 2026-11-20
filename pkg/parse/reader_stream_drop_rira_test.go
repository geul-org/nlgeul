//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDropRira — end to end with the word table of testdata, the copula's stem 이 is read as dropped before 리라 after a noun that ends in a vowel (세대리라, 사과리라), and not after one that ends in a consonant (입력리라 gives no 입력+(이)+리라)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDropRira(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, in := range []string{"세대리라", "사과리라"} {
		want := "\n" + strings.TrimSuffix(in, "리라") + "[열쇠:체언]+(이)[서술격:이(빠짐)]+리라[맺음:리라]\n"
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "입력리라"); strings.Contains(got, "입력[열쇠:체언]+(이)[서술격:이(빠짐)]+리라") {
		t.Errorf("이 dropped after a consonant\n%s", got)
	}
}
