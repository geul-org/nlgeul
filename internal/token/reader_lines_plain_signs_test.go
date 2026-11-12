//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesPlainSigns — with no word table, signs no reading stands for come up as their letters with no name: x < y and A -> B have < and -> as lines of their own, 5'10" is the word 5'10 and a ", 가,나 keeps the comma-closed reading 가+,[쉼표] 나 beside the whole word and the reading with 나 as a particle
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesPlainSigns(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"x < y", []string{"x[열:영숫자]", "<", "y[열:영숫자]"}},
		{"A -> B", []string{"A[열:영숫자]", "->", "B[열:영숫자]"}},
		{"5'10\"", []string{"5'10", "\""}},
		{"가,나", []string{"가,나\t경합(", "  가,+나[보조사:나]", "  가,나\t", "  가+,[쉼표] 나+(아)[맺음:어(합침)]"}},
	} {
		lines := r.Lines(c.in, false)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
