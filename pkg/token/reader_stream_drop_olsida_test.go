//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamDropOlsida — end to end with the word table of testdata, the copula's stem 이 is read as dropped before the closing 올시다 after a noun that ends in a vowel (세대올시다, 사과올시다), and not after one that ends in a consonant (입력올시다 gives no 입력+(이)+올시다), where it stands written (입력이올시다)
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamDropOlsida(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, in := range []string{"세대올시다", "사과올시다"} {
		want := "\n" + strings.TrimSuffix(in, "올시다") + "[열쇠:체언]+(이)[서술격:이(빠짐)]+올시다[맺음:올시다]\n"
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
	}
	if got := streamText(r, "입력올시다"); strings.Contains(got, "입력[열쇠:체언]+(이)[서술격:이(빠짐)]+올시다") {
		t.Errorf("이 dropped after a consonant\n%s", got)
	}
	if got := streamText(r, "입력이올시다"); !strings.Contains(got, "\n입력[열쇠:체언]+이[서술격:이]+올시다[맺음:올시다]\n") {
		t.Errorf("이 written\n%s", got)
	}
}
