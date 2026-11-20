//ff:func feature=parse type=reader control=sequence
//ff:what Notes — the trace lines of the units read since Notes was last called, in the order they were read (empty for a Stream made without StreamTrace); taken once
package parse

func (s *Stream) Notes() []string {
	out := s.notes
	s.notes = nil
	return out
}
