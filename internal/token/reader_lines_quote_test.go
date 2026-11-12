//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesQuote — with no word table, a single quotation mark is read in every way it may stand: 'I'm (O'Brien)' opens with 열기:1 or the letter, the apostrophes inside words are no signs, the bracket is 열기:1 or 열기:2 and closes as 닫기:1 or 닫기:2, and the last ' is 열기:1, 닫기:1 or the letter; it's is one word; '90년대 gives the quote as 열기:1 or its letter; a quote closed in a later word with a particle on it is 닫기:1 beside 열기:1 and the letter
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesQuote(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"'I'm (O'Brien)'", []string{"'\t경합(열기:1 / ')", "  '[열기:1]\t성립 1.00", "I'm", "(\t경합(열기:1 / 열기:2)", "O'Brien", ")\t경합(닫기:1 / 닫기:2)", "'\t경합(열기:1 / 닫기:1)", "  '[열기:1]\t성립 1.00", "  '[닫기:1]\t성립 1.00", "  '\t성립 1.00"}},
		{"it's", []string{"it's"}},
		{"'90년대", []string{"'\t경합(열기:1 / ')", "90년대"}},
		{"곰이 '사자 호랑이'는 싫다", []string{"곰이", "'\t경합(열기:1 / ')", "사자", "호랑이", "'는\t경합(열기:1 + 보조사:는 / 닫기:1 + 보조사:는)", "  '[열기:1]+는[보조사:는]", "  '[닫기:1]+는[보조사:는]", "  '+는[보조사:는]", "싫다"}},
	} {
		lines := r.Lines(c.in, false)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
