//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterAt — whether a piece the tables read after a number (a greater number, 여, a unit, 째) begins at a place of the unit
package token

func (sc *scan) afterAt(k int) bool {
	for j := k + 1; j <= len(sc.s); j++ {
		if sc.readsAfter(string(sc.s[k:j]), 0) {
			return true
		}
	}
	return false
}
