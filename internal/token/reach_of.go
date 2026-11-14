//ff:func feature=token type=helper control=sequence
//ff:what of — the value of the best way to a state before another: 0 for a seed or the root (nothing is before it), the value it was given otherwise, and false when it has none yet
package token

func (rc *reach) of(b *path) (float64, bool) {
	if b.prev == nil {
		return 0, true
	}
	v, ok := rc.value[b]
	return v, ok
}
