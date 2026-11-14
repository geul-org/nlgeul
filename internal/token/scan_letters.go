//ff:func feature=token type=helper control=sequence
//ff:what letters — the letters of the unit from place b to place e as a string: a part of the unit's string when the scan was prepared (no copy), else made from the letters
package token

func (sc *scan) letters(b, e int) string {
	if len(sc.offsets) == len(sc.s)+1 {
		return sc.text[sc.offsets[b]:sc.offsets[e]]
	}
	return string(sc.s[b:e])
}
