//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what midStart — where the run of one sign that ends right before place end begins (a sign between letters may stand two or more in a row: ——); the place before end itself when the letters there are not this unit's
package parse

func (sc *scan) midStart(end int) int {
	q := end - 1
	for q > 0 && q < len(sc.s) && sc.s[q-1] == sc.s[q] {
		q--
	}
	return q
}
