//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunHypoPage — page 0 of runHypo prints every hypothesis however many pages they fill; a page of runHypo: the first line is still the count of all the hypotheses, then only the lines of that page, each ranked as in the whole list and written as hypo.Line writes it; the last page holds what is left; a page past the last one prints the first line only; a sentence past hypo.Limit is only counted and refused, a page asked for or not
package cli

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/internal/hypo"
	"github.com/geul-org/nlgeul/internal/token"
)

func TestRunHypoPage(t *testing.T) {
	bare, err := token.New(token.Options{})
	if err != nil {
		t.Fatal(err)
	}
	sentence := strings.TrimSpace(strings.Repeat("작은 ", 5))
	choices := bare.Bundles(sentence)
	count := hypo.Count(choices)
	total := int(count.Int64())
	if total <= 2*hypo.PerPage || total >= 3*hypo.PerPage {
		t.Fatalf("%d hypotheses, want more than two pages of hypo.PerPage and less than three", total)
	}
	// lines[0] is the count, lines[r] the line of rank r in the whole list.
	lines := []string{"가설 " + count.String()}
	for i, h := range hypo.List(choices) {
		lines = append(lines, hypo.Line(i+1, h, choices))
	}
	cases := []struct {
		page, perPage int
		from, to      int // the ranks printed; none when from is past to
	}{
		{0, hypo.PerPage, 1, total},
		{1, hypo.PerPage, 1, hypo.PerPage},
		{2, hypo.PerPage, hypo.PerPage + 1, 2 * hypo.PerPage},
		{3, hypo.PerPage, 2*hypo.PerPage + 1, total},
		{4, hypo.PerPage, total + 1, total},
		{1, 1, 1, 1},
		{3, 2, 5, 6},
		{total, 1, total, total},
		{total + 1, 1, total + 1, total},
		{1, total, 1, total},
		{2, total, total + 1, total},
	}
	var out bytes.Buffer
	for _, c := range cases {
		out.Reset()
		want := strings.Join(append(lines[:1:1], lines[c.from:c.to+1]...), "\n") + "\n"
		if err := runHypo(&out, bare, sentence, c.page, c.perPage); err != nil || out.String() != want {
			t.Errorf("page %d of %d: %v\n%q\nwant the count and ranks %d to %d", c.page, c.perPage, err, out.String(), c.from, c.to)
		}
	}

	many := strings.TrimSpace(strings.Repeat("작은 ", 13))
	over := hypo.Count(bare.Bundles(many))
	out.Reset()
	if over.Cmp(big.NewInt(hypo.Limit)) <= 0 || runHypo(&out, bare, many, 1, hypo.PerPage) == nil || out.String() != "가설 "+over.String()+"\n" {
		t.Fatalf("past the limit with page 1: %s hypotheses, %q", over, out.String())
	}
}
