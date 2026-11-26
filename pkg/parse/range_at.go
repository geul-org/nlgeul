//ff:func feature=parse type=reader control=sequence
//ff:what rangeAt — whether the sign at place q is the range sign written right after a number and its unit and right before a digit (the tilde of 1월~3월: tildeAt reads it as a piece of the number's word, 범위): it is not read as a word of its own there
package parse

import (
	"strings"
	"unicode"
)

func (sc *scan) rangeAt(q int) bool {
	if !strings.ContainsRune(tildeSigns, sc.s[q]) || !unicode.IsDigit(sc.at(q+1)) {
		return false
	}
	return len(sc.units(q, 0, false)) > 0
}
