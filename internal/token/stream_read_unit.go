//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readNext — read the next unit of the stream: its stretches are read by the net (readUnitNet) and marked as Contests marks them (the unit, Glued) and each becomes a region at its place in the sentence, queued in order; the trace lines of the unit are kept
package token

func (s *Stream) readNext() {
	u := s.units[s.read]
	stretches, notes := readUnitNet(s.ix, u, s.trace)
	for k, st := range stretches {
		c := st.contest
		c.Unit, c.Glued = s.read, k > 0 || (s.read > 0 && !s.spaced)
		s.queue = append(s.queue, st.region(c, s.cursor))
		s.cursor += len([]rune(c.Text))
	}
	s.notes = append(s.notes, notes...)
	s.spaced = u.spaced
	s.read++
}
