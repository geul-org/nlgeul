//ff:func feature=token type=helper control=iteration dimension=1
//ff:what placeLetters — the unit's letters as one string, and for each place where it begins in that string, so that the letters between two places are a part of the string (letters) and not a copy
package token

import "unicode/utf8"

func (sc *scan) placeLetters() {
	buf := make([]byte, 0, len(sc.s)*3)
	sc.offsets = make([]int, len(sc.s)+1)
	for i, r := range sc.s {
		sc.offsets[i] = len(buf)
		buf = utf8.AppendRune(buf, r)
	}
	sc.offsets[len(sc.s)] = len(buf)
	sc.text = string(buf)
}
