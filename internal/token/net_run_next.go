//ff:func feature=token type=reader control=iteration dimension=1
//ff:what next — the next stretch of the unit: those made and not yet sent first; else a cut is tried (tryCut) before each place is worked out (step), and when the unit's end is reached the stretches that are left are made (finish); false when none is left
package token

func (r *netRun) next() (unitStretch, bool) {
	for {
		if len(r.rest) > 0 {
			st := r.rest[0]
			r.rest = r.rest[1:]
			return r.send(st), true
		}
		if r.done {
			return unitStretch{}, false
		}
		if st, ok := r.tryCut(); ok {
			return r.send(st), true
		}
		if r.j == r.n {
			r.rest, r.done = r.finish(), true
			continue
		}
		r.step()
	}
}
