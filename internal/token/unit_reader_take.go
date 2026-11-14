//ff:func feature=token type=helper control=sequence
//ff:what take — the trace lines of the unit made since they were last taken, in the order they were made; taken once
package token

func (r *unitReader) take() []string {
	out := r.notes
	r.notes = nil
	return out
}
