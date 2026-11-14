//ff:func feature=token type=reader control=iteration dimension=1
//ff:what next — the next stretch of the unit: those made first, then those of the net as it is worked out, then those of the signs after the letters; false when none is left
package token

func (r *unitReader) next() (unitStretch, bool) {
	for {
		if len(r.queue) > 0 {
			st := r.queue[0]
			r.queue = r.queue[1:]
			return st, true
		}
		if r.run == nil {
			return unitStretch{}, false
		}
		if st, ok := r.run.next(); ok {
			return st, true
		}
		r.run, r.queue = nil, r.after()
	}
}
