//ff:func feature=token type=reader control=sequence
//ff:what fitsNames — articles 49 and 50: the word before the boundary is a nominal with no part after, and the stem of the word after it is a nominal
package token

func fitsNames(g gap) bool {
	return len(g.left.Parts) == 1 && nominalStem(stemPart(g.left)) && nominalStem(stemPart(g.right))
}
