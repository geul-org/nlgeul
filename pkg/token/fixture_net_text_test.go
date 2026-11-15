//ff:func feature=token type=helper control=sequence
//ff:what netText — the lines of netLines joined with newlines, with a newline before the first and after the last, so a whole line is found as "\n"+line+"\n"
package token

import "strings"

func netText(r *Reader, sentence string) string {
	return "\n" + strings.Join(netLines(r, sentence), "\n") + "\n"
}
