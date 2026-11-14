//ff:func feature=token type=helper control=iteration dimension=1
//ff:what edgeAt — the letter of a line that ends exactly at the cell given (the line counted in cells, a Hangul letter two), and whether there is one; a letter that would stand across the cell does not end there
package token

func edgeAt(line string, cell int) (rune, bool) {
	at := 0
	for _, r := range line {
		at += runeWidth(r)
		if at == cell {
			return r, true
		}
	}
	return 0, false
}
