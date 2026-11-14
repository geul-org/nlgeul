//ff:func feature=token type=reader control=iteration dimension=1
//ff:what front — the states at the window's end: those that end where a word may still begin that the cut must not split (open); none (false) when that place is where the last cut was, since then no word read so far is before it
package token

func (r *netRun) front() ([]*path, bool) {
	open := r.open()
	if open[r.c] {
		return nil, false
	}
	var out []*path
	for _, s := range r.sc.registry().all {
		if open[s.last.end] {
			out = append(out, s)
		}
	}
	return out, len(out) > 0
}
