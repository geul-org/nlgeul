//ff:func feature=token type=reader control=sequence
//ff:what wordOf — the first reading of stepsOf: the word from b to end with its markers; the one reading of a word that has one
package token

func (sc *scan) wordOf(b, end int, t tail) step {
	return sc.stepsOf(b, end, t)[0]
}
