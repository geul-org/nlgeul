//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPairsOrder — the candidates of a sign of a pair come in the order of the ways the signs before it were read, first seen first: '가'는 closes first at the second quote, '가' '나' opens first at the third and closes first at the fourth, two straight double quotes close at the fourth, a bracket inside an open quote is depth 2 first and depth 1 after it where with no quote it is depth 1 only, and it's is no sign
package parse

import (
	"slices"
	"testing"
)

func TestReaderStreamPairsOrder(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, sign string
		at       int
		want     []string
	}{
		{"'가'는", "'", 1, []string{"닫기:1", "열기:1", ""}},
		{"'가' '나'", "'", 2, []string{"열기:1", "닫기:1", ""}},
		{"'가' '나'", "'", 3, []string{"닫기:1", "열기:1", ""}},
		{"\"가\"다 \"나\"다", "\"", 3, []string{"닫기:1"}},
		{"'가 (나)'", "(", 0, []string{"열기:2", "열기:1"}},
		{"'가 (나)'", ")", 0, []string{"닫기:2", "닫기:1"}},
		{"(나)", "(", 0, []string{"열기:1"}},
		{"(나)", ")", 0, []string{"닫기:1"}},
	} {
		signs := pairsSigns(r, c.in, c.sign)
		if len(signs) <= c.at {
			t.Errorf("%q: %d of %q", c.in, len(signs), c.sign)
			continue
		}
		got := kindsOf(signs[c.at])
		if !slices.Equal(got, c.want) {
			t.Errorf("%q: %q number %d: %q, want %q", c.in, c.sign, c.at, got, c.want)
		}
	}
	if got := streamContests(r, "it's"); len(got) != 1 || got[0].Text != "it's" {
		t.Errorf("it's is one word: %v", got)
	}
}
