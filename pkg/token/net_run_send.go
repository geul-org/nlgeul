//ff:func feature=token type=helper control=sequence
//ff:what send — a stretch taken from the net as it is sent: the places worked out when the first of a net is sent are noted (stats), and the stretches are counted
package token

func (r *netRun) send(st unitStretch) unitStretch {
	if r.made == 0 {
		r.stats.firsts = append(r.stats.firsts, r.j)
	}
	r.made++
	return st
}
