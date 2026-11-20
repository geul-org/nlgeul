//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamCorner — the halfwidth corner brackets ｢ ｣ of the table of pairs are signs of a pair of their own kind: ｢ opens, ｣ closes it with a particle written onto it read as one, and ｣ written onto a word with none open closes at depth 1 as the other closing brackets do
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamCorner(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"｢가나｣를 본다", []string{"｢[열기:1]\t성립 1.00", "가나", "｣[닫기:1]\t성립 1.00", "+를[목적격:를]", "본다"}},
		{"「가 ｢나｣」", []string{"「[열기:1]", "가", "｢[열기:2]", "나", "｣[닫기:2]", "」[닫기:1]"}},
		{"옛말｣을 본다", []string{"옛말", "｣[닫기:1]\t성립 1.00", "+을[목적격:을]", "본다"}},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
