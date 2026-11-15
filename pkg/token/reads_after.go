//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readsAfter — whether letters are a piece the tables read after a number (a greater number, 여, a unit, 째), of the place given (any place for 0)
package token

func (sc *scan) readsAfter(form string, place int) bool {
	f := []rune(form)
	if len(f) == 0 {
		return false
	}
	for _, e := range sc.ix.after[f[len(f)-1]] {
		if e.form == form && (place == 0 || e.slot == place) {
			return true
		}
	}
	return false
}
