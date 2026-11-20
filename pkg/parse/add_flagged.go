//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what addFlagged — a word put at the end of the list of each name whose bit is set in the bits given
package parse

func addFlagged(out map[string][]string, word string, bits int, names map[string]int) {
	for name, bit := range names {
		if bits&bit != 0 {
			out[name] = append(out[name], word)
		}
	}
}
