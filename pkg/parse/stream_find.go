//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what find — the region sent by Next whose span, as sent up (outer), is the one given, or nil
package parse

func (s *Stream) find(span Span) *region {
	for _, rg := range s.sent {
		if s.outer(rg.span) == span {
			return rg
		}
	}
	return nil
}
