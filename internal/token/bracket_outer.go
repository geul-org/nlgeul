//ff:func feature=token type=helper control=iteration dimension=1
//ff:what drawOuter — the outer bracket of the contest of ways, from the first line of the first way to the first line of the last way shown (starts): "┐" first, "┘" last, "│" between, "┤" at the first line of each way between; the middle line (the first and the last, halved, rounded down) is "├" (or "┼" when it is the first line of a way) and takes the tag; a contest of one way shown is a single "├"
package token

func (b *bracket) drawOuter(starts []int, tag string) {
	first, last := starts[0], starts[len(starts)-1]
	for i := first; i <= last; i++ {
		b.outer[i] = '│'
	}
	for i := 1; i < len(starts)-1; i++ {
		b.outer[starts[i]] = '┤'
	}
	mid := (first + last) / 2
	if first == last {
		b.outer[mid] = '├'
	} else {
		b.outer[first], b.outer[last] = '┐', '┘'
	}
	if b.outer[mid] == '┤' {
		b.outer[mid] = '┼'
	} else if b.outer[mid] == '│' {
		b.outer[mid] = '├'
	}
	b.tag(mid, tag)
}
