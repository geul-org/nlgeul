//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what droppedIn — does a tail hold the copula's stem dropped after a vowel ((이), which leaves no trace in the letters)
package parse

func droppedIn(t tail) bool {
	for _, p := range t.pieces {
		if p.names == droppedName {
			return true
		}
	}
	return false
}
