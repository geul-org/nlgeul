//ff:func feature=token type=helper control=iteration dimension=1
//ff:what maxCells — the widest of the texts in cells (0 for none)
package token

func maxCells(ss []string) int {
	w := 0
	for _, s := range ss {
		w = max(w, cellWidth(s))
	}
	return w
}
