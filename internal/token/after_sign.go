//ff:func feature=token type=helper control=iteration dimension=1
//ff:what afterSign — where the unit of a sign of a pair ends: after what is written onto the sign when one of its candidates closes a pair (a particle is written onto a closing sign), else right after the sign
package token

func (c *cutter) afterSign(from int, reads []pairing) int {
	for _, rd := range reads {
		if rd.name == shutName {
			return c.ontoEnd(from)
		}
	}
	return from
}
