//ff:func feature=token type=reader control=sequence
//ff:what newCutter — a cutter at the start of a sentence: its spans of code found (backticks two by two), and one way the signs of a pair stand open, with none open
package token

func newCutter(ix *index, sentence string, trace bool) *cutter {
	r := []rune(sentence)
	code := twos(r, '`', nil)
	return &cutter{ix: ix, r: r, code: code, inside: covered(code, len(r)), open: []string{""}, trace: trace}
}
