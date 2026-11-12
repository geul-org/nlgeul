//ff:func feature=token type=reader control=iteration dimension=1
//ff:what rootGrown — the readings of a unit that is one word with no marker before it: the root grown by the word from 0 to n that tail t makes, once for each way its stem stands
package token

func (sc *scan) rootGrown(root *path, n int, t tail) []*path {
	var out []*path
	for _, e := range sc.stepsOf(0, n, t) {
		out = append(out, grow(root, e))
	}
	return out
}
