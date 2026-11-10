//ff:func feature=token type=reader control=iteration dimension=1
//ff:what partsPref — the preference of a stem of bound words: the sum of the words' own (a word the table gives no count costs what its letters would)
package token

func (ix *index) partsPref(parts []string) float64 {
	sum := 0.0
	for _, part := range parts {
		sum += ix.stemPref(part, true, len([]rune(part)))
	}
	return sum
}
