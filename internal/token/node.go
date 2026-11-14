//ff:type feature=token type=model
//ff:what Node — one word of a stream (an eojeol, as the layer above reads it): where it stands and the contest of its candidates, each candidate one word (a stem with the particles or endings after it), best first by lambda
package token

// Contest.Text is the letters of the span, Contest.Unit the unit of the
// sentence it is in, and Contest.Glued whether no space stood right
// before it. Every branch of the contest holds exactly one word.
type Node struct {
	Span    Span
	Contest Contest
}
