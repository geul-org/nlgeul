//ff:func feature=token type=reader control=sequence
//ff:what Next — the best way of cutting the next stretch of the sentence, the stretches one after another; false at the end. A stretch with no candidate at all (none is known) is sent as it is, with no nodes and Rank 0. A stretch a window cut off is kept after it is sent only as the letters it covers and how to read it again (Again)
package token

import "math/big"

func (s *Stream) Next() (Segment, bool) {
	s.fill()
	if len(s.queue) == 0 {
		return Segment{}, false
	}
	rg := s.queue[0]
	s.queue = s.queue[1:]
	seg, _ := rg.deal()
	seg.Region = rg.span
	if seg.Ways == nil {
		seg.Ways = new(big.Int)
	}
	if rg.redo != nil {
		rg = &region{span: rg.span, redo: rg.redo, sent: 1}
	}
	s.sent = append(s.sent, rg)
	return seg, true
}
