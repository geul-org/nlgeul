//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runMini — print the lines of mini.Lines for the sentence
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/koparse/internal/mini"
)

// lines judges a sentence; tests replace it to observe a failure.
var lines = mini.Lines

func runMini(out io.Writer, sentence string) error {
	ls, err := lines(sentence)
	if err != nil {
		return err
	}
	for _, l := range ls {
		fmt.Fprintln(out, l)
	}
	return nil
}
