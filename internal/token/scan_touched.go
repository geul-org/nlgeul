//ff:func feature=token type=reader control=sequence
//ff:what touched — the candidates of one word-place of a unit as readUnit finishes them: at the place where the unit begins the sign of a pair at its head is given once for each way it is read (spread) when head is set, the signs between letters are given the candidates of that sign (expandMid), and at the place where the unit begins a sign of the unit table as a unit of measure is added (measured); the order of the candidates and what they hold are kept
package token

func (sc *scan) touched(bs []Branch, u unit, head bool) []Branch {
	c := []Contest{{Branches: bs}}
	if head {
		c = spread(c, u)
	}
	c = sc.expandMid(c)
	if head {
		c = sc.ix.measured(c, u)
	}
	return c[0].Branches
}
