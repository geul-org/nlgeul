//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestMeasuredSentence — end to end: a quotation mark right after a digit that closes a quote open before it is read first as the closing ('A12'에서, "높이 4"), the units staying candidates behind it; with no quote open before it (5'10") the last quote is read first as a unit (inches), as before
package token

import (
	"strings"
	"testing"
)

func TestMeasuredSentence(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, c := range []struct{ in, want string }{
		{"'A12'에서", "'[닫기:1]"},
		{"\"높이 4\"", "\"[닫기:1]"},
		{"키는 5'10\"", "\"[단위:인치]"},
	} {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		if got := " │ " + strings.Join(words, " │ ") + " │ "; !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
	}
	if !netHas(r, "'A12'에서", "'[단위:피트]") {
		t.Fatal("the unit is gone")
	}
}
