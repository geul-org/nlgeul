//ff:func feature=data type=helper control=iteration dimension=1
//ff:what jamoMap — from each initial or vowel jamo to the compatibility jamo that decomposes to it (U+3131 to U+318E)
package main

func jamoMap(t *ucdTable) map[rune]rune {
	out := map[rune]rune{}
	for o := rune(0x3131); o <= 0x318E; o++ {
		d := t.Chars[o].Decomp
		if tagOf(d) != "<compat>" {
			continue
		}
		v := decompValue(d)
		if len([]rune(v)) == 1 {
			out[[]rune(v)[0]] = o
		}
	}
	return out
}
