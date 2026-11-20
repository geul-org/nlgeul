//ff:type feature=parse type=model
//ff:what edge — two words that stand side by side in a way of cutting: the span of the first and the span of the one after it
package parse

type edge struct {
	from Span
	to   Span
}
