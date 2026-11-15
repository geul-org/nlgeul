//ff:func feature=token type=reader control=iteration dimension=1
//ff:what redoAgain — the next way of a stretch a window cut off: the stretch is read again (rebuild, with what the rules of a sentence added to its candidates: lift) unless the one read again last is of the same stretch, the ways sent already are passed over, and the next is given; only that one stretch read again is held
package token

func (s *Stream) redoAgain(rg *region) (Segment, bool) {
	if s.redoOf != rg {
		s.redone, s.redoOf = rg.redo.rebuild(), rg
		s.redone.lift(rg.lifts)
		for i := 0; i < rg.sent; i++ {
			s.redone.deal()
		}
	}
	seg, ok := s.redone.deal()
	if ok {
		rg.sent++
	}
	return seg, ok
}
