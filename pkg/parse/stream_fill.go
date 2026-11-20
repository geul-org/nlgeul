//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what fill — read, a stretch at a time, until a region is queued or no unit is left, and on to the predicate while a candidate read waits for it (held: the words up to the predicate are then queued together); the rules of a sentence then judge what is queued and not yet judged (settle)
package parse

func (s *Stream) fill() {
	for (len(s.queue) == 0 || s.held) && (s.cur != nil || s.read < len(s.units)) {
		s.pull()
	}
	s.settle()
}
