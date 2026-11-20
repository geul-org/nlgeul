//ff:func feature=parse type=reader control=sequence
//ff:what wordOf — the first candidate of stepsOf: the word from b to end with its markers; the one candidate of a word that has one
package parse

func (sc *scan) wordOf(b, end int, t tail) step {
	return sc.stepsOf(b, end, t)[0]
}
