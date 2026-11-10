//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterUnits — a tail read once more after each run of pieces that follows a number and ends where the tail starts (3개+이+다): the tail keeps what it tells of its closing
package token

func (sc *scan) afterUnits(t tail) []tail {
	if t.base != 0 {
		return nil
	}
	out := joined(sc.units(t.start, 0), t.pieces...)
	for i := range out {
		out[i].noYo, out[i].close = t.noYo, t.close
	}
	return out
}
