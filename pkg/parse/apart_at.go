//ff:func feature=parse type=helper control=sequence
//ff:what apartAt — does a space, or the edge of the letters, stand on both sides of place i
package parse

import "unicode"

func apartAt(r []rune, i int) bool {
	left, right := atOf(r, i-1), atOf(r, i+1)
	return (left == 0 || unicode.IsSpace(left)) && (right == 0 || unicode.IsSpace(right))
}
