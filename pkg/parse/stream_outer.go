//ff:func feature=parse type=helper control=sequence
//ff:what outer — a span as the layer above counts it, the letters of the sentence with every space left out, from a span of the stream (which counts the spaces inside a span of code); the same span when no unit holds such a space
package parse

func (s *Stream) outer(sp Span) Span {
	if s.places == nil {
		return sp
	}
	return Span{Start: s.places[sp.Start], End: s.places[sp.End]}
}
