//ff:func feature=token type=reader control=sequence
//ff:what pull — take the next stretch of the unit being read (the next unit is begun when none is): it is marked with its unit and Glued (true for every stretch of a unit after its first, and for the first after a unit with no space after it) and becomes a region at its place in the sentence, queued, with what is needed to read it again when a window cut it off; the trace lines made since the last region go with the region (given when it is sent), and the region tells whether the stream must read on to a predicate (hold); when the unit has no stretch left, the next unit is the one to read
package token

func (s *Stream) pull() {
	u := s.units[s.read]
	if s.cur == nil {
		s.cur, s.k, s.glued = newUnitReader(s.ix, u, s.trace, &s.stats), 0, s.read > 0 && !s.spaced
	}
	st, ok := s.cur.next()
	s.pending = append(s.pending, s.cur.take()...)
	if !ok {
		s.cur, s.spaced = nil, u.spaced
		s.read++
		return
	}
	c := st.contest
	c.Unit, c.Glued = s.read, s.k > 0 || s.glued
	rg := st.region(c, s.cursor)
	if st.redo != nil {
		st.redo.head, st.redo.at = c, s.cursor
		rg.redo = st.redo
	}
	rg.notes, s.pending = s.pending, nil
	s.queue = append(s.queue, rg)
	s.hold(rg)
	s.cursor += len([]rune(c.Text))
	s.k++
}
