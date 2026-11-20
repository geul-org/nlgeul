//ff:func feature=parse type=helper control=sequence
//ff:what scoredText — the lines of scoredLines joined with newlines, with a newline before the first and after the last, so a whole candidate is found as "\n"+line+"\t" and with its verdict as "\n"+line+"\t성립 0.33"
package parse

import "strings"

func scoredText(r *Reader, sentence string) string {
	return "\n" + strings.Join(scoredLines(r, sentence), "\n") + "\n"
}
