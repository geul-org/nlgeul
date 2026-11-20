//ff:func feature=cli type=helper control=selection
//ff:what parseLines — the lines nlgeul parse prints for one sentence: the word ends with cuts; the word ends with the spacing rules of each boundary with marks; else the boxes of the stream (a box for each word, the ways of a stretch at the right in brackets)
package cli

import "github.com/geul-org/nlgeul/pkg/parse"

func parseLines(r *parse.Reader, s string, trace, cuts, marks bool, head int) []string {
	switch {
	case cuts:
		return r.Cuts(s)
	case marks:
		return r.Marks(s)
	}
	return r.BoxLines(s, trace, head)
}
