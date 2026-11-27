//ff:func feature=data type=helper control=sequence
//ff:what mergeJamo — adds one character after the ones so far, joining a hangul initial and vowel (and a final) into a syllable as normalization form C does
package main

const (
	syllableBase = 0xAC00
	vowelCount   = 21
	finalCount   = 28
)

func mergeJamo(out []rune, r rune) []rune {
	n := len(out)
	if n == 0 {
		return append(out, r)
	}
	last := out[n-1]
	if last >= 0x1100 && last <= 0x1112 && r >= 0x1161 && r <= 0x1175 {
		out[n-1] = syllableBase + ((last-0x1100)*vowelCount+(r-0x1161))*finalCount
		return out
	}
	if last >= syllableBase && last <= 0xD7A3 && (last-syllableBase)%finalCount == 0 && r >= 0x11A8 && r <= 0x11C2 {
		out[n-1] = last + (r - 0x11A7)
		return out
	}
	return append(out, r)
}
