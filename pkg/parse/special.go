//ff:func feature=parse type=helper control=selection
//ff:what special — what a unit is when it is no sentence's word, and where that ends: an address (http://, https://), a mail address, a path (a path's head, or two slashes and more between names), each up to the last letter that is no Hangul syllable (a particle may follow), a path past the Hangul syllables of the last part of its name when a sign joins them on (joinedEnd: a/b-가격정책); an option (- or -- and Latin letters); "" and 0 for anything else
package parse

import "strings"

// pathName is what a path is called: one that special knows, and two
// runs that a slash joins into a name (nameSigns).
const pathName = "경로"

// pathTailName and pathWeakName name the graphs of the two readings of
// a path that the letters alone do not settle (addPaths): both are
// printed as a path (pathName).
const (
	pathTailName = "경로:빗금 뒤 한글"
	pathWeakName = "경로:머리 없음"
)

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
		return pathName, joinedEnd(s, end)
	case optionEnd(s) > 0:
		return "옵션", optionEnd(s)
	}
	return "", 0
}
