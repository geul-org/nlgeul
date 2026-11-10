//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runToken — read each sentence with the reader and print its lines, or its word ends when those are asked for; several sentences are parted by an empty line
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/koparse/internal/token"
)

func runToken(out io.Writer, r *token.Reader, sentences []string, trace, cuts bool) {
	for _, s := range sentences {
		ls := r.Lines(s, trace)
		if cuts {
			ls = r.Cuts(s)
		}
		for _, l := range ls {
			fmt.Fprintln(out, l)
		}
		if len(sentences) > 1 {
			fmt.Fprintln(out)
		}
	}
}
