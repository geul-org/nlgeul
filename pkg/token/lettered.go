//ff:func feature=token type=helper control=sequence
//ff:what lettered — does a Latin letter or a digit stand on both sides of a place (an apostrophe there is part of a word, not a quotation mark: it's, '90s); a Hangul letter on either side does not count, so that a closing mark with a particle written onto it pairs (… 호랑이'는)
package token

import "unicode"

func lettered(r []rune, i int) bool {
	both := func(c rune) bool { return unicode.Is(unicode.Latin, c) || unicode.IsDigit(c) }
	return i > 0 && i+1 < len(r) && both(r[i-1]) && both(r[i+1])
}
