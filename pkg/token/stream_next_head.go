//ff:func feature=token type=reader control=sequence
//ff:what NextHead — Next with a head on what is given back: the contest of each node of the way holds only its first head candidates (0 or less: all of them, the same as Next). The candidates are made as for Next and the stream goes on as after Next, so what is given is the first head of what Next gives
package token

// head is a bound on what the caller takes, not on what is read: the
// candidates kept are exactly the first head of all of them, ties in the
// same order.
func (s *Stream) NextHead(head int) (Segment, bool) {
	seg, ok := s.Next()
	return segmentHead(seg, head), ok
}
