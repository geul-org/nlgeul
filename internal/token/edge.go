//ff:type feature=token type=model
//ff:what edge — two words that stand side by side in a way of cutting: the span of the first and the span of the one after it
package token

type edge struct {
	from Span
	to   Span
}
