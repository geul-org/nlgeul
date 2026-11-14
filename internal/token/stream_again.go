//ff:func feature=token type=reader control=sequence
//ff:what Again — the way of cutting a stretch after the one sent last for it (Next sent the first), the next best by lambda; false when the stretch is not one Next has sent or no way is left. It makes only that one way, and does not move Next. A stretch a window cut off is read again for it (redoAgain)
package token

func (s *Stream) Again(region Span) (Segment, bool) {
	rg := s.find(region)
	if rg == nil {
		return Segment{}, false
	}
	if rg.redo != nil {
		return s.redoAgain(rg)
	}
	return rg.deal()
}
