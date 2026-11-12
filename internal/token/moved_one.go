//ff:func feature=token type=reader control=iteration dimension=1
//ff:what movedOne — one reading carried over to the letters after a sign at its head: the sign taken off the first word (movedFirst; a word with nothing left is left out), every end counted from the letters after the sign; nil when no word is left
package token

func (sc *scan) movedOne(words []step, head int) *path {
	root := &path{verdict: 1}
	if e, ok := movedFirst(words[0], head, sc.wordOf(0, head, tail{start: head}).pref); ok {
		root = grow(root, e)
	}
	for _, e := range words[1:] {
		e.end -= head
		e.word.End = e.end
		root = grow(root, e)
	}
	if root.prev == nil {
		return nil
	}
	return root
}
