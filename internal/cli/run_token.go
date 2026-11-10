//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runToken — print the lines of token.Lines for the sentence, or of token.Cuts when the word ends are asked for
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/koparse/internal/token"
)

func runToken(out io.Writer, sentence string, trace, cuts bool) error {
	ls := token.Lines(sentence, trace)
	if cuts {
		ls = token.Cuts(sentence)
	}
	for _, l := range ls {
		fmt.Fprintln(out, l)
	}
	return nil
}
