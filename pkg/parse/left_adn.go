//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what leftAdn — the longest reading of left as a predicate with a modifying closing (아는, 전략인, 된다는): how it divides and its length; false when there is none
package parse

func leftAdn(left []rune) (string, int, bool) {
	var best tailFact
	for _, t := range collectTails(left, len(left), false) {
		if t.kind&kindAdn != 0 && t.size > best.size {
			best = t
		}
	}
	return best.marked, best.size, best.size > 0
}
