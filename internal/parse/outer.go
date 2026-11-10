//ff:func feature=parse type=reader control=sequence
//ff:what outer — the case of a chain's outermost case particle: the right particle's if it marks one, else the one found to its left
package parse

func outer(right, left int) int {
	if right != 0 {
		return right
	}
	return left
}
