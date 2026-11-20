//ff:func feature=cli type=command control=iteration dimension=2
//ff:what runParse — read each sentence with the reader and print its boxes (a box cut to its first head candidates and a stretch cut to its first head ways, 0 for all), or its word ends when those are asked for (cuts), or the word ends with the spacing rules (marks); several sentences are parted by an empty line
package cli

import (
	"fmt"
	"io"

	"github.com/geul-org/nlgeul/pkg/token"
)

func runParse(out io.Writer, r *token.Reader, sentences []string, trace, cuts, marks bool, head int) {
	for _, s := range sentences {
		for _, l := range parseLines(r, s, trace, cuts, marks, head) {
			fmt.Fprintln(out, l)
		}
		if len(sentences) > 1 {
			fmt.Fprintln(out)
		}
	}
}
