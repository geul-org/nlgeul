//ff:type feature=token type=model
//ff:what cut — one way of cutting a stretch into words: the span of each word, and with it that word as a branch of its own (one word, its verdict and its preference)
package token

type cut struct {
	spans []Span
	words []Branch
}
