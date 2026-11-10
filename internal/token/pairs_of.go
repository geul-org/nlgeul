//ff:func feature=token type=reader control=iteration dimension=1
//ff:what pairsOf — the signs of a sentence that open and close a pair, each place with its partner's: brackets by nesting, quotation marks and backticks two by two in the order they come; an apostrophe between two letters is no quotation mark, and a sign inside a span of code is no sign
package token

// Brackets inside a span of code, and quotation marks, pair only
// outside it: the code is not read.
func pairsOf(r []rune) map[int]int {
	out := map[int]int{}
	code := twos(r, '`', nil)
	for open, shut := range code {
		out[open] = shut
	}
	inside := covered(code, len(r))
	for _, q := range quoteSigns {
		for open, shut := range twos(r, q, inside) {
			out[open] = shut
		}
	}
	for open, shut := range nested(r, inside) {
		out[open] = shut
	}
	return out
}
