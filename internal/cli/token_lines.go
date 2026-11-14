//ff:func feature=cli type=helper control=selection
//ff:what tokenLines — the lines nlgeul token prints for one sentence: the word ends with cuts; the word ends with the spacing rules of each boundary with marks; the old one line a candidate form with plain; else the boxes of the stream (a box for each word, the ways of a stretch at the right in brackets)
package cli

import "github.com/geul-org/nlgeul/internal/token"

func tokenLines(r *token.Reader, s string, trace, cuts, marks, plain bool, head int) []string {
	switch {
	case cuts:
		return r.Cuts(s)
	case marks:
		return r.Marks(s)
	case plain:
		return r.LinesHead(s, trace, head)
	}
	return r.BoxLines(s, trace, head)
}
