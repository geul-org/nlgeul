//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestTokenLines — cuts give the word ends (whatever plain is), marks the word ends with the spacing rules of each boundary, plain gives LinesHead as it is (with trace and head), and the default gives BoxLines; none of the three is the same as another
package cli

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestTokenLines(t *testing.T) {
	r, err := token.New(token.Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	s := "고양이가 마당에서 물을"
	cuts := tokenLines(r, s, false, true, false, true, 0)
	if !slices.Equal(cuts, r.Cuts(s)) {
		t.Fatal(cuts)
	}
	for _, c := range []struct {
		trace bool
		head  int
	}{{false, 0}, {false, 1}, {false, 20}, {true, 0}, {true, 1}, {true, 20}} {
		plain := tokenLines(r, s, c.trace, false, false, true, c.head)
		if !slices.Equal(plain, r.LinesHead(s, c.trace, c.head)) {
			t.Fatal(c, plain)
		}
		box := tokenLines(r, s, c.trace, false, false, false, c.head)
		if !slices.Equal(box, r.BoxLines(s, c.trace, c.head)) || slices.Equal(box, plain) {
			t.Fatal(c, box)
		}
	}
	if slices.Equal(cuts, tokenLines(r, s, false, false, false, false, 0)) {
		t.Fatal("cuts is not the boxes")
	}
	marks := tokenLines(r, s, false, false, true, true, 0)
	if !slices.Equal(marks, r.Marks(s)) || slices.Equal(marks, cuts) {
		t.Fatal("marks", marks)
	}
	if got := tokenLines(r, s, false, true, true, true, 0); !slices.Equal(got, cuts) {
		t.Fatal("cuts come first", got)
	}
}
