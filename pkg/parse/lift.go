//ff:type feature=parse type=model
//ff:what lift — what the rules of a sentence do to one candidate: the span of its word, the candidate as it is printed, how much is added to its preference (add), and the force of a rule of rebuttal that stands against it (force: its verdict is multiplied by 2/(1+force)-1; 0 for none, which multiplies by 1)
package parse

type lift struct {
	span  Span
	word  string
	add   float64
	force float64
}
