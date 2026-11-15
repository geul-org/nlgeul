//ff:type feature=token type=model
//ff:what lift — what the rules of a sentence add to the preference of one candidate: the span of its word, the candidate as it is printed, and how much is added
package token

type lift struct {
	span Span
	word string
	add  float64
}
