//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what wholeFirst — of the lines streamLines gives for a sentence of one word, the first that reads it as one node (no " " between nodes): the first candidate of the word read whole, whichever way comes first; "" when no way reads it whole
package parse

import "strings"

func wholeFirst(lines []string) string {
	for _, l := range lines {
		if !strings.Contains(l, " ") {
			return l
		}
	}
	return ""
}
