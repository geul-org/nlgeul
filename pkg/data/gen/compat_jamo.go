//ff:func feature=data type=helper control=iteration dimension=1
//ff:what compatJamo — joins hangul jamo into syllables and turns the initials and vowels left alone into compatibility jamo (the parenthesized initial becomes ㄱ, initial and vowel become 가)
package main

import "strings"

func compatJamo(jamo map[rune]rune, s string) string {
	var b strings.Builder
	for _, r := range composeHangul(s) {
		if c, ok := jamo[r]; ok {
			r = c
		}
		b.WriteRune(r)
	}
	return b.String()
}
