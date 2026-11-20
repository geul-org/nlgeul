//ff:type feature=parse type=model
//ff:what Span — a place in a sentence: letters from Start up to End (End not included), counted over the letters of the sentence with its spaces left out (a space inside a span of code is not counted either: Node.Contest.Text keeps it, the letters of the stretch do not)
package parse

type Span struct {
	Start int
	End   int
}
