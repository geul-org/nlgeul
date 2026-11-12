//ff:func feature=token type=reader control=sequence
//ff:what readGroup — judge every entry of one class that ends at j and merge the readings of the same token sequence of the same place or kind into one piece (groupAt), with nothing told of a sign that closes the unit
package token

func (sc *scan) readGroup(class, j, next int, more bool) (map[string]piece, map[string]*entry, map[string]int) {
	return sc.groupAt(class, j, next, more, false, false)
}
