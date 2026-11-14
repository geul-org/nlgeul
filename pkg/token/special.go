//ff:func feature=token type=helper control=selection
//ff:what special — what a unit is when it is no sentence's word, and where that ends: an address (http://, https://), a mail address, a path (a path's head, or two slashes and more between names), each up to the last letter that is no Hangul syllable (a particle may follow); an option (- or -- and Latin letters); "" and 0 for anything else
package token

import "strings"

func special(s []rune) (string, int) {
	text, end := string(s), lastPlain(s)
	switch {
	case end == 0:
		return "", 0
	case strings.HasPrefix(text, "http://") || strings.HasPrefix(text, "https://"):
		return "주소", end
	case isMail(s[:end]):
		return "전자우편", end
	case isPath(s[:end]):
		return "경로", end
	case optionEnd(s) > 0:
		return "옵션", optionEnd(s)
	}
	return "", 0
}
