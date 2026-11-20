//ff:func feature=parse type=helper control=sequence
//ff:what streamText — the lines of streamLines joined with newlines, with a newline before the first and after the last, so a whole line is found as "\n"+line+"\n"
package parse

import "strings"

func streamText(r *Reader, sentence string) string {
	return "\n" + strings.Join(streamLines(r, sentence), "\n") + "\n"
}
