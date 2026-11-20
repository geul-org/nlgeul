//ff:func feature=parse type=reader control=sequence
//ff:what Again — the way of cutting a stretch after the one sent last for it (Next sent the first), the next best by lambda; false when the stretch is not one Next has sent or no way is left. It makes only that one way, and does not move Next. A stretch a window cut off is read again for it (redoAgain). What is sent has its places as the layer above counts them (outerSegment), and the stretch asked for is given in those places
package parse

func (s *Stream) Again(region Span) (Segment, bool) {
	rg := s.find(region)
	if rg == nil {
		return Segment{}, false
	}
	if rg.redo != nil {
		seg, ok := s.redoAgain(rg)
		return s.outerSegment(seg), ok
	}
	seg, ok := rg.deal()
	return s.outerSegment(seg), ok
}
