//ff:type feature=token type=model
//ff:what Span — a place in a sentence: letters from Start up to End (End not included), counted over the letters of the sentence with its spaces left out
package token

type Span struct {
	Start int
	End   int
}
