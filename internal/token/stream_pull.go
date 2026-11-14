//ff:func feature=token type=reader control=sequence
//ff:what pull — take the next stretch of the unit being read (the next unit is begun when none is): it is marked as Contests marks it (the unit, Glued) and becomes a region at its place in the sentence, queued, with what is needed to read it again when a window cut it off; the trace lines made are kept; when the unit has no stretch left, the next unit is the one to read
package token

func (s *Stream) pull() {
	u := s.units[s.read]
	if s.cur == nil {
		s.cur, s.k, s.glued = newUnitReader(s.ix, u, s.trace, &s.stats), 0, s.read > 0 && !s.spaced
	}
	st, ok := s.cur.next()
	s.notes = append(s.notes, s.cur.take()...)
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
	s.queue = append(s.queue, rg)
	s.cursor += len([]rune(c.Text))
	s.k++
}
