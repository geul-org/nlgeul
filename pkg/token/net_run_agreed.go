//ff:func feature=token type=reader control=sequence
//ff:what agreed — the first spans the best ways to every state at the window's end begin with (the front; the spans of a way are the lines of reach), none when no state is there or a word may still begin where the last cut was; the last span of the part is left out when a way ends with it (withWordAfter)
package token

func (r *netRun) agreed() []Span {
	front, ok := r.front()
	if !ok {
		return nil
	}
	lines := r.lines(front)
	return withWordAfter(commonSpans(lines), lines)
}
