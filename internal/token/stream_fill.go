//ff:func feature=token type=reader control=iteration dimension=1
//ff:what fill — read units, one at a time, until a region is queued or no unit is left
package token

func (s *Stream) fill() {
	for len(s.queue) == 0 && s.read < len(s.units) {
		s.readNext()
	}
}
