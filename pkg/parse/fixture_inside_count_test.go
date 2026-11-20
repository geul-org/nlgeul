//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what insideCount — how many places of a unit are inside an atom
package parse

func insideCount(sc *scan) int {
	n := 0
	for p := range sc.s {
		if sc.inside[p] {
			n++
		}
	}
	return n
}
