//ff:func feature=data type=helper control=iteration dimension=1
//ff:what composeHangul — the string with every hangul initial, vowel and final in a row joined into a syllable
package main

func composeHangul(s string) string {
	var out []rune
	for _, r := range s {
		out = mergeJamo(out, r)
	}
	return string(out)
}
