//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestTokenLines — cuts give the word ends, marks the word ends with the spacing rules of each boundary, and the default gives BoxLines as it is (with trace and head); none of the three is the same as another
package cli

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/token"
)

func TestTokenLines(t *testing.T) {
	r, err := token.New(token.Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	s := "고양이가 마당에서 물을"
	cuts := tokenLines(r, s, false, true, false, 0)
	if !slices.Equal(cuts, r.Cuts(s)) {
		t.Fatal(cuts)
	}
	for _, c := range []struct {
		trace bool
		head  int
	}{{false, 0}, {false, 1}, {false, 20}, {true, 0}, {true, 1}, {true, 20}} {
		box := tokenLines(r, s, c.trace, false, false, c.head)
		if !slices.Equal(box, r.BoxLines(s, c.trace, c.head)) || slices.Equal(box, cuts) {
			t.Fatal(c, box)
		}
	}
	marks := tokenLines(r, s, false, false, true, 0)
	if !slices.Equal(marks, r.Marks(s)) || slices.Equal(marks, cuts) || slices.Equal(marks, tokenLines(r, s, false, false, false, 0)) {
		t.Fatal("marks", marks)
	}
	if got := tokenLines(r, s, false, true, true, 0); !slices.Equal(got, cuts) {
		t.Fatal("cuts come first", got)
	}
}
