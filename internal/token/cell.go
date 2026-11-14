//ff:type feature=token type=model
//ff:what cell — one word-place of a stretch read by the net: its span (counted in the letters of the unit) and the candidates that stand there, best first by lambda and, of the same lambda, in the order they were born
package token

type cell struct {
	span     Span
	branches []Branch
}
