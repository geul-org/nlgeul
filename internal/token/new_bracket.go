//ff:func feature=token type=helper control=iteration dimension=1
//ff:what newBracket — a bracket for lines lines, nothing drawn yet (blanks)
package token

func newBracket(lines int) *bracket {
	b := &bracket{inner: make([]rune, lines), outer: make([]rune, lines), tags: make([]string, lines)}
	for i := range b.inner {
		b.inner[i], b.outer[i] = ' ', ' '
	}
	return b
}
