//ff:type feature=token type=model
//ff:what bracket — the two brackets drawn at the right of the boxes of a contest of ways, kept for each line: the letter of the inner bracket (one way), the letter of the outer bracket (the whole contest), and the tag after them (empty for none)
package token

type bracket struct {
	inner []rune
	outer []rune
	tags  []string
}
