//ff:func feature=parse type=helper control=sequence
//ff:what unitHas — whether some line the stream of a sentence gives (streamLines) has the given text in it
package parse

import "strings"

func unitHas(r *Reader, s, text string) bool {
	return strings.Contains(streamText(r, s), text)
}
