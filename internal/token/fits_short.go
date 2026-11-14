//ff:func feature=token type=reader control=sequence
//ff:what fitsShort — article 46: both words of the boundary are one letter long without their particles
package token

func fitsShort(g gap) bool {
	return bodyLen(g.left, g.leftLen) == 1 && bodyLen(g.right, g.rightLen) == 1
}
