//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddNumbersStopped — two numbers joined by a full stop with a full stop right after them and letters after that (1.28.은, 8.3.부터, 1961.12.이전) are read first as a date, the decimal of the same letters right behind it; with no full stop after them (1.28에서) or with the full stop at the unit's end (2.5.) the decimal stays first and no date is read
package token

import (
	"strings"
	"testing"
)

func TestAddNumbersStopped(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, c := range []struct{ text, nums string }{{"1.28.은", "1.28"}, {"8.3.부터", "8.3"}, {"1961.12.이전", "1961.12"}} {
		seg, _ := r.Stream(c.text).Next()
		got := seg.Nodes[0].Contest.Branches
		if !strings.HasPrefix(got[0].String(), c.nums+"[날짜]+.[마침표]") || !strings.HasPrefix(got[1].String(), c.nums+"[소수점]+.[마침표]") {
			t.Errorf("%s: date first, decimal right behind: %q", c.text, got)
		}
	}
	for _, text := range []string{"1.28에서", "2.5."} {
		if got := streamText(r, text); strings.Contains(got, "[날짜]") || !strings.Contains(got, "[소수점]") {
			t.Errorf("%s: no date\n%s", text, got)
		}
	}
}
