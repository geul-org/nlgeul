//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestGatedClosingLse — end to end with a word table, the closing ㄹ세 of 하게체 (gated-closings.tsv) is read fused under the copula's stem 이 (학교일세) and under a predicate's stem of the table that ends in a vowel (아닐세), and not under the letters of a stem the table does not have (월세 has no closing ㄹ세)
package token

import (
	"strings"
	"testing"
)

func TestGatedClosingLse(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-hao.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ in, want string }{
		{"학교일세", "학교[열쇠:체언]+이[서술격:이]+(ㄹ)세[맺음:ㄹ세]"},
		{"아닐세", "아니[열쇠:형용사]+(ㄹ)세[맺음:ㄹ세]"},
	} {
		lines := streamLines(r, c.in)
		if len(lines) == 0 || lines[0] != c.want {
			t.Errorf("%s: want first %q\n%s", c.in, c.want, strings.Join(lines, "\n"))
		}
	}
	if got := streamText(r, "월세"); strings.Contains(got, "맺음:ㄹ세]") {
		t.Errorf("월세: a closing after a stem not in the table\n%s", got)
	}
}
