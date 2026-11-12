//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runToken — read each sentence with the reader and print its lines (a contest cut to its first head readings, 0 for all), or its word ends when those are asked for; several sentences are parted by an empty line
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/nlgeul/internal/token"
)

func runToken(out io.Writer, r *token.Reader, sentences []string, trace, cuts bool, head int) {
	for _, s := range sentences {
		ls := r.LinesHead(s, trace, head)
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
