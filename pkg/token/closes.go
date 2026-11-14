//ff:func feature=token type=reader control=sequence
//ff:what closes — do the signs after a unit's letters begin with one that ends a sentence: a question mark, an exclamation mark, or a full stop that is not the first of dots in a row (those trail off: 글쎄...)
package token

import "strings"

func closes(signs string) bool {
	return strings.HasPrefix(signs, "?") || strings.HasPrefix(signs, "!") || (strings.HasPrefix(signs, ".") && !strings.HasPrefix(signs, ".."))
}
