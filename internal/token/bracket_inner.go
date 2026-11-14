//ff:func feature=token type=helper control=iteration dimension=1
//ff:what drawInner — the inner bracket of one way, from the line of the title of its first box to the line of the title of its last: "┬" first, "│" between, "┘" last (a single "─" when they are one line), with the tag on the first line
package token

func (b *bracket) drawInner(first, last int, tag string) {
	if first == last {
		b.inner[first] = '─'
	} else {
		b.inner[first], b.inner[last] = '┬', '┘'
	}
	for i := first + 1; i < last; i++ {
		b.inner[i] = '│'
	}
	b.tag(first, tag)
}
