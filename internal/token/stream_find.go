//ff:func feature=token type=helper control=iteration dimension=1
//ff:what find — the region sent by Next whose span is the one given, or nil
package token

func (s *Stream) find(span Span) *region {
	for _, rg := range s.sent {
		if rg.span == span {
			return rg
		}
	}
	return nil
}
