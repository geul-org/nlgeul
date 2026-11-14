//ff:func feature=token type=helper control=sequence
//ff:what memberOf — whether a line of spans is one of the ways of cutting that are listed (by their spans, lineKey): the test a region asks of a way its walks made
package token

func memberOf(forms map[string]bool) func([]Span) bool {
	return func(spans []Span) bool { return forms[lineKey(spans)] }
}
