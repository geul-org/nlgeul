//ff:func feature=token type=helper control=iteration dimension=1
//ff:what cellWidth — the cells a text takes in a terminal, the sum of runeWidth of its letters
package token

func cellWidth(s string) int {
	n := 0
	for _, r := range s {
		n += runeWidth(r)
	}
	return n
}
