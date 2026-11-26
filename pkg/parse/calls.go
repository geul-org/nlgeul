//ff:func feature=parse type=reader control=sequence
//ff:what calls — do the signs after a unit's letters begin with one written right after one who is called: a comma (also the fullwidth one) or an exclamation mark (영숙이, 영숙이!)
package parse

import "strings"

func calls(signs string) bool {
	return strings.HasPrefix(signs, ",") || strings.HasPrefix(signs, "，") || strings.HasPrefix(signs, "!")
}
