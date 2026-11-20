//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamLean — a sentence that begins with a quotation mark left over from the sentence before: the straight quotation mark written onto the word after it opens first though one may stand open, the one written onto the word before it closes first, and the bracket after them is read on from those (depth 1 first, then 2)
package parse

import (
	"slices"
	"testing"
)

func TestReaderStreamLean(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	const in = "', '봄날 간다' 하고 (집에) 갔다"
	for _, c := range []struct {
		sign string
		at   int
		want []string
	}{
		{"'", 0, []string{"열기:1", ""}},
		{"'", 1, []string{"열기:1", "닫기:1", ""}},
		{"'", 2, []string{"닫기:1", "열기:1", ""}},
		{"(", 0, []string{"열기:1", "열기:2"}},
		{")", 0, []string{"닫기:1", "닫기:2"}},
	} {
		signs := pairsSigns(r, in, c.sign)
		if len(signs) <= c.at {
			t.Fatalf("%d of %q", len(signs), c.sign)
		}
		if got := kindsOf(signs[c.at]); !slices.Equal(got, c.want) {
			t.Errorf("%q number %d: %q, want %q", c.sign, c.at, got, c.want)
		}
	}
}
