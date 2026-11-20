//ff:func feature=parse type=reader control=sequence
//ff:what Next — the best way of cutting the next stretch of the sentence, the stretches one after another; false at the end. A stretch with no candidate at all (none is known) is sent as it is, with no nodes and Rank 0. A stretch a window cut off is kept after it is sent only as the letters it covers and how to read it again (Again), with what the rules of a sentence added to its candidates. The trace lines of a stretch are given with it (Notes), those left at the end with the last false. What is sent has its places as the layer above counts them (outerSegment)
package parse

import "math/big"

func (s *Stream) Next() (Segment, bool) {
	s.fill()
	if len(s.queue) == 0 {
		s.notes, s.pending = append(s.notes, s.pending...), nil
		return Segment{}, false
	}
	rg := s.queue[0]
	s.queue = s.queue[1:]
	s.notes = append(s.notes, rg.notes...)
	seg, _ := rg.deal()
	seg.Region = rg.span
	if seg.Ways == nil {
		seg.Ways = new(big.Int)
	}
	if rg.redo != nil {
		rg = &region{span: rg.span, redo: rg.redo, sent: 1, lifts: rg.lifts}
	}
	s.sent = append(s.sent, rg)
	return s.outerSegment(seg), true
}
