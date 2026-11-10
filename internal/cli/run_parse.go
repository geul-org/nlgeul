//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runParse — print the lines of parse.Lines for the sentence
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/koparse/internal/parse"
)

// lines judges a sentence; tests replace it to observe a failure.
var lines = parse.Lines

func runParse(out io.Writer, sentence string) error {
	ls, err := lines(sentence)
	if err != nil {
		return err
	}
	for _, l := range ls {
		fmt.Fprintln(out, l)
	}
	return nil
}
