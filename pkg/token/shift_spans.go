//ff:func feature=token type=helper control=iteration dimension=1
//ff:what shiftSpans — the spans moved by delta places (spans of the unit's letters to those of the sentence, and back)
package token

func shiftSpans(spans []Span, delta int) []Span {
	out := make([]Span, 0, len(spans))
	for _, sp := range spans {
		out = append(out, Span{Start: sp.Start + delta, End: sp.End + delta})
	}
	return out
}
