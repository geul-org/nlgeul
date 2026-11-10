//ff:func feature=token type=helper control=sequence
//ff:what lettered — does a letter or digit stand on both sides of a place (an apostrophe there is part of a word, not a quotation mark)
package token

import "unicode"

func lettered(r []rune, i int) bool {
	both := func(c rune) bool { return unicode.IsLetter(c) || unicode.IsDigit(c) }
	return i > 0 && i+1 < len(r) && both(r[i-1]) && both(r[i+1])
}
