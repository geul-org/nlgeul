//ff:func feature=token type=reader control=sequence
//ff:what AgainHead — Again with a head on what is given back: the contest of each node of the next way of the stretch holds only its first head candidates (0 or less: all of them, the same as Again). The ways are made and counted as for Again, so what is given is the first head of what Again gives
package token

func (s *Stream) AgainHead(region Span, head int) (Segment, bool) {
	seg, ok := s.Again(region)
	return segmentHead(seg, head), ok
}
