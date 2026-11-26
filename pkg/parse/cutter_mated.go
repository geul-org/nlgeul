//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what mated — does the sign at place i, written with a space (or the sentence's edge) on both sides, have a sign that pairs with it written so too on the other side of what it encloses: after an opening sign a closing sign of its kind, before a closing sign an opening one (< 제목 >); signs in a span of code are not counted
package parse

func (c *cutter) mated(i int, row pairRow) bool {
	if !apartAt(c.r, i) {
		return false
	}
	for j, sign := range c.r {
		other, ok := pairRows[sign]
		if !ok || other.kind != row.kind || c.inside[j] || !apartAt(c.r, j) {
			continue
		}
		if (j > i && row.opens && other.shuts) || (j < i && row.shuts && other.opens) {
			return true
		}
	}
	return false
}
