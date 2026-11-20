//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what kindPlace — the place of a kind in kindOrder (0 for a kind that is not in it)
package parse

func kindPlace(kind int) int {
	for i, k := range kindOrder {
		if k == kind {
			return i
		}
	}
	return 0
}
