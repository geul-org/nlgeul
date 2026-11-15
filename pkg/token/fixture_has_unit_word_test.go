//ff:func feature=token type=helper control=sequence
//ff:what hasUnitWord — whether some line the stream of a sentence gives (streamLines) prints exactly as the given text
package token

import "slices"

func hasUnitWord(r *Reader, s, text string) bool {
	return slices.Contains(streamLines(r, s), text)
}
