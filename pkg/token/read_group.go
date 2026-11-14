//ff:func feature=token type=reader control=sequence
//ff:what readGroup — judge every entry of one class that ends at j and file each candidate that stands under its own key (groupAt: the same token sequence of the same place or kind, and the same name), with nothing told of a sign that closes the unit
package token

func (sc *scan) readGroup(class, j, next int, more bool) (map[string]piece, map[string]*entry, map[string]int) {
	return sc.groupAt(class, j, next, more, false, false)
}
