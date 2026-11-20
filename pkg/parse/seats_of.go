//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what seatsOf — the words of the best way of each region given, in order, as the line the rules of a sentence judge (a region with no way gives none)
package parse

func seatsOf(line []*region) []seat {
	var out []seat
	for k, rg := range line {
		seg, _ := rg.peek()
		out = append(out, seatsIn(k, seg.Nodes)...)
	}
	return out
}
