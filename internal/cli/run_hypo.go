//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runHypo — read a sentence with the token reader, print how many hypotheses its bundles give, then every hypothesis on a line, best first, or only those of one page when a page is asked for; a sentence with more than hypo.Limit is only counted and refused
package cli

import (
	"fmt"
	"io"
	"math/big"

	"github.com/geul-org/nlgeul/internal/hypo"
	"github.com/geul-org/nlgeul/internal/token"
)

// Page 0 prints every hypothesis. A page keeps the ranks of the whole
// list: the first line of page 2 of perPage 100 is ranked 101.
func runHypo(out io.Writer, r *token.Reader, sentence string, page, perPage int) error {
	choices := r.Bundles(sentence)
	count := hypo.Count(choices)
	fmt.Fprintf(out, "가설 %s\n", count)
	if count.Cmp(big.NewInt(hypo.Limit)) > 0 {
		return fmt.Errorf("가설이 %d개를 넘어 나열하지 않는다", hypo.Limit)
	}
	first, listed := hypo.Page(hypo.List(choices), page, perPage)
	for i, h := range listed {
		fmt.Fprintln(out, hypo.Line(first+i+1, h, choices))
	}
	return nil
}
