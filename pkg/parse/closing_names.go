//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what closingNames — the names the index gives the entries of some closings of the tables, in order: a list that stands for every form of a closing (어, 아, 여 and the fused one; 면서 and 으면서) is made from the closing's row, as its entries are, and not written out a second time
package parse

func closingNames(rows ...Marker) []string {
	var out []string
	for _, m := range rows {
		for _, e := range closingEntries(m) {
			out = append(out, e.name)
		}
	}
	return out
}
